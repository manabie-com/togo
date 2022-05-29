
.PHONY: test integrated_test test_coverage migrate_up migrate_down setup_test_env remove_test_env

#migrate_up:
#
#	migrate -source file://deployment/migrations -database "mysql://root:123456aA@(localhost:3309)/todo?multiStatements=true" -verbose up
#
#migrate_down:
#	migrate -source file://deployment/migrations -database "mysql://root:123456aA@(localhost:3309)/todo?multiStatements=true" -verbose down

setup_env:
	docker-compose -f ./deployment/docker-compose.test.yml up
remove_env:
	docker-compose -f ./deployment/docker-compose.test.yml down

test:
	go test -v -coverprofile cover.out ./...


coverage:
	go tool cover -html=cover.out -o cover.html


integrated_test:
	go test ./...

local_run:
	go run main.go