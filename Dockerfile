FROM golang:1.16-alpine3.14 AS build_base

ENV CGO_ENABLED=1
ENV GO111MODULE=on

RUN apk add --no-cache git gcc g++

# Set the Current Working Directory inside the container
WORKDIR /src

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./bin/server ./cmd/togo

RUN chmod +x bin/server

# Start fresh from a smaller image
FROM alpine:3.14

RUN apk add ca-certificates tzdata

ENV TZ=Asia/Ho_Chi_Minh

COPY --from=build_base /src/bin/server /app/server

ENTRYPOINT ["/app/server"]