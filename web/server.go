package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/eubyt/goeubyt/urlshort"
)

type Response struct {
	Message string `json:"message"`
}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}

func JsonHandler(w http.ResponseWriter, r *http.Request, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func ListenAndServe() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		JsonHandler(w, r, &Response{Message: "pong"}, http.StatusOK)
	})

	mux.HandleFunc("/api/shorten", RegisterShortener)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		const mydir = "./"
		isFileExist := checkFileExists(mydir + r.URL.Path)

		if isFileExist {
			// Pagina HTML - Index
			http.FileServer(http.Dir(mydir)).ServeHTTP(w, r)
		} else {
			// Se não achar o nome do arquivo na pasta do front-end realizar o redirect para os sites - Url Shortener
			result, err := urlshort.SearchShortURL(strings.Split(r.URL.Path, "/")[1])
			if err != nil {
				JsonHandler(w, r, &Response{Message: "not exist"}, http.StatusNotFound)
				return
			}
			http.Redirect(w, r, result.Url, http.StatusPermanentRedirect)
		}
	})

	println("Listening on port 3000")

	return http.ListenAndServe(":3000", mux)
}
