FROM golang:1.19

WORKDIR /app/manabie

COPY go.* ./

RUN go mod tidy

COPY . ./

RUN cd /app/manabie/cmd/ && go build -o /server

EXPOSE 9000

CMD ["/server"]