# Use the official Go image as a base
FROM golang:1.22 as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy the entire project to the container
COPY . .

# Build the user
RUN go build -o main init.go

# Use a minimal base image for the final container
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /root/

COPY --from=builder /app/main .

# Expose ports for both microservices
EXPOSE 80

# Command to run the microservices
ENTRYPOINT ["/root/main"]