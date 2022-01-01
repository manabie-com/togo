# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git
ADD . /src
RUN cd /src && go build -o goapp

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/. /app/
ENTRYPOINT ./goapp