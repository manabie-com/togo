mock:
	mockgen -package mockdb -destination internal/storages/mock/mock.go github.com/surw/togo/internal/services ILiteDB

run:
	docker-compose up -d --build