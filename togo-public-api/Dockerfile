# syntax=docker/dockerfile:1

FROM golang:1.16 as builder

ENV GO111MODULE=on
WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app

FROM alpine:3.14 as final

COPY --from=builder /build/app /app
COPY --from=builder /build/config/dev.toml /config.toml

ENV TZ Asia/Saigon
EXPOSE 9090 8080

ENV PORT 9090
ENV METRIC_PORT 8080
ENV CONFIG_FILE config.toml

ENTRYPOINT ["/app", "serve"]

