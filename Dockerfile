# Start from a minimal base image with Go 1.21.6 installed
FROM golang:1.21.6-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go application with prod flags by default
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o app ./app

# Start from a smaller base image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the executable from the builder image
COPY --from=builder /app/app .

# Expose the port the app runs on
EXPOSE 8080

# Set the GIN_MODE environment variable to "release" by default
ENV GIN_MODE=release

# Command to run the executable
CMD ["./app"]
