package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I can do evertyhing!\n\n")
	fmt.Fprintf(w, "Release: v6.0.0")
}

func main() {
	fmt.Println("Basic web server is starting on port 8080...")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
