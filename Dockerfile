
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN apk add --no-cache git build-base gcc abuild binutils binutils-doc gcc-doc
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/togo /togo
ENTRYPOINT ./togo
LABEL Name=togo Version=0.0.1
EXPOSE 5050
