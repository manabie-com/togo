FROM golang:1.16-alpine

WORKDIR /app

COPY ./internal/go.mod ./
RUN go mod tidy

COPY ./internal ./

RUN go build -o /togo

EXPOSE 5050

CMD [ "/togo" ]