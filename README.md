### Requirements

- Implement one single API which accepts a todo task and records it
  - There is a maximum **limit of N tasks per user** that can be added **per day**.
  - Different users can have **different** maximum daily limit.
- Write integration (functional) tests
- Write unit tests
- Choose a suitable architecture to make your code simple, organizable, and maintainable
- Write a concise README
  - How to run your code locally?
  - A sample “curl” command to call your API
  - How to run your unit tests locally?
  - What do you love about your solution?
  - What else do you want us to know about however you do not have enough time to complete?
## Prerequisites:

[Go](https://go.dev/dl/)

[PostgreSQL](https://www.postgresql.org/download/)

## How to run source code locally?
**Installation**
- Clone the project
```bash
git clone https://github.com/lntvan166/togo.git
cd togo
```
- Install dependencies
```bash
go mod tidy
```
- Create `.env` by checkout `.env.example` to see all required environment variables.

**Setup Postgres database**

  *The first way: dockerize postgresql*
- Config environment variables in `Dockerfile` like `.env`
- With PORT is DATABASE_PORT in `.env`
```bash
sudo docker build -t togo .
sudo docker run -p {PORT}:5432 togo
```


  *The second way: migrate data by script*
- Open psql in your Terminal by:
```bash
  sudo -u postgres psql
```
- Create togo database:
``` bash
CREATE DATABASE togo; \c togo
```
- Copy and paste script in migrations/migrate.sql to generate necessary data, then exit psql
```bash
\q
```

**Run app:**
``` bash
make run
```

## A sample “curl” command to call APIs

- I created the mock users:
  - free plan: the limit of tasks is **10**
    - username: **free**
    - password: **password**
  - vip tier: the limit of tasks is **20**
    - username: **vip**
    - password: **password**

- You can login with the above accounts or register (free plan by default) by the following
```bash
curl --location --request POST 'localhost:8080/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "your_name",
    "password": "password"
}'
```

- Or login:

```bash
curl --location --request POST 'localhost:8080/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "free",
    "password": "password"
}'
```
- If login successfully, you can get a token:
``` json
{
  "message":"login successfully",
  "token":"this-token-is-very-very-very-long"
}
```

- Assign a task. I created the mock tasks for user **free** have (**10**) tasks.
  The user **vip** has (**0**) tasks.
  - If we assign one more task to user **free** the system will return error limit message.
  - Assign task: you need to replace **ADD-TOKEN-HERE** which is the token you got in the login API.
```bash
  curl --location --request POST 'localhost:8080/task' \
--header 'Authorization: Bearer ADD-TOKEN-HERE' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"todo",
    "description": "test"
}'
```
  - Retry with vip user:
```bash
curl --location --request POST 'localhost:8080/task' \
--header 'Authorization: Bearer ADD-TOKEN-HERE' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"todo",
    "description": "test"
}'
```

  - Know more about my api: [Togo](https://documenter.getpostman.com/view/21343860/UzJERdrX) (Please add this header to curl:
  ```--header 'Authorization: Bearer ADD-TOKEN-HERE'```)
  - Features: 
    - Authentication: Login, register
    - User: Get all users, get one user, get plan, upgrade plan
    - Task: Get all own tasks, get one own task, create task, complete task, delete task


## How to run unit tests locally?

- unit tests:

```bash
  make unit_test
```
```text
PASS
coverage: 72.7% of statements
ok      lntvan166/togo/internal/usecase 0.003s  coverage: 72.7% of statements
?       lntvan166/togo/pkg      [no test files]
?       lntvan166/togo/pkg/mock [no test files]
```

- integration tests:

```bash
  make integration_test
```
```text
example result:
PASS
ok      lntvan166/togo/internal/integration     (cached)
```

- all tests:

```bash
  make test
```

## What do you love about your solution?

- Learning about clean architecture, I found myself able to easily manage the project
- Using mock makes it easy to write test cases
- Easily add new features


## What else do you want us to know about however you do not have enough time to complete?

- Write more error unit test
- Dockerize my application
- Deploy

## Folder structure

```text
.
├── cmd
│   └── server
│       └── main.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── delivery
│   │   ├── auth.go
│   │   ├── auth_test.go
│   │   ├── handler.go
│   │   ├── plan.go
│   │   ├── plan_test.go
│   │   ├── task.go
│   │   ├── task_test.go
│   │   ├── user.go
│   │   └── user_test.go
│   ├── domain
│   │   └── domain.go
│   ├── entities
│   │   ├── task.go
│   │   └── user.go
│   ├── integration
│   │   └── integration_test.go
│   ├── middleware
│   │   ├── middleware.go
│   │   └── middleware_test.go
│   ├── repository
│   │   ├── db.go
│   │   ├── task.go
│   │   ├── task_test.go
│   │   ├── user.go
│   │   └── user_test.go
│   ├── routes
│   │   ├── auth.go
│   │   ├── plan.go
│   │   ├── routes.go
│   │   ├── task.go
│   │   └── user.go
│   └── usecase
│       ├── task.go
│       ├── task_test.go
│       ├── user.go
│       └── user_test.go
├── LICENSE
├── Makefile
├── migrations
│   ├── create_tasks_table.sql
│   ├── create_users_table.sql
│   └── migrate.sql
├── pkg
│   ├── crypto.go
│   ├── json.go
│   ├── jwt.go
│   ├── mock
│   │   └── mock.go
│   └── utils.go
└── README.md
```