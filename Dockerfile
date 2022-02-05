FROM golang:1.15.2

# Install the air binary so we get live code-reloading when we save files
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /api

COPY go.mod .
COPY go.sum .

# pull go dependencies
RUN go mod download

COPY . ./

CMD ["air"]