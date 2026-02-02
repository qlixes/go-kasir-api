FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env.dev .env

RUN go build -o main cmd/main.go

EXPOSE 8080
CMD ["./main"]
