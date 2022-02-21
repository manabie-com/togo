# Manabie Assignment

### Requirements

- Implement one single API which accepts a todo task and records it
    - There is a maximum **limit of N tasks per user** that can be added **per day**.
    - Different users can have **different** maximum daily limit.

### Solutions
- The best way to get userID from the token. But, To make it simple, I use get userID from the header of the request with key: `auth-user-id`
- If the is no auth key or invalid key => cannot create the task
- If auth key is valid but there is no settings on `limitations` table => cannot create the task
- User can create a task when total daily tasks is less than setting on the DB

# Architecture: Clean Architecture + SOLID principle + Dependency Injection
- All values should be injected from main
- We have 3 main layers: handler -> service -> repo

### RUN unittests
- I use docker test to create database automatically. So, You have to run Docker successfully
- Run: `./gotest.sh`

# Run on the local
In Mac: set env by run command
- create a postgres db
`export BLOG_TEST_DB_URL="username:password@/dbName?parseTime=true"`

### Run with curl command

- With Invalid user
```bash
curl --location --request POST 'localhost:8081/todos' \
--header 'auth-user-id: abc-xyz' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Coding"
}'
```

- With valid user but without limit config
```bash
curl --location --request POST 'localhost:8081/todos' \
--header 'auth-user-id: 0a1a88c2-18f9-42dd-b4b0-99ee4dc88851' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Coding"
}'
```

- With valid users
```bash
curl --location --request POST 'localhost:8081/todos' \
--header 'auth-user-id: 0a1a88c2-18f9-42dd-b4b0-99ee4dc77751' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Coding"
}'
```

# Run on the local
- Update TODO_DB_URL with your Postgres DNS
- Run command: `local.sh`

