package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	_ "embed"
)

func main() {
	port := "0.0.0.0:8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	http.HandleFunc("/", ServeHTTP)

	log.Printf("Listening on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

//go:embed index.html
var indexHtml string

const uCDN_JSDELIVER_NET = "https://cdn.jsdelivr.net"

var startup = time.Now() // time the server started

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// index: server index html
	if r.URL.Path == "" || r.URL.Path == "/" {
		http.ServeContent(w, r, ".html", startup, strings.NewReader(indexHtml))
		return
	}

	// everything else: do a redirect
	dest := uCDN_JSDELIVER_NET + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
		dest += "?" + r.URL.RawQuery
	}
	http.Redirect(w, r, dest, http.StatusTemporaryRedirect)
	return
}
