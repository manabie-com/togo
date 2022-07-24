# syntax = docker/dockerfile:1-experimental
FROM golang:1.18-alpine3.15 AS base

RUN apk --no-cache --update-cache --upgrade --latest add build-base git gcc bash

WORKDIR /go/src/github.com/trangmaiq/togo

ADD go.mod go.mod
ADD go.sum go.sum

ENV GO111MODULE on
RUN go mod download

ADD . .

ARG VERSION
ARG COMMIT
ARG BUILD_DATE

ENV GOBIN=/usr/bin

RUN --mount=type=cache,target=/root/.cache/go-build go build \
    -ldflags="-X 'github.com/trangmaiq/togo/build.Version=${VERSION}' -X 'github.com/trangmaiq/togo/build.Date=${BUILD_DATE}' -X 'github.com/trangmaiq/togo/build.Commit=${COMMIT}'" \
    -o /usr/bin/togo

FROM alpine:3.15.4

RUN addgroup -S togo; \
    adduser -S togo -G togo -D -u 10000 -h /home/togo -s /bin/nologin; \
    chown -R togo:togo /home/togo

COPY --from=base /usr/bin/togo /usr/bin/togo
EXPOSE 9000
USER 10000

ENTRYPOINT ["togo"]
