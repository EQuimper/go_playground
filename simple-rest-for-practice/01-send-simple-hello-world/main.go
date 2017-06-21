package main

import (
	"io"
	"net/http"
)

// For running this server only run `go run main.go`

func main() {
	// Create a request who response on '/' with the function index
	http.HandleFunc("/", index)
	// Make server listen on port :8080 and put nil for make use of the defaultservermux
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}
