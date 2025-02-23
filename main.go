package main

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type AddressBind struct {
	Path string `gorm:"uniqueIndex"`
	Url  string `gorm:"unique"`
}

var db *gorm.DB

func main() {
	db = getDB()
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

func setUrlsHandler(w http.ResponseWriter, r *http.Request) {
	var addressBinds []AddressBind

	err := parseFormat(w, r, &addressBinds)

	if err != nil {
		fmt.Fprintln(w, err.Error())
	}

	cleanTable(db)
	updateTable(db, addressBinds)
}

func dbProvider(value string) (string, error) {
	return getValue(db, value)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		showAPI(w)
		return
	}

	mapHandler(dbProvider, w, r)
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
