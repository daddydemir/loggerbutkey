package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func MainRouting() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	// endpoint
	url := "/api/v1/"
	r.HandleFunc(url+"uploadImage", UploadImage).Methods("POST")

	return r
}

func UploadImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	file, handler, err := r.FormFile("file")
	if err != nil {
		// bad request
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "işlem başarısız oldu."})
	}

	path := filepath.Join(".", "files")
	os.MkdirAll(path, os.ModePerm)
	fullpath := path + "/" + handler.Filename
	myfile, _ := os.OpenFile(fullpath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	defer myfile.Close()

	io.Copy(myfile, file)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "dosya kaydedildi."})

}
