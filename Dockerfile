ARG DOCKER_IMAGE_TAG="alpine"

FROM golang:${DOCKER_IMAGE_TAG} AS builder
WORKDIR /usr/togo
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -o "./server" "./cmd/server.go"

# ===== Production =====
FROM golang:${DOCKER_IMAGE_TAG}
WORKDIR /usr/togo
COPY --from=builder /usr/togo/server ./
CMD ["./server"]