# Stage 1: Build Stage for Dependencies
FROM golang:1.23-alpine AS builder

# Set working directory to /app
WORKDIR /app

# Copy Go module files (go.mod and go.sum)
COPY server/go.mod server/go.sum /app/

# Install Go dependencies
RUN go mod tidy

# Copy Go source files
COPY server/ /app/

# Stage 2: Build the Go application
RUN go build -o /app/bin/server /app/main.go  # Change the path to your main file if needed

# Stage 3: Runtime Stage
FROM alpine:latest

# Set working directory to /app
WORKDIR /app

# Copy the Go binary from the build stage
COPY --from=builder /app/bin/server /app/

# Expose the application port
EXPOSE 4000

# Set the default command to run the Go application
CMD ["/app/server"]
