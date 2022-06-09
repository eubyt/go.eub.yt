package web

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/eubyt/goeubyt/urlshort"
)

type UrlShortener struct {
	Url string
}

func IsUrl(s string) bool {
	u, err := url.ParseRequestURI(s)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func RegisterShortener(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var urlShortener UrlShortener
		err := decoder.Decode(&urlShortener)

		if err != nil || urlShortener.Url == "" {
			JsonHandler(w, r, &Response{Message: "Invalid JSON"}, http.StatusUnprocessableEntity)
			return
		}

		if !IsUrl(urlShortener.Url) {
			JsonHandler(w, r, &Response{Message: "Invalid URL"}, http.StatusUnprocessableEntity)
			return
		}

		shortener, err := urlshort.CreateShortURL(urlShortener.Url)
		if err != nil {
			panic(err)
		}

		JsonHandler(w, r, shortener, http.StatusCreated)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
