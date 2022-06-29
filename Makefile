run:
	go run cmd/server/main.go
mock:
	mockgen -source=internal/domain/domain.go   -destination=./pkg/mock/mock.go -package=mock
unit_test:
	go test -v -cover -short ./...