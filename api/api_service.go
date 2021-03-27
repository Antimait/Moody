package api

import "net/http"

// Get all the Services or add one
func ServicesMux(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getServices(w, r)
	case http.MethodPost:
		postService(w, r)
	case http.MethodOptions:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Get/delete a situation or patch it to activate it
func ServiceMux(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getService(w, r)
	case http.MethodDelete:
		deleteService(w, r)
	case http.MethodOptions:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
