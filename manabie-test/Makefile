.PHONY: up down nginx php phplog nginxlog db coverage vendor

MAKEPATH := $(abspath $(lastword $(MAKEFILE_LIST)))
PWD := $(dir $(MAKEPATH))
CONTAINERS := $(shell docker ps -a -q -f "name=todo*")

up:
	docker-compose up -d --build

down:
	docker-compose down

nginx:
	docker exec -it todo-nginx-container bash

php: 
	docker exec -it todo-php-container bash

phplog: 
	docker logs todo-php-container

nginxlog:
	docker logs todo-nginx-container

db:
	docker-compose exec mysql mysql -e 'DROP DATABASE IF EXISTS todo_test ; CREATE DATABASE todo_test;'
	docker-compose exec mysql sh -c "mysql todo_test < docker-entrypoint-initdb.d/database.sql"

coverage:
	docker-compose exec php-fpm sh -c "./vendor/bin/phpunit --coverage-text --coverage-html coverage"

vendor:
	docker-compose exec php-fpm sh -c "composer install"
