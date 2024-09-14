# Stage 1: Build
FROM golang:1.23.1-alpine3.19 AS builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o backend ./cmd/api/main.go

# Stage 2: Run
FROM alpine:3.19

# Copy the binary from the builder stage
COPY --from=builder /app/backend /usr/local/bin/backend

# Expose port
EXPOSE 8000

# Set the entry point
ENTRYPOINT [ "/usr/local/bin/backend" ]
