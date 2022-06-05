package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes(r *mux.Router, service Scaffold) {
	r.HandleFunc("/", Invoker[CreateRequest, CreateResponse](service.Create)).Methods(http.MethodPost)
}
