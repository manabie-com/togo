# Build stage
FROM golang:1.16-alpine3.13 AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/main.go
RUN apk --no-cache add curl

# Run stage
FROM alpine:3.13
WORKDIR /app
COPY . .
COPY --from=builder /app/main .
COPY start.sh .
COPY wait-for.sh .

EXPOSE 5000
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]