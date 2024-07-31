// main.go
// Simple Go Web Server
//
// This program implements a simple web server that responds with a "Hello, World!" message.
//
// Usage:
//
//     go run main.go
//
// Or build and run the Docker container:
//
//     docker build -t simple-go-webserver .
//     docker run -p 8080:8080 simple-go-webserver
//
// Visit http://localhost:8080 in your browser to see the output.

package main

import (
    "fmt"
    "log"
    "net/http"
)

// helloHandler handles HTTP requests and responds with a "Hello, World!" message.
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

// main is the entry point for the application.
func main() {
    // Register the helloHandler function for the root URL path.
    http.HandleFunc("/", helloHandler)

    // Start the server on port 8080 and log any errors.
    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("could not start server: %s\n", err)
    }
}
