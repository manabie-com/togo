FROM golang:1.14 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
#RUN go build -o .

FROM alpine:3.13
WORKDIR /build
COPY --from=builder /build/main ./
CMD ["./main"]