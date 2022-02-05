
FROM golang:1.17

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o bin/ all 


EXPOSE 8000

CMD [ "bin/json_handler" ]
