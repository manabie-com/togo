.PHONY: install

MIGRATE_TODO := migrate -source file://migrate/ddl  -database 'postgres://postgres:password@postgresql.manabie.todo:5432/todo?sslmode=disable'

install:
	go mod tidy
	go install -tags "postgres" github.com/golang-migrate/migrate/v4/cmd/migrate

run:
	go run main.go

format:
	go fmt ./...

test:
	go test -count=1 -cover ./...

test-e2e:
	go test -count=1 -tags=e2e_test ./api/...

migrate-todo:
	@echo "Migrating todo DB.."
	@${MIGRATE_TODO} up
	@echo "Done!"

down-todo:
	@echo "Migrating(down) todo DB.."
	@echo y | $(MIGRATE_TODO) down 1
	@echo "Done!"

drop-todo:
	@echo "Dropping todo DB.."
	@echo y | $(MIGRATE_TODO) drop
	@echo "Dropped!"
