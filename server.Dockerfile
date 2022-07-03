FROM golang:1.18.1-alpine
WORKDIR /app
COPY . . 
RUN go mod download
RUN go build -o /togo
EXPOSE 8000
CMD [ "/togo" ]