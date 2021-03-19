setup:
	docker-compose -f docker/docker-compose.yml --env-file config/envs/.env up

check_env:
	docker-compose -f docker/docker-compose.yml --env-file config/envs/.env config