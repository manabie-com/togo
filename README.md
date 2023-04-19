### Run Project.
Use docker run command:
- `docker-compose up`

### List Api:
#### Call curl below or import file postman in package document.
- auth/register: register account.
  ``
  curl --location --request POST 'localhost:9000/auth/register' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "userName": "admin",
  "passWord": "123456"
  }'
  ``
- auth/login: login
  ``
  curl --location --request POST 'localhost:9000/auth/login' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "userName": "admin",
  "passWord": "123456"
  }'
  ``
- task/create: create task.
  ``
  curl --location --request POST 'localhost:9000/task/create' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODcwMjAwMDEsImlzcyI6Im1hbmliaWUtdG9kbyIsIklkIjoyfQ.VYSoGawZE_6XrXRfurkswKYo1x1doAAmFDpuplrUFcU' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "content": "Task test danh 4"
  }'
  ``
- task/list: get task by created date.
  ``
  curl --location --request GET 'localhost:9000/task/list?createdDate=2023-04-18' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODcwMjAwMDEsImlzcyI6Im1hbmliaWUtdG9kbyIsIklkIjoyfQ.VYSoGawZE_6XrXRfurkswKYo1x1doAAmFDpuplrUFcU'
  ``

### Solution
Limit the user to add work by day if the limit is reached then use redis cache to
reduce the load on the database, cache on redis with userId and expire at the end of the day.

### Technology used
- Golang version 1.19
- Postgres
- Redis
- Gin, Gorm
- Docker
