FROM golang:1.15.14 AS build
WORKDIR /go/src/github.com/surw

COPY go.mod ./
COPY go.sum ./

RUN go mod tidy
RUN go mod download

COPY . ./
RUN go build -o /togo


FROM gcr.io/distroless/base-debian10
WORKDIR /

COPY --from=build /togo /togo
EXPOSE 5050
USER nonroot:nonroot
ENTRYPOINT ["/togo"]

