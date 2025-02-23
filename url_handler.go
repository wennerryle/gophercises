package main

import (
	"fmt"
	"net/http"
)

func mapHandler(
	provider func(string) (string, error),
	w http.ResponseWriter,
	r *http.Request) {

	longUrl, err := provider(r.URL.Path)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, err.Error())
		return
	}

	w.Header().Add("Location", longUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
