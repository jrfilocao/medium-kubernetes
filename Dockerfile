# Use the official Go base image
FROM golang:1.12.0-alpine3.9

# Set the working directory
WORKDIR /app

# Copy the Go source code into the container
COPY server.go .

# Build the Go binary
RUN go build -o server .

# Expose the port that the server listens on
EXPOSE 8080

# Start the server when the container is run
CMD ["./server"]