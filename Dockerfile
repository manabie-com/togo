## Build
FROM golang:1.17.2-alpine AS builder

ARG APP_NAME=manabie-test

WORKDIR /manabie

COPY . .

RUN CGO_ENABLED=0 go build -o ./$APP_NAME

# Run
FROM alpine:3.14
WORKDIR /manabie
COPY --from=builder /manabie/$APP_NAME .
RUN apk add --no-cache tzdata
RUN cp /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime
EXPOSE 8080

CMD /manabie/manabie-test