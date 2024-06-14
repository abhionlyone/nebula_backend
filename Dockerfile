# Use the latest Go image
FROM golang:latest AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Install air for live reloading
RUN go install github.com/air-verse/air@latest

# Build the Go app
RUN go build -o ./tmp/main cmd/nebula_backend/main.go

# Start a new stage from scratch
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app .

# Copy the .air.toml configuration file
COPY .air.toml .

# Copy the .env file
COPY .env .

# Copy the air binary from the builder stage
COPY --from=builder /go/bin/air /usr/local/bin/air

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable with air
CMD ["air", "-c", ".air.toml"]