FROM golang:1.17.1

WORKDIR /var/www/src

RUN go get github.com/cespare/reflex
COPY reflex.conf /

COPY start.sh /

ENTRYPOINT ["reflex", "-c", "/reflex.conf"]
