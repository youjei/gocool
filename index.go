package main

import (
	"fmt"
	"net/http"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hegggggggllo %s!</h1>", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080", nil)
}
