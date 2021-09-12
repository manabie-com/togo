FROM alpine:3.14
RUN apk add --no-cache ca-certificates
COPY togo .
EXPOSE 5050
ENTRYPOINT [ "./togo" ]
