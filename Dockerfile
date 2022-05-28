# Build stage
FROM golang:1.18-alpine as build

WORKDIR /app

COPY . .

RUN go build -o /togo ./cmd/togo


# Deploy stage
FROM alpine
WORKDIR /

COPY configs /configs
COPY scripts/start.sh .
COPY scripts/wait-for.sh .
COPY --from=build /togo /togo

EXPOSE 8080

CMD [ "/togo" ]

ENTRYPOINT [ "/start.sh" ]
