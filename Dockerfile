# Use the official Go image as a base
FROM golang:1.17-alpine as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the entire project to the container
COPY . .

# Build the user
RUN go build -o user ./srv/user

# Build the spelling microservice
RUN go build -o spelling ./srv/spelling

# Use a minimal base image for the final container
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the built binaries from the builder stage
COPY --from=builder /app/user .
COPY --from=builder /app/spelling .

# Expose ports for both microservices
EXPOSE 8080
EXPOSE 8081

# Command to run the microservices
CMD ["./user","./spelling"]
