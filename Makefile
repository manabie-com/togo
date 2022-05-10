.PHONY: deploy
# Remember to adjust env file and main.go before make

build-test:
	cd dockerdb \
		&& docker build -t manabie-mssql-test:latest . \
		&& cd ..
build:
	docker build --platform linux/amd64 -t manabie-test:latest . \
		&& cd dockerdb \
		&& docker build -t manabie-mssql:latest . \
		&& cd ..
deploy:
	make build && \
	docker-compose -f docker-compose.yaml down && docker-compose -f docker-compose.yaml up
setup-integration-test:
	make build-test \
		&& docker-compose -f docker-compose-test.yaml down \
		&& docker-compose -f docker-compose-test.yaml up -d
integration-test:
	go test -tags=integration -v ./...
unit-test:
	go test -tags=unit -v ./...
test-coverage:
	go test -tags="unit integration" -v ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
test:
	make integration-test \
		&& make unit-test
test-all:
	docker-compose -f docker-compose-test.yaml down \
		&& make build-test \
		&& docker-compose -f docker-compose-test.yaml up -d \
  		&& make integration-test \
		&& make unit-test \
		&& docker-compose -f docker-compose-test.yaml down
