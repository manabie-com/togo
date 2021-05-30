# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang alpine base image
FROM golang:alpine as builder

# Add Maintainer Info
LABEL maintainer="Pham Hoang Nam<nampham.today@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /go/src/togo

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download dependencies
RUN go get -d -v ./...

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/togo .


######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

#WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/bin .

EXPOSE 5050

CMD ["/togo"]