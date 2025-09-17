package rest

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, code int, b []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b)
}

func RespondJSON(w http.ResponseWriter, code int, response any) {
	b, err := json.Marshal(response)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	Respond(w, code, b)
}

func RespondError(w http.ResponseWriter, code int, err error) {
	RespondJSON(w, code, &Error{
		Code:    code,
		Message: err.Error(),
	})
}
