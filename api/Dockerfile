FROM golang:1.17

ENV ROOT /api

# Create and change to the app directory.
WORKDIR $ROOT

COPY go.mod ./
COPY go.sum ./
COPY . .

RUN make install
