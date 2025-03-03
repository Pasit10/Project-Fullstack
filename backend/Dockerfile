FROM golang:1.24.0 AS builder

# Set the working directory inside the container
WORKDIR /backend

# Copy the Go modules manifest and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Use a minimal base image for running the application
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port (change if needed)
EXPOSE 8080

# Run the application
CMD ["./main"]
