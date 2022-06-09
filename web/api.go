package web

import (
	"encoding/json"
	"net/http"
	"net/url"
	"regexp"

	"github.com/eubyt/goeubyt/urlshort"
)

type UrlShortener struct {
	Url       string
	CustomUrl string
}

func isUrl(s string) bool {
	u, err := url.ParseRequestURI(s)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func validCustomUrl(customUrl string) bool {
	regexTest := regexp.MustCompile(`^[A-zZ0-9_]*$`)
	return regexTest.MatchString(customUrl)
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

		if !isUrl(urlShortener.Url) {
			JsonHandler(w, r, &Response{Message: "Invalid URL"}, http.StatusUnprocessableEntity)
			return
		}

		if urlShortener.CustomUrl != "" && !validCustomUrl(urlShortener.CustomUrl) {
			JsonHandler(w, r, &Response{Message: "Invalid Custom URL"}, http.StatusUnprocessableEntity)
			return
		}

		shortener, err := urlshort.CreateShortURL(urlShortener.Url, urlShortener.CustomUrl)
		if err != nil {
			panic(err)
		}

		JsonHandler(w, r, shortener, http.StatusCreated)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
