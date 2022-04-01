### Reference: https://hub.docker.com/_/golang

FROM golang:1.18 as builder

WORKDIR /go/src/github.com/TrinhTrungDung/togo

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -v -o build/server ./cmd/api/main.go

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/TrinhTrungDung/togo/build/server /usr/bin/server
EXPOSE 8080 8080
ENTRYPOINT [ "/usr/bin/server" ]