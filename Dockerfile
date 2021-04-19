FROM golang:1.16.3-alpine3.13 as builder

COPY go.mod go.sum /go/src/github.com/manabie-com/togo/
WORKDIR /go/src/github.com/manabie-com/togo
RUN go mod download
COPY . /go/src/github.com/manabie-com/togo

# RUN echo "ipv6" >> /etc/modules

RUN apk add --no-cache build-base ca-certificates && update-ca-certificates

RUN GOOS=linux go build -a -installsuffix cgo -o build/manabie_togo github.com/manabie-com/togo

FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/manabie-com/togo/build/manabie_togo /usr/bin/manabie_togo

EXPOSE 5050 5050

ENTRYPOINT ["/usr/bin/manabie_togo"]