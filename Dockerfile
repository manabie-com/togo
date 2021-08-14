FROM golang:1.14.4 AS build
LABEL maintainer="thuocnv1802@gmail.com"

# Switches to /tmp/app as the working directory, similar to 'cd'
WORKDIR /build/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Builds the current project to a binary file.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o task-service .


FROM ubuntu:18.04

ENV APP_DIR /home/task/app

RUN apt-get update \
    && apt-get install -y ca-certificates && update-ca-certificates \
    && apt-get install -y tzdata xz-utils  unzip libaio1 \
    && mkdir -p "${APP_DIR}"

# Copies the binary file from the BUILD container to /app folder
COPY --from=build /build/app/task-service $APP_DIR/task-service

# Switches working directory to /app
WORKDIR $APP_DIR

# Exposes the 8088 port from the container
EXPOSE 5050

# Runs the binary once the container starts
ENTRYPOINT ["./task-service"]