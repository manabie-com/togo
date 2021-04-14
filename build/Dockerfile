FROM golang:alpine

RUN apk add --no-cache git

RUN apk add bash

# All these steps will be cached
WORKDIR /togo
# COPY go.mod and go.sum files to the workspace
COPY go.mod go.sum ./
# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .
# This step is required to create dir vendor
RUN go mod vendor
# Build the binary
RUN CGO_ENABLED=0 \
GOOS=linux \
GOARCH=amd64 \
go build -mod=vendor -o cli.bin cmd/main.go