FROM golang:1.13 as builder

ENV GO111MODULE=on
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM scratch

COPY --from=builder /app/togo /app/

EXPOSE 5050
ENTRYPOINT ["/app/togo"]