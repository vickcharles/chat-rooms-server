# Use the official Go image as the base image
FROM golang:1.17 as builder

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build cmd/main.go

# Create the final image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/main .

# Expose the application's port
EXPOSE 80

# Run the Go binary
CMD ["./main"]
