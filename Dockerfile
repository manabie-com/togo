FROM golang:1.17 as build-stage

WORKDIR /app
COPY . /app

RUN CGO_ENABLED=0 && go build -o app-exe ./cmd/app/...

FROM  golang:1.17 as production-stage

COPY --from=build-stage /app/app-exe /bin
COPY --from=build-stage /app/docker.env /bin

RUN chmod +x /bin/app-exe

EXPOSE 4000

CMD [ "sh","-c","set -o allexport;. /bin/docker.env; set +o allexport && /bin/app-exe" ]