FROM golang:1.15.3
ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir -p /app
COPY . /app
WORKDIR /app

RUN go mod download && \
    go build -ldflags="-w -s" && \
find . -type f -not -name 'togo' -delete

ENTRYPOINT ["./togo"]
