# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Cao Dinh <bacao1994@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

RUN mkdir -p /usr/src/app/manabie-com/togo

# Set the current working directory inside the container
WORKDIR /usr/src/app/manabie-com/togo

# Copy the source from the current directory to the working Directory inside the container
COPY . .
# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go get ./...

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /usr/src/app/manabie-com/togo/main .
COPY --from=builder /usr/src/app/manabie-com/togo/.env .

# Expose port 8080 to the outside world
EXPOSE 9005

#Command to run the executable
CMD ["./main"]