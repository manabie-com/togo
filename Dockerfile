FROM golang:1.16.3-buster as builder
ARG GIT_USER
ARG GIT_TOKEN
RUN git config --global url."https://${GIT_USER}:${GIT_TOKEN}@github.com".insteadOf "https://github.com"
WORKDIR /app
ENV GOSUMDB off
COPY ./go.mod ./go.sum ./

RUN go mod download
COPY ./ ./
RUN go build /app/cmd/server


FROM gcr.io/distroless/base-debian10

COPY --from=builder /app/server /app/server

WORKDIR /app
CMD ["/app/server"]
