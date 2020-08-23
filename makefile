.PHONY: test dev remove-infras init
SRC_PATH=$(GOPATH)/src/github.com/phuwn/togo
POSTGRES_CONTAINER?=togo_db
POSTGRES_TEST_CONTAINER?=togo_test_db

dev:
	@GO111MODULE=on RUN_MODE=local go run main.go

test:
	go test -p 1 ./internal/services/...
	go test -p 1 ./integration/...

remove-infras:
	docker-compose stop; docker-compose  rm -f

seed-db-local:
	@docker cp data/seed/. $(POSTGRES_CONTAINER):/
	@docker exec -t $(POSTGRES_CONTAINER) sh -c "chmod +x seed.sh;./seed.sh"

migrate-testing-environment:
	@docker cp integration/migration.sql $(POSTGRES_TEST_CONTAINER):/
	@docker exec -t $(POSTGRES_TEST_CONTAINER) sh -c "PGPASSWORD=password psql -U admin -d togo_test -f /migration.sql"

init: remove-infras
	@docker-compose  up -d 
	@echo "Waiting for database connection..."
	@while ! docker exec $(POSTGRES_CONTAINER) pg_isready -h localhost -p 5432 > /dev/null; do \
		sleep 1; \
	done
	sql-migrate up -config=dbconfig.yml -env="local"
	make seed-db-local
	sql-migrate up -config=dbconfig.yml -env="local_test"
	make migrate-testing-environment