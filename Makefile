setup:
	docker-compose -f config/docker-compose.yaml up -d

check_env:
	docker-compose -f config/docker-compose.yaml config
