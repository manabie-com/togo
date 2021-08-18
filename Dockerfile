FROM golang:1.17-buster

ENV GROUP_ID=1000
ENV USER_ID=1000
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV APP_USER app
ENV APP_HOME /app

WORKDIR $APP_HOME

# Install dependencies
COPY ./internal/go.mod .
COPY ./internal/go.sum .
RUN go mod tidy

# Copy app source
COPY ./internal ${APP_HOME}

# Run as non root user
RUN groupadd --gid $GROUP_ID app && useradd -m -l --uid $USER_ID --gid $GROUP_ID $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME
USER $APP_USER

EXPOSE 5050

CMD ["go", "run", "main.go"]