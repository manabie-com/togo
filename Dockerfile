# staging:
FROM golang:1.16 as build-stage

WORKDIR /app
COPY . /app 

RUN CGO_ENABLED=0 && go build -o app-exe ./cmd/main.go
RUN CGO_ENABLED=0 && go build -o migrate-exe ./cmd/migrate/main.go

# production: 
FROM golang:1.16 as production-stage

COPY --from=build-stage /app/app-exe /bin
COPY --from=build-stage /app/docker.env /bin
COPY --from=build-stage /app/migrate-exe /bin

RUN chmod +x /bin/app-exe
RUN chmod +x /bin/migrate-exe
RUN chmod +x /bin/docker.env

EXPOSE 5050

CMD [ "sh","-c","set -o allexport;. /bin/docker.env; set +o allexport && /bin/migrate-exe && /bin/app-exe" ]
