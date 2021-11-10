MIGRATE_TAG = init

start-docker-compose:
	docker-compose -p togo up -d

create-migration:
	migrate create -ext -sql -dir migrations -seq ${MIGRATE_TAG}

migrate-up:
	migrate -path migrations -database ${DB__SOURCE} -verbose up

migrate-down:
	migrate -path migrations -database ${DB__SOURCE} -verbose down 1

test:
	go test ./...  -count=1 -v -cover

run:
	go run cmd/main.go

lint:
	golangci-lint run --allow-parallel-runners

generate-mocks:
	mockery --all

.PHONY:start-docker-compose create-migration migrate-up migrate-down lint generate-mocks