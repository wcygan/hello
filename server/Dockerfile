# syntax=docker/dockerfile:1

# Build Stage
FROM golang:1.24-alpine AS builder

# Set necessary environment variables
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

# Copy go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
# Copy cmd directory specifically as that contains main.go
COPY cmd/ ./cmd/

# Build the application
# Output the binary to /server (this will be copied to the final stage)
RUN go build -o /server ./cmd/main.go

# Final Stage
FROM alpine:latest

# Copy the static binary from the builder stage
COPY --from=builder /server /server

# Expose port 8080 (as used in main.go)
EXPOSE 8080

# Command to run the executable
CMD ["/server"] 