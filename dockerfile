# Build stage
FROM golang:1.22.1-alpine3.19 as build

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker layer caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . ./

# Build the Go application
RUN go build -o api ./cmd/api

# Run stage
FROM alpine:3.19

WORKDIR /app

# Copy the compiled binary from the build stage
COPY --from=build /app/api /app/api

# Copy environment file
COPY .env .

# Expose the application port
EXPOSE 3000

# Command to run the application
CMD ["./api"]
