migrate:
	go run ./cmd/migrate/main.go	

# run:
# 	go run ./cmd/main.go

test:
	chmod +x ./test.sh && ./test.sh

docker-start:
	docker-compose down
	docker-compose up 