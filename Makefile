all: build test

build:
	docker-compose build

start-storage-services:
	docker-compose up -d redis postgres

stop-storage-services:
	docker-compose down redis postgres

test:
	make start-storage-services; \
	sleep 2; \
	DB_URI="postgresql://postgres:postgres@127.0.0.1:5433/togo_db" \
	REDIS_URI="127.0.0.1:6380" \
	REDIS_PASS="" \
	go test ./... -v ; \
	code=$$? ;\
	printf "\n\n\n\n=================\n" && \
	if [ "$$code" -eq "0" ];\
	then printf "TEST ALL: SUCCESS\n";\
	else printf "TEST ALL: FAIL\n"; \
	fi;\
	printf "=================\n" \
	make stop-storage-services;

run:
	docker-compose up
