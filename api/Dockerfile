FROM golang:1.17

ENV ROOT /api

# Create and change to the app directory.
WORKDIR $ROOT

COPY . .

RUN make install
