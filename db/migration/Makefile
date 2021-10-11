postgres:
	docker run --name postgres-todo -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -ti postgres-todo createdb --username=root --owner=root todo

dropdb:
	docker exec -ti postgres-todo dropdb todo

.PHONY: postgres createdb dropdb