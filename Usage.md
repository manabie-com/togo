### run app
```
	docker-compose up --build
```


### run test
```
	docker-compose exec togoapp go test -v ./internal/integration_test/
```


### access database by psql. run database service
if you don't install postgres or psql on your local machine
```
	docker-compose run database psql -h database -U togoapp -d togodb
```


if you install postgres 
```
	psql -h localhost -U togoapp -d togodb
```

then password promt appear: 
	password: togoapp


### backup database by using pg_dump
```
	docker-compose exec database pg_dump -h database -U togoapp togodb > togodb.sql
```
then enter password: togoapp

