package route

import (
	"github.com/gorilla/mux"
	"github.com/mykytaserdiuk/shaream/pkg/inventory/handler"
	"github.com/mykytaserdiuk/shaream/pkg/inventory/service"
)

func NewRouter(services *service.Services) *mux.Router {
	router := mux.NewRouter().PathPrefix("/v1/api").Subrouter()

	router.HandleFunc("/file/upload", handler.NewUploadFileHandler(services.FileSvc)).
		Methods("POST")

	return router

}
