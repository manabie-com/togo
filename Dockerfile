FROM golang:1.16.4-alpine3.13 as builder
WORKDIR /app
RUN apk update && apk upgrade  -U -a && \
    apk add bash git openssh gcc libc-dev
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o /app/todo /app


FROM golang:1.16.4-alpine3.13
RUN apk add --update ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime && \
    echo "Asia/Ho_Chi_Minh" > /etc/timezone && \
    rm -rf /var/cache/apk/*
COPY --from=builder /app/todo /app/todo

## Add the wait script to the image
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.8.0/wait /wait
RUN chmod +x /wait

WORKDIR /app
CMD /wait && /app/todo
