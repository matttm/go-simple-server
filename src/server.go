package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
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

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Form submitted\n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "name %s\n", name)
	fmt.Fprintf(w, "address %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	//
	// requwst handler
	//
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
