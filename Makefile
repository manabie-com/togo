startDB:
	docker-compose -f docker-compose.yml up -d --build postgres_db

seedDB:
	go run ./internal/seed/seed.go

test:
	go test -v -cover -covermode=atomic ./internal/module/
	go test -v -cover -covermode=atomic ./internal/module/user
	go test -v -cover -covermode=atomic ./internal/module/task

stop:
	docker-compose down