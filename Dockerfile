FROM golang:1.18-alpine AS BUILDER
WORKDIR /app
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main
#CMD ["/app/main"]

FROM alpine:latest
WORKDIR /app
COPY --from=BUILDER /app/main ./
EXPOSE 8080
CMD ["/app/main"]
