package handler

import (
	"encoding/json"
	"errors"
	"github.com/mykytaserdiuk/shaream/pkg/models"
	"github.com/mykytaserdiuk/shaream/pkg/rest"
	"github.com/mykytaserdiuk/shaream/pkg/user-management/service"
	"net/http"
)

func NewCreateUserHandler(svc service.UserSvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var credentials models.Credentials
		err := json.NewDecoder(r.Body).Decode(&credentials)
		if err != nil {
			rest.RespondError(w, http.StatusBadRequest, err)
			return
		}
		if credentials.Password == "" {
			rest.RespondError(w, http.StatusBadRequest, errors.New("pass is empty"))
			return
		}

		user, err := svc.CreateUser(r.Context(), &credentials)
		if err != nil {
			rest.RespondError(w, http.StatusInternalServerError, err)
			return
		}

		rest.RespondJSON(w, http.StatusOK, user)
	}
}
