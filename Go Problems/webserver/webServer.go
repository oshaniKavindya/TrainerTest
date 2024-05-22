package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Set up a file server to serve the static HTML file
	Server := http.FileServer(http.Dir("./"))
	http.Handle("/", Server)

	// Start the server on port 8080
	fmt.Println("Starting server at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
