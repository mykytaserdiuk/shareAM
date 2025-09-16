package handler

import (
	"encoding/json"
	"github.com/mykytaserdiuk/shaream/pkg/models"
	"github.com/mykytaserdiuk/shaream/pkg/user-management/service"
	"net/http"
)

func NewCreateUserHandler(svc service.UserSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var credentials models.Credentials
		err := json.NewDecoder(r.Body).Decode(&credentials)
		if err != nil {
			return
		}

		user, err := svc.CreateUser(r.Context(), &credentials)
		if err != nil {
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}
