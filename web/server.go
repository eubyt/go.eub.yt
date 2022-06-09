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

func staticFilesHandler(w http.ResponseWriter, r *http.Request) bool {
	const staticDir = "/go/bin"
	path := r.URL.Path
	extensionList := []string{".css", ".js", ".png", ".jpg", ".jpeg", ".gif", ".ico", ".svg", ".txt", ".html"}

	if path == "/" {
		path = "/index.html"
	}

	for _, extension := range extensionList {
		if strings.HasSuffix(path, extension) {
			isFileExist := checkFileExists(staticDir + r.URL.Path)
			if isFileExist {
				http.FileServer(http.Dir(staticDir)).ServeHTTP(w, r)
				return true
			}
		}
	}

	return false
}

func ListenAndServe() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		JsonHandler(w, r, &Response{Message: "pong"}, http.StatusOK)
	})

	mux.HandleFunc("/api/shorten", RegisterShortener)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Se não achar o nome do arquivo na pasta do front-end realizar o redirect para os sites - Url Shortener
		if !staticFilesHandler(w, r) {
			result, err := urlshort.SearchShortURL(strings.Split(r.URL.Path, "/")[1])
			if err != nil {
				JsonHandler(w, r, &Response{Message: "not exist"}, http.StatusNotFound)
				return
			}
			http.Redirect(w, r, result.Url, http.StatusPermanentRedirect)
		}
	})

	println("Listening on port 8080")

	return http.ListenAndServe(":8080", mux)
}
