# Stage 1: Build the Go app with Air
FROM golang:1.23-alpine AS builder

# Set the working directory in the container
WORKDIR /app

# Install Air for live reloading (updated to the new repo)
RUN go install github.com/air-verse/air@latest

# Copy Go module files and install dependencies
COPY server/go.mod /app/server/go.sum ./
RUN go mod tidy

# Copy the Go source code to the container
COPY server/ /app/

# Stage 2: Run the Go app with Air
CMD ["air"]
