FROM golang:latest

LABEL maintainer="DeLV <lede0510@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download -x

COPY . .

COPY --from=itinance/swag /root/swag /usr/local/bin

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -exclude-dir=.git -exclude-dir=docs --build="make build-dev" --command=./main