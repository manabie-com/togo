#set up postgres server
postgres:
	docker run --name=psql-togo -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret postgres:11.4
	
createdb:
	docker exec -it psql-togo createdb --username=root --owner=root todo

dropdb:
	docker exec -it psql-togo dropdb todo

# .PHONY: postgres createdb dropdb