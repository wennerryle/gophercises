package main

import (
	"errors"
	"fmt"
	"net/http"
)

type TempRealization map[string]string

func (m TempRealization) Get(key string) (string, error) {
	v, exists := m[key]
	if !exists {
		return "", errors.New("that URL address was not found")
	}
	return v, nil
}

func (m TempRealization) Set(key string, value string) error {
	m[key] = value
	return nil
}

// Build the MapHandler using the mux as the fallback
var db = TempRealization{
	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
}

func main() {
	mux := defaultMux()

	fmt.Println("Starting the server on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println(err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/set_urls", setUrlsHandler)
	mux.HandleFunc("/", rootHandler)
	return mux
}

type AddressBind struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func setUrlsHandler(w http.ResponseWriter, r *http.Request) {
	var addressBind []AddressBind

	err := parseFormat(w, r, &addressBind)

	fmt.Printf("setUrlsHandler err: %v\n", err)

	for k := range db {
		delete(db, k)
	}

	fmt.Println(addressBind)

	for _, v := range addressBind {
		db[v.Path] = v.Url
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		showAPI(w)
		return
	}

	mapHandler(db, w, r)
}

func showAPI(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintln(
		w,
		`<h3 style="text-align: center;"><strong>API</strong></h3>
<p><strong>/set_urls</strong><em> | </em><strong>POST</strong><em>, Content-Type: JSON(application/json) or <a href="https://datatracker.ietf.org/doc/rfc9512/">YML (application/yaml)</a></em></p>
<h3><strong>Interface: </strong></h3>
<blockquote>
<p><strong>Array of { path: string, url: string }</strong></p>
</blockquote>`,
	)
}
