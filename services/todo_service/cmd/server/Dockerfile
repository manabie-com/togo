FROM golang:1.19.1-alpine as builder
WORKDIR /app
RUN apk update && apk add --no-cache gcc musl-dev git
COPY . .
WORKDIR /app/services/todo_service
RUN go mod download
RUN go build -ldflags '-w -s' -a -o togo ./cmd/server/main.go

# Deployment environment
# ----------------------
FROM alpine:3.17
WORKDIR /app
RUN apk update
RUN chown nobody:nobody /app
USER nobody:nobody
COPY --from=builder --chown=nobody:nobody ./app/services/todo_service/togo .
COPY --from=builder --chown=nobody:nobody ./app/services/todo_service/cmd/server/run.sh .

ENTRYPOINT sh run.sh

