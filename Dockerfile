FROM golang:alpine
WORKDIR /app
COPY togo .

ENV GO_LOG_STDERR=true
ENTRYPOINT [ "./togo" ]
