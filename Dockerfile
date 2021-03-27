# Load golang image to build
FROM golang:1.14 as builder
ARG APP_NAME
ARG APP_PATH

ENV APP_USER app
ENV APP_HOME /go/src/app-builder

RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME

WORKDIR $APP_HOME
USER $APP_USER
COPY ./ .

RUN go build -o $APP_NAME $APP_PATH


# Deploy execute file to simple linux server
FROM debian:buster
ARG APP_NAME
ARG APP_PORT

ENV APP_USER app
ENV BUILDER_APP_HOME /go/src/app-builder
ENV APP_PORT $APP_PORT

RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p /app
WORKDIR /app

#RUN echo "Copy app from Build PATH: $BUILDER_APP_HOME/$APP_NAME"
COPY --chown=0:0 --from=builder $BUILDER_APP_HOME/$APP_NAME ./runner

RUN echo "APP_NAME: $APP_NAME"

EXPOSE $APP_PORT
USER $APP_USER
CMD ./runner