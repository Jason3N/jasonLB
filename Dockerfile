# Use the official Golang image as the build stage
FROM golang:1.20-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod .

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o go-load-balancer .

# Start a new stage from scratch
FROM alpine:latest

# Copy the pre-built binary file from the builder stage
COPY --from=builder /app/go-load-balancer /go-load-balancer

# Expose the load balancer port
EXPOSE 8080

# Command to run the executable
CMD ["/go-load-balancer"]
