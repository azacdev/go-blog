# Stage 1: Build stage
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

# Install git (only for downloading modules if needed)
RUN apk add --no-cache git

# Copy go mod files and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy full source
COPY . .

# Build the Go binary
RUN go build -o app main.go

# Stage 2: Production stage
FROM alpine:latest

WORKDIR /app

# Copy only the built binary from builder stage
COPY --from=builder /app/app .

# Copy config files
COPY config config

# Copy env file into image
COPY .env .

# Expose app port
EXPOSE 8080

# Run the binary directly
ENTRYPOINT ["./app"]
CMD ["serve"]
