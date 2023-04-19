# syntax = docker/dockerfile:1

FROM --platform=${BUILDPLATFORM} golang:1.19-alpine as base

RUN apk --update add tzdata \
                     ca-certificates \
                     git \
                     openssh-client \
                     build-base

WORKDIR /src

COPY go.* .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=ssh \
    go mod download

FROM base AS build

ARG TARGETOS
ARG TARGETARCH
ARG BUILD_DATE
ARG BUILD_REF

RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -tags musl -ldflags "-X main.build=${BUILD_REF} -extldflags -static" -o /app/todo-api cmd/todo-api/main.go

FROM scratch AS bin-unix

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/todo-api /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM bin-${TARGETOS} as bin

ENTRYPOINT [ "./todo-api" ]

# vim: ft=dockerfile
