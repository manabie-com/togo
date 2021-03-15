FROM golang:1.16.1 as builder

RUN mkdir -p /build

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Run test
# RUN go test ./...

# Build the application
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o main .

RUN mkdir -p /dist
# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

FROM scratch
COPY --from=builder /dist/main /

# Command to run the executable
ENTRYPOINT ["/main"]