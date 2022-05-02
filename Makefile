services-up:
	$(info Make: Starting all services.)
	./scripts/start-services.sh

services-down:
	$(info Make: Shutdown all services.)
	./scripts/down-services.sh

mongo-up:
	$(info Make: Start MongoDB service.)
	./scripts/mongo/mongo-start.sh

mongo-down:
	$(info Make: Shutdown MongoDB service.)
	./scripts/mongo/mongo-down.sh

kafka-up:
	$(info Make: Start Kafka service.)
	./scripts/kafka/kafka-start.sh

kafka-down:
	$(info Make: Shutdown Kafka service.)
	./scripts/kafka/kafka-down.sh

redis-up:
	$(info Make: Start Redis service.)
	./scripts/redis/redis-start.sh

redis-down:
	$(info Make: Shutdown Redis service.)
	./scripts/kafka/redis-down.sh