FROM golang:1.17.7-alpine3.15 AS build
WORKDIR /go/src/project/
COPY . /go/src/project

RUN set -ex &&\
    apk add --no-progress --no-cache \
      gcc \
      musl-dev git

RUN go get -d -v
RUN GOOS=linux GOARCH=amd64 go build -a -v -tags musl -o /bin/project

#FROM alpine:3.15
#COPY --from=build /bin/project /bin/project
#COPY --from=build /go/src/project/.env .

EXPOSE 9999
ENTRYPOINT [ "/bin/project" ]