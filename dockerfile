# build phase
from golang:1.14-alpine as builder
run apk add --no-cache git make build-base
workdir /app
copy go.mod .
copy go.sum .
run go mod download
copy . .
run GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o togo


# final phase
from alpine
workdir /app
copy --from=builder /app/togo .
copy --from=builder /app/data.db .
expose 8888
entrypoint ["./togo"]
