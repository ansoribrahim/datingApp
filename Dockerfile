# Use an official Go image as a base image
FROM golang:1.20 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files for dependency management
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Use a minimal base image for the final container
FROM debian:bullseye-slim

# Set the working directory inside the container
WORKDIR /app

# Copy the built application binary from the builder stage
COPY --from=builder /app/main .

# Expose the application's port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
