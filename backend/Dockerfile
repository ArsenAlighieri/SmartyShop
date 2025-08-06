# Stage 1: Builder
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
# -o /app/server: output file name
# CGO_ENABLED=0: disable CGO for static linking
# GOOS=linux: specify the target OS
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/main.go

# Stage 2: Runtime
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/server .

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["./server"]
