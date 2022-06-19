FROM golang:1.17-alpine as builder
WORKDIR /build
ENV GO111MODULE=on
RUN apk add git &&\
    go get github.com/cespare/reflex
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app cmd/monolith/main.go