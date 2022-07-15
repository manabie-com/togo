FROM golang:1.16-buster AS build

WORKDIR /go/src/datshiro/togo

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/server ./cmd/server


###

FROM alpine:3.9

COPY --from=0 /usr/local/bin/server /usr/local/bin/server
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["server"]

