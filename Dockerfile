# Use the official Golang image as a parent image
FROM golang:alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project directory to the container
COPY . .

# Fetch dependencies and build the binary
RUN go mod download && \
    go build -o main .

# Create a new image from scratch
FROM scratch

# Copy the binary and any other files needed from the builder image
COPY --from=builder /app/main /app/main

# Set the working directory inside the container
WORKDIR /app

# Run the binary
ENTRYPOINT ["/app/main"]


