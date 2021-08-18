package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *mux.Router, service Service) {
	res := resource{service}

	r.HandleFunc("/user", res.create).Methods("POST")
	r.HandleFunc("/user/{id}", res.get).Methods("GET")
}

type resource struct {
	service Service
}

func (r resource) get(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	result, err := r.service.Get(request.Context(), params["id"])
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Error("Ошибка во время получения пользователя")
		return
	}
	response.Header().Add("content-type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}

func (r resource) create(response http.ResponseWriter, request *http.Request) {
	var input CreateUserRequest
	_ = json.NewDecoder(request.Body).Decode(&input)
	err := r.service.Create(request.Context(), input)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		log.WithError(err).Error("Ошибка во время создания пользователя")
		return
	}
	response.Header().Add("content-type", "application/json")
	response.WriteHeader(http.StatusCreated)
}
