package main

import (
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/fatabbas/ascii-art-web/projects/httpserver/index.html")
}

func main() {
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/home/fatabbas/ascii-art-web/projects/httpserver/style.css")
	})
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":2003", nil)
}
