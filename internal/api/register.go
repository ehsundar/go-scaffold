package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes(r *mux.Router, service Scaffold) {
	r.HandleFunc("/", Invoke[CreateRequest, CreateResponse](service.Create)).
		Methods(http.MethodPost)
	r.HandleFunc("/{id:[0-9]+}/", Invoke[RetrieveRequest, RetrieveResponse](service.Retrieve)).
		Methods(http.MethodGet)
	r.HandleFunc("/{id:[0-9]+}/", Invoke[DeleteRequest, DeleteResponse](service.Delete)).
		Methods(http.MethodDelete)
}
