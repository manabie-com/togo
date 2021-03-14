### Quick start
1. create env
```bigquery
cp env.example .env
```

2.1 with docker-compose, must install Docker and docker-compose
```bigquery
docker-compose up --build
```

2.2 without docker
```bigquery
go run main.go
```
- you can get help by flag `-h`, `--help`
    ```bigquery
    go run main.go -h
    ```
  
3 run testcase
```bigquery
go test -race -count=1 -v -cover ./...
```
### what is missing? 
The key point missing:

- API /login, can't deactivate a token when user logout.
- API POST /tasks, missing handle mutex a function AddTask. 


### What did I do?
- _**use redis for mange token, remove token in redis for set deactivate token.**_
- _**handle mutex a AddTask function using Redlock**_
- refactor code base structure
- write unit test for services layer
- add run command
- use postgre to store all data.
- add index (user_id, created_date), improve performance of count user's task
- etc...

### Technical debt
- improve API's response
- add log to issue tracing
- improve code coverage
- detect sql injection
- graceful shutdown
- CI/CD

