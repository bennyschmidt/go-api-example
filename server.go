package main

// Dependencies

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

// Main

func main() {

	// Ping

	http.HandleFunc("/", func(writer http.ResponseWriter, reader *http.Request) {
		fmt.Fprintf(writer, "Ping!")
	})

	// Hello

	http.HandleFunc("/hello", func(writer http.ResponseWriter, reader *http.Request) {
		fmt.Fprintf(writer, "Hello! The route is %q :)", html.EscapeString(reader.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
