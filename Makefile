build:
	docker-compose build

serve:
	docker-compose up

down:
	docker-compose down

run_test:
	docker-compose run --rm api go test /api/api

# database commands

sql_generate:
	docker-compose exec api sqlc generate

create_migration:
	docker-compose run --rm migrate create -ext sql -dir /migrations $(name)
	sudo chown -R "$(id -u):$(id -g)" db/migrations

migrate_up:
	docker-compose run --rm migrate -path /migrations -database postgres://postgres:postgres@db:5432/todo_app?sslmode=disable up

migrate_down:
	docker-compose run --rm migrate -path /migrations -database postgres://postgres:postgres@db:5432/todo_app?sslmode=disable down

