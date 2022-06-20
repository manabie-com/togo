# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

ENV HOME /app
ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /app

# Download all the dependencies
COPY go.mod go.sum ./
RUN go mod download
COPY . .

EXPOSE 8080
RUN go build -o togo
CMD [ "/app/togo" ]