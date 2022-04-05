DB_IMAGE=postgres:13.2
DB_PORT=5432
DB_TEST_PORT=5433
APP_PORT=5000
DB_ENV=-e POSTGRES_USER=togo -e POSTGRES_DBNAME=togo -e POSTGRES_PASSWORD=password -e POSTGRES_PORT=${DB_PORT}
DB_TEST_ENV=-e POSTGRES_USER=togo_test -e POSTGRES_DBNAME=togo_test -e POSTGRES_PASSWORD=password -e POSTGRES_PORT=${DB_TEST_PORT}
#app
togo-image:
	@echo "Build togo image"
	docker build -t togo .
togo-cont:
	@echo ${CUR_DIR}
	@echo "Run togo container"
	docker run -d --env-file=.env -p ${APP_PORT}:${APP_PORT} --name togo-cont togo
remove-togo-cont:
	@echo "Remove togo container"
	-docker stop togo-cont
	-docker rm togo-cont

#database
build-db:
	@echo "Run database container"
	docker run -d -p ${DB_PORT}:${DB_PORT} ${DB_ENV} --name db-cont ${DB_IMAGE}

remove-db-cont:
	@echo ":::Remove database container"
	-docker stop db-cont
	-docker rm db-cont

#database test
build-test-db:
	@echo "Run database test container"
	docker run -d -p ${DB_TEST_PORT}:${DB_PORT} ${DB_TEST_ENV} --name db-test-cont ${DB_IMAGE}

remove-db-test-cont:
	@echo ":::Remove database test container"
	-docker stop db-test-cont
	-docker rm db-test-cont

build-app: remove-togo-cont togo-image togo-cont

integration-test:
	-go test ./integration_tests/

build-integration-test: build-test-db integration-test remove-db-test-cont

unit-test:
	-go test ./internal/pkg/...
