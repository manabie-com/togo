include .env
export
start:
	go run cmd/monolith/main.go
migration:
	./scripts/make-migration.sh
migrate:
	sql-migrate up -config=dbconfig.yml
migrate-rollback:
	sql-migrate down -config=dbconfig.yml
test:
	go test -v -cover ./...