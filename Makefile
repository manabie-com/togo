INTEGRATION_TEST_PATH?=github.com/manabie-com/togo/tests/...

# set of env variables that you need for testing
ENV_LOCAL_TEST=\
  JWK_KEY=wqGyEBBfPK9w3Lxw \
  PG_CONN_URI=postgresql://admin:123456@localhost:5432/togodb?sslmode=disable 

# this command will start a docker components that we set in docker-compose.yml
compose-all:
	docker-compose up -d --remove-orphans postgres;

# shutting down docker components
compose-down:
	docker-compose down;

# this command will trigger integration test
test-integration:
	go test -tags=integration $(INTEGRATION_TEST_PATH) -count=1
