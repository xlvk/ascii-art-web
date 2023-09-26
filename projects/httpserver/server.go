package main

import (
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":2003", nil)
}
