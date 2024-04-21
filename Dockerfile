# Use the official Go image as a base
FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build
RUN go build -o main main.go

EXPOSE 8080

# Run
CMD ["./main"]