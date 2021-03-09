FROM golang:1.12 AS builder
WORKDIR /usr/togo
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.10
WORKDIR /usr/togo
COPY --from=builder /usr/togo/main ./
CMD ["./main"]