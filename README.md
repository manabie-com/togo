
How to run:
- `docker-compose up -d` to deploy local MySQL (port 3306) and Redis (port 6379).
- `go run ./app/main.go` to start running back-end application (port 9090).

Sample request:
```
curl --location --request POST 'localhost:9090/user/todo' \
--header 'Content-Type: application/json' \
--data-raw '{
    "content": "Hello world",
    "user_id": 1
}'
```
