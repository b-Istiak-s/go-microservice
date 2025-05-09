# Dockerfile

# Stage 1: Build the app
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code of auth-service
COPY . ./auth-service/

# Set the working directory to where the main.go is located
WORKDIR /app/auth-service/cmd/myapp

# Build the Go app with CGO disabled for a fully static binary
RUN CGO_ENABLED=0 go build -o /auth-service/myapp .

# Stage 2: Create the final lightweight image
FROM alpine:3.17

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /auth-service/myapp .

# Expose the port your app will run on
EXPOSE 8080

# Command to run the app
CMD ["./myapp"]
