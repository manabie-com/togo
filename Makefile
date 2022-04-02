run-dev:
	go run cmd/app/main.go

start:
	docker-compose up -d --build

stop:
	docker-compose down

unit-test:
	chmod +x ./scripts/unit_test.sh	
	./scripts/unit_test.sh

integration-test:
	chmod +x ./scripts/integration_test.sh
	./scripts/integration_test.sh