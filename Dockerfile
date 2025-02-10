# Use the latest Go 1.23 image
FROM golang:1.23 as builder

WORKDIR /app

# Copy go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source files
COPY . .

# Build the binary explicitly for Linux
RUN GOOS=linux GOARCH=amd64 go build -o mock-api-server

# Final stage: Use a minimal base image
FROM ubuntu:latest

WORKDIR /app

# Install CA certificates for secure HTTPS calls
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

# Copy the compiled binary from the builder
COPY --from=builder /app/mock-api-server .

# Ensure the binary has execute permissions
RUN chmod +x /app/mock-api-server

# Expose API port
EXPOSE 3002

# Run the application
CMD ["/app/mock-api-server"]
