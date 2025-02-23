package main

import (
	"fmt"
	"net/http"
)

type UrlHandlerProvider interface {
	Get(string) (string, error)
}

func mapHandler(
	provider UrlHandlerProvider,
	w http.ResponseWriter,
	r *http.Request) {

	longUrl, err := provider.Get(r.URL.Path)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Add("Location", longUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
