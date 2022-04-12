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
	docker-compose -f docker-compose.yaml down && docker-compose -f docker-compose.yaml up -d
setup-integration-test:
	make build-test \
		&& docker-compose -f docker-compose-test.yaml down \
		&& docker-compose -f docker-compose-test.yaml up -d
integration-test:
	go test -tags=integration -timeout 30s -coverprofile=coverage.out github.com/manabie-com/togo/handlers
unit-test:
	go test -tags=unit -timeout 30s -coverprofile=coverage.out github.com/manabie-com/togo/models github.com/manabie-com/togo/initializer
test-coverage:
	go test -tags="unit integration" ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
test:
	make integration-test \
		&& make unit-test
test-all:
	docker-compose -f docker-compose-test.yaml down \
		&& make build-test \
		&& docker-compose -f docker-compose-test.yaml up -d \
  		&& make test \
		&& docker-compose -f docker-compose-test.yaml down
