PLATFORM=local
export DOCKER_BUILDKIT=1


# $(shell git rev-parse --short HEAD)
VERSION := 1.0 

all: todo-api

test-auth:
	curl -il 'http://localhost:3000/v1/users/auth' \
		-H 'Content-Type: application/json' \
		--data '{ "email": "user@example.com", "password": "gophers" }'

test-create-todo:
	curl -il 'http://localhost:3000/v1/todos' \
		-H 'Content-Type: application/json' \
		-H "Authorization: Bearer ${TOKEN}" \
		--data '{ "title": "test title", "content": "test content" }'

todo-api:
	@docker build \
		--ssh default=${SSH_AUTH_SOCK} \
		--platform ${PLATFORM} \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		--target bin \
		--tag togo-todo-api:$(VERSION) \
		--tag togo-todo-api \
		.

run:
	docker compose up

restart: todo-api recreate

recreate:
	docker compose up -d todo-api --force-recreate

tidy:
	go mod tidy

db:
	docker compose up -d cloudbeaver
	@echo 'http://localhost:8978/'

# Admin
migrate:
	go run cmd/todo-admin/main.go migrate

seed:
	go run cmd/todo-admin/main.go seed


test:
	go test -count=1 ./...

