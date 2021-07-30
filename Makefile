.PHONY: migrate, run

migrate:
	goose -dir migrations \
		postgres "user=${DB_USERNAME} dbname=${DB_NAME} password=${DB_PASSWORD} host=${DB_HOST} port=${DB_PORT} sslmode=disable" up

run:
	go run main.go
