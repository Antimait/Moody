package api

import (
	"bufio"
	"errors"
	"log"
	"moody/communication"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

func HttpListenAndServer(port string, panelPath string) *http.Server {
	router := mux.NewRouter()

	// Panel
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(panelPath)))

	// Situation
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/situation", SituationsMux)
	apiRouter.HandleFunc("/situation/{id}", SituationMux)

	// Service
	apiRouter.HandleFunc("/service", ServicesMux)
	apiRouter.HandleFunc("/service/{id}", ServiceMux)

	// Internal Gateway endpoints
	apiRouter.HandleFunc("/neural_state", neuralStateMux)
	apiRouter.HandleFunc("/actuator_mode", actuatorModeMux)
	apiRouter.HandleFunc("/actuators", actuatorMux)
	apiRouter.HandleFunc("/sensor_service", serviceMux)
	apiRouter.HandleFunc("/data_table", tableMux)
	apiRouter.HandleFunc("/current_situation", situationMux)
	apiRouter.HandleFunc("/service_ws", communication.ServeServiceWS)
	apiRouter.HandleFunc("/actuator_ws", communication.ServeActuatorWS)

	apiRouter.Use(allowAllCorsMiddleware)
	apiRouter.Use(logRequestResponseMiddleWare)
	server := &http.Server{Addr: port, Handler: router}
	return server
}

func allowAllCorsMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("origin")
		if origin == "" {
			origin = "*"
		}
		applyHeaders(origin, &w)
		if r.Method == http.MethodOptions {
			respOptions(origin, &w)
		}
		h.ServeHTTP(w, r)
	})
}

func logRequestResponseMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)

		lrw := NewLoggingResponseWriter(w, r)
		h.ServeHTTP(lrw, r)

		statusCode := lrw.statusCode
		method := lrw.method
		url := lrw.url

		log.Printf("%s %s %d %s", method, url, statusCode, http.StatusText(statusCode))
	})
}

// CORS Middleware Headers

func applyHeaders(origin string, w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", origin)
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE, PUT")
	(*w).Header().Set("Access-Control-Allow-Headers", "DNT,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range,Access-Control-Allow-Headers,Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")
	(*w).Header().Set("Access-Control-Expose-Headers", "Content-Length,Content-Range")
}

func respOptions(origin string, w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", origin)
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE, PUT")
	(*w).Header().Set("Access-Control-Allow-Headers", "DNT,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range,Access-Control-Allow-Headers,Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")
	(*w).Header().Set("Access-Control-Max-Age", "1728000")
	(*w).Header().Set("Content-Type", "text/plain; charset=utf-8")
	(*w).Header().Set("Content-Length", "0")
	(*w).WriteHeader(http.StatusNoContent)
}

// Request/Response logger

type loggingResponseWriter struct {
	http.ResponseWriter
	request    *http.Request
	statusCode int
	method     string
	url        string
}

func NewLoggingResponseWriter(w http.ResponseWriter, r *http.Request) *loggingResponseWriter {
	return &loggingResponseWriter{w, r, http.StatusOK, "", ""}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.method = lrw.request.Method
	lrw.url = lrw.request.URL.Path
}

func (lrw *loggingResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := lrw.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("hijack not supported")
	}
	return h.Hijack()
}
