.PHONY: up

COMPOSE := docker-compose

up:
	$(COMPOSE) up -d

down: ## Kill all containers
	$(COMPOSE) kill
