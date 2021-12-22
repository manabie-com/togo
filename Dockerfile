ARG GO_VERSION=1.17

FROM golang:${GO_VERSION}-alpine AS builder

RUN mkdir -p /togo
WORKDIR /togo

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./server ./cmd/api/main.go

EXPOSE 3000

ENTRYPOINT ["./server"]
