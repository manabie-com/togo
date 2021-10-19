FROM golang:1.14-stretch

WORKDIR /var/www/src

RUN go get github.com/cespare/reflex
COPY reflex.conf /
ENTRYPOINT ["reflex", "-c", "/reflex.conf"]