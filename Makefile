# Executable file name
BINARY_NAME=togo

# we will put our unit testing in this path
UNIT_TEST_PATH?=./ut

# we will put our integration testing in this path
INTEGRATION_TEST_PATH?=./it

all: build
build: 
	go build -o $(BINARY_NAME) -v

clean: 
	go clean
	rm -f $(BINARY_NAME)

run:
	./$(BINARY_NAME)

# this command will start a docker components that we set in docker-compose.yml
docker.start.components:
	docker-compose up -d --remove-orphans postgres;

# shutting down docker components
docker.stop:
	docker-compose down;

# this command will trigger unit test
test.unit:
	go test -tags=unit $(UNIT_TEST_PATH) -count=1 -v

# this command will trigger integration test
# INTEGRATION_TEST_SUITE_PATH is used for run specific test in Golang, if it's not specified
# it will run all tests under ./it directory
test.integration:
	go test -tags=integration $(INTEGRATION_TEST_PATH) -count=1 -v -run=$(INTEGRATION_TEST_SUITE_PATH)
