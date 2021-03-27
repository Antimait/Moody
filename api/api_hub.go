package api

import "net/http"

func neuralStateMux(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getNeuralState(w, r)
	case http.MethodPut:
		setNeuralState(w, r)
	case http.MethodOptions:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func actuatorModeMux(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getActuatorMode(w, r)
	case http.MethodPost:
		setActuatorMode(w, r)
	case http.MethodOptions:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func actuatorMux(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getActuators(w, r)
	case http.MethodPost:
		addMapping(w, r)
	case http.MethodDelete:
		deleteMappings(w, r)
	case http.MethodOptions:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func serviceMux(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		activateService(w, r)
	case http.MethodGet:
		getActivatedServices(w, r)
	case http.MethodDelete:
		deactivateService(w, r)
	case http.MethodOptions:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func situationMux(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		setHubSituation(w, r)
	case http.MethodGet:
		getHubSituation(w, r)
	case http.MethodOptions:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func tableMux(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getDataTable(w, r)
	case http.MethodOptions:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
