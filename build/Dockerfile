FROM golang:1.16-alpine

RUN apk add --no-cache ca-certificates musl-dev gcc build-base tzdata && update-ca-certificates

RUN go get -u github.com/cosmtrek/air

WORKDIR /app

ENV PATH $PATH:/go/bin

EXPOSE 3000

ENTRYPOINT ["air", "-c", ".air.toml"]