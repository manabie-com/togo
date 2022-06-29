run:
	go run cmd/server/main.go
mock:
	mockgen -source=internal/domain/domain.go   -destination=./pkg/mock/mock.go -package=mockdb