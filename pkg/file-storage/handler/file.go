package handler

import (
	"github.com/mykytaserdiuk/shaream/pkg/file-storage/service"
	"net/http"
)

func NewUploadFileHandler(svc service.FileSvc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("User-Token")

		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Ошибка при получении файла: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		link, err := svc.Upload(r.Context(), &token, file, header)
		if err != nil {
			http.Error(w, "Ошибка при сохранинии файла: "+err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(200)
		w.Write([]byte(link))
	})
}
