# Use the official Go image as a base
FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV GIN_MODE=release
ENV CGO_ENABLED=0

# Build
RUN go build -o main main.go

# Create new image and import just the binary
FROM alpine:3.16

WORKDIR /root/

COPY --from=builder /app/main .

ENV GIN_MODE=release
ENV CGO_ENABLED=0

EXPOSE 80

CMD ["./main"]