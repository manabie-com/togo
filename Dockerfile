FROM golang:alpine AS builder

ENV APP_ENV="docker"
ENV GIN_MODE="release"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN go build -o main .

EXPOSE 3000

CMD ["./main"]

