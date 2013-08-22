package main

import (
	"os"
	"log"
	"net/http"
	"path/filepath"
)

var webappPath string = os.Getenv("ROAMAOR_WEBAPP")
var indexPage string = filepath.Join(webappPath, "static/index.html")

func appHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, indexPage)
}

func assetsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(webappPath, r.URL.Path))
}

func main() {
	_, err := os.Stat(indexPage)
	if err != nil {
		log.Fatal("Could not find index.html")
	}
	http.HandleFunc("/app", appHandler)
	http.HandleFunc("/static/", assetsHandler)
	http.ListenAndServe(":8080", nil)
}
