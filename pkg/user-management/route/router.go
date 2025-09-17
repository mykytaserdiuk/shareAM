package route

import (
	"github.com/gorilla/mux"
	"github.com/mykytaserdiuk/shaream/pkg/user-management/handler"
	"github.com/mykytaserdiuk/shaream/pkg/user-management/service"
)

func NewRouter(services *service.Services) *mux.Router {
	router := mux.NewRouter().PathPrefix("/api/v1").Subrouter()

	router.HandleFunc("/users", handler.NewCreateUserHandler(services.UserSvc)).
		Methods("POST")

	return router

}
