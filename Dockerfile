FROM golang:1.14

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

ENV APP_HOME go/src/togoapp
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

# fetch dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# build app
COPY . .
RUN go build -o togoapp main.go

#
CMD ["./togoapp"]