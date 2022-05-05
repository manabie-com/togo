test:
	CGO_ENABLED=1 go test -v ./...
build:test
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -ldflags="-linkmode external -extldflags '-static'" -o togo ./cmd/togo/
docker-build:build
	docker build -t togo .
run:docker-build
	docker run -p 8080:8080 togo 