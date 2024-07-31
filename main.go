// main.go
// Simple Go Web Server
//
// This program implements a simple web server that responds with an HTML page containing a "Hello, World!" message
// and a button that links to your GitHub account.
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

// helloHandler handles HTTP requests and responds with an HTML page containing a "Hello, World!" message
// and a button linking to your GitHub profile.
func helloHandler(w http.ResponseWriter, r *http.Request) {
    htmlContent := `
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Hello, World!</title>
        <style>
            body {
                font-family: Arial, sans-serif;
                display: flex;
                flex-direction: column;
                align-items: center;
                justify-content: center;
                height: 100vh;
                margin: 0;
                background-color: #f4f4f4;
            }
            h1 {
                color: #333;
            }
            .button {
                padding: 10px 20px;
                font-size: 16px;
                color: white;
                background-color: #007bff;
                border: none;
                border-radius: 5px;
                text-decoration: none;
                cursor: pointer;
                transition: background-color 0.3s ease;
            }
            .button:hover {
                background-color: #0056b3;
            }
        </style>
    </head>
    <body>
        <h1>Hello, World!</h1>
        <a href="https://github.com/AgorastisMesaio" class="button">Visit My GitHub</a>
    </body>
    </html>`
    
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, htmlContent)
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
