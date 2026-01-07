# Stage 1: Build the Go binary
FROM golang:1.25-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Run tests (fail build if tests fail)

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /go_book_api ./cmd/main.go

# Stage 2: Runtime image
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /go_book_api .
# Copy .env file if you want it baked in, 
# though it's better to pass env via docker-compose
COPY .env .

# Expose the port the app runs on
EXPOSE 8000

# Command to run the executable
CMD ["./go_book_api"]