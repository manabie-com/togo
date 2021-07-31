.PHONY: migrate, migrate-down, run, build, compile

migrate:
	goose -dir migrations \
		postgres "user=${DB_USERNAME} dbname=${DB_NAME} password=${DB_PASSWORD} host=${DB_HOST} port=${DB_PORT} sslmode=${SSL_MODE}" up

migrate-down:
	goose -dir migrations \
		postgres "user=${DB_USERNAME} dbname=${DB_NAME} password=${DB_PASSWORD} host=${DB_HOST} port=${DB_PORT} sslmode=${SSL_MODE}" down

build:
	go build -o bin/main main.go

run:
	go run main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
