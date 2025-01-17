package api

import (
	"log"
	"moody/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type servicesResponse struct {
	Services     []*models.Service `json:"services"`
	ServiceCount int64             `json:"count"`
}

// Service API handlers

func getServices(w http.ResponseWriter, _ *http.Request) {
	services, count, err := models.GetAllServices()
	if err != nil {
		services = []*models.Service{}
	}
	resp := servicesResponse{
		Services:     services,
		ServiceCount: count,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	models.MustEncode(w, resp)
}

func postService(w http.ResponseWriter, r *http.Request) {
	newService := &models.Service{}
	w.Header().Set("Content-Type", "application/json")
	ok := models.MustValidate(r, newService)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		models.MustEncode(w, models.ErrorResponse{"Bad syntax"})
		return
	}

	if err := models.AddService(newService); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		models.MustEncode(w, models.ErrorResponse{"Record with pk already exists"})
		return
	}
	w.WriteHeader(http.StatusOK)
	models.MustEncode(w, newService)
	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}
}

// Get handler for /situation/{name}
func getService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		models.MustEncode(w, models.ErrorResponse{"Bad syntax"})
		return
	}
	service, err := models.GetService(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		models.MustEncode(w, models.ErrorResponse{Error: "not found"})
		return
	}
	// here the response is the situation itself
	w.WriteHeader(http.StatusOK)
	models.MustEncode(w, service)
}

// Delete handler for /situation/{name}
func deleteService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		models.MustEncode(w, models.ErrorResponse{"Bad syntax"})
		return
	}

	service, err := models.GetService(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		models.MustEncode(w, models.ErrorResponse{"The situation does not exist"})
		return
	}

	if err := models.DeleteService(service); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		models.MustEncode(w, models.ErrorResponse{"An error occurred while trying to delete the situation"})
		return
	}
	w.WriteHeader(http.StatusOK)
	models.MustEncode(w, service)
}
