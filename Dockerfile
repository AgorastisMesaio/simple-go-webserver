# Use an official Go runtime as a parent image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Install any required Go dependencies
RUN go mod init example.com/m

# Build the Go application
RUN go build -o main .

# Make port 8080 available to the world outside this container
EXPOSE 8080

# Run the Go app
CMD ["./main"]
