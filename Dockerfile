# Use the official Go image as the base image
FROM golang:1.23.1 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o pos-backend .

# Use a smaller base image for the final image
FROM debian:bookworm-slim

# Install ca-certificates for HTTPS requests
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/pos-backend .

# Expose the port your application listens on (adjust if needed)
EXPOSE 8080

# Command to run the application
CMD ["./pos-backend"]
