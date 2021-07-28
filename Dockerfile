# build stage
FROM golang:1.15-alpine AS build

ENV APP togo
ENV GO111MODULE=on

RUN apk add git

COPY . /go/src/$APP/
WORKDIR /go/src/$APP

RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/$APP main.go

###
FROM golang:1.15-alpine
ENV APP togo

RUN apk add git; \
    apk add build-base; \
    apk add --no-cache tzdata; \
    go get -u github.com/pressly/goose/cmd/goose

COPY ./internal/database/migrations/ /internal/database/migrations/
COPY ./entrypoint.sh /entrypoint.sh
COPY --from=build /go/bin/$APP /go/bin/$APP

EXPOSE 5050
ENTRYPOINT ["sh","/entrypoint.sh"]