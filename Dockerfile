
FROM golang:1.16.4

WORKDIR /go/src/
COPY . .

RUN go get -d -v ./...
RUN go build

CMD ["./togo"]
EXPOSE 5050