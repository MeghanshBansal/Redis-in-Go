# Stage 1: Build the Go application
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server .

# Stage 2: Create the runtime container
FROM scratch

# Copy the compiled Go application from the builder stage
COPY --from=builder /app/server /server

# Expose the application port
EXPOSE 8080

# Command to run the Go server
ENTRYPOINT ["/server"]
