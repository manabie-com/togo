FROM golang:1.14-alpine as builder
WORKDIR /src/go

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./internal ./internal
COPY ./cmd ./cmd


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cli ./cmd/cli
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/app

CMD ["./server"]
