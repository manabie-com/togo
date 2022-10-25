FROM golang:1.16-alpine
WORKDIR /app

RUN apk update &&\
    apk add libc-dev &&\
    apk add gcc &&\
    apk add make

RUN go install github.com/golang/mock/mockgen@v1.6.0

COPY ./go.mod go.sum ./
RUN go mod download && go mod verify

RUN go get github.com/githubnemo/CompileDaemon

COPY . .

COPY ./entrypoint.sh /entrypoint.sh

ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
RUN chmod +rx /usr/local/bin/wait-for /entrypoint.sh

ENTRYPOINT [ "sh", "/entrypoint.sh" ]