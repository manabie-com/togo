docker:
	docker build .

run:
	docker-compose up --build -d
	docker-compose exec db pgloader data.db postgres://postgres:password@localhost:5432/postgres?sslmode=disable

test: 
	go test -v -cover -covermode=atomic ./internal/services

stop:
	docker-compose down