package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "HTTP 404: File Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "HTTP Method not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello Request Handler")
}

func main() {
	fileServer := http.FileServer(http.Dir("/public"))
	//
	// requwst handler
	//
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloRequestHandler)

	fmt.Printf("Starting server on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
