![GitHub action workflow status](https://github.com/AgorastisMesaio/simple-go-webserver/actions/workflows/docker-publish.yml/badge.svg)

# Simple Go Web Server

This is a simple web server written in Go. The purpose of this project is to demonstrate how to create a Docker container for a Go application.

## Getting Started

### Prerequisites

- Docker
- Go (for local development)

### Running the Application

#### Using Docker

1. Build the Docker image:

    ```sh
    docker build -t simple-go-webserver .
    ```

2. Run the Docker container:

    ```sh
    docker run -p 8080:8080 simple-go-webserver
    ```

3. Open your browser and go to `http://localhost:8080` to see the application running.

#### Local Development

1. Install Go from the official website: https://golang.org/dl/

2. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/simple-go-webserver.git
    cd simple-go-webserver
    ```

3. Run the application:

    ```sh
    go run main.go
    ```

4. Open your browser and go to `http://localhost:8080` to see the application running.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
