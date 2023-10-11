package main

import (
	"net/http"
	"os"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	// Save path links in var.
	filePath := "/home/fatabbas/ascii-art-web/projects/httpserver/index.html"
	error404 := "/home/fatabbas/ascii-art-web/projects/httpserver/404Error.html"

	// If the path link not found.
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, error404)
		return
	}

	http.ServeFile(w, r, filePath)

	//http.ServeFile(w, r, "/home/fatabbas/ascii-art-web/projects/httpserver/index.html")
}

func main() {
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/home/fatabbas/ascii-art-web/projects/httpserver/style.css")
	})
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":2003", nil)
}
