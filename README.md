## Run project locally
- Install go
- Set up mysql and create a new database, example: togo
- Config http port and database on .env file. example
```
HTTP_PORT=8080
APP_TIMEZONE="UTC"
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=togo
DB_USERNAME=root
DB_PASSWORD=123456
```
- Install goose package https://github.com/pressly/goose to migrate
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```
- Run migration using goose
```
cd ./migration && GOOSE_DRIVER=mysql GOOSE_DBSTRING="root:123456@/togo" goose status
```

- Run the app
```
go run app/main.go
```

### Run unit test
```
go test ./...
```

### Run a curl example
- Create a task and assign to user email
```
curl --location --request POST 'http://localhost:9099/tasks' \
--header 'Content-Type: application/json' \
--data-raw '{
    "content": "This is a task 6789",
    "user_email": "thang2223@gmail.com",
    "task_limit": 1
}'
```

- Create a task and assign to user id
```
curl --location --request POST 'http://localhost:9099/tasks' \
--header 'Content-Type: application/json' \
--data-raw '{
    "content": "This is a task example",
    "user_id": 1
}'
```
