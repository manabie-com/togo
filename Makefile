INTEGRATION_TEST_PATH?=github.com/manabie-com/togo/tests/...

# set of env variables that you need for testing
ENV_LOCAL_TEST=\
  POSTGRES_PASSWORD=123456 \
  POSTGRES_DB=togodb \
  POSTGRES_HOST=postgres \
  POSTGRES_USER=admin

# this command will start a docker components that we set in docker-compose.yml
compose-all:
	docker-compose up -d --remove-orphans postgres;

# shutting down docker components
compose-down:
	docker-compose down;

# this command will trigger integration test
test-integration:
	$(ENV_LOCAL_TEST) \
	go test -tags=integration $(INTEGRATION_TEST_PATH) -count=1
