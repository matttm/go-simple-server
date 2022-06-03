package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//
	// requwst handler
	//
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Request Handler")
	})

	fmt.Printf("Starting server on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
