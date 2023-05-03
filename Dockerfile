# Use golang:alpine as the base image
FROM golang:alpine as builder

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Install dependencies
RUN apk update && apk add --no-cache git
RUN go get -d -v
RUN go install -v

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Use alpine as the base image
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder image
COPY --from=builder /app/app .

# Expose port 8080
EXPOSE 8080

# Set the entry point for the container
ENTRYPOINT ["/root/app"]
