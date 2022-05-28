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

- Go [1.18](https://go.dev/dl/)
- [Docker](https://docs.docker.com/engine/install/)
- [Docker compose](https://docs.docker.com/compose/install/)

## How to run source code locally?

```bash
  make up
```

## A sample “curl” command to call APIs

- Login to get JWT token. I created the mock users:
  - free tier: the limit tasks is **5**
    - username: **free1**
    - password: **password**
  - standard tier: the limit tasks is **10**
    - username: **standard1**
    - password: **password**
  - Since the system does not implement policy feature yet, so the user can assign any task to another user
```bash
curl --location --request POST 'localhost:8080/api/v1/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "free1",
    "password": "password"
}'
```

- Assign a task. I created the mock tasks [id: 1-> 6] for user **free1** have (**5**) tasks, id [1-5].
  The user **standard1** have (**0**) task.
  - If we assign one more task to user **free1** the system will return error limit message.
  - Assign task: you need to replace **ADD-TOKEN-HERE** with the token got in the login API.
```bash
  curl --location --request PATCH 'localhost:8080/api/v1/tasks' \
--header 'Authorization: Bearer ADD-TOKEN-HERE' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id":6,
    "assignee": "free1"
}'
```

## How to run unit tests locally?

- unit tests:

```bash
  make unit_test
```
```text
--- PASS: TestHandler_GetByUserName (0.00s)       
    --- PASS: TestHandler_GetByUserName/OK (0.00s)
PASS
coverage: 42.9% of statements

=== RUN   TestQueries_CreateUserIntegration
    user.sql_test.go:38: skipping integration test
=== RUN   TestQueries_GetUserIntegration
    user.sql_test.go:45: skipping integration test
--- SKIP: TestQueries_GetUserIntegration (0.00s)

```

- integration tests:

```bash
  make integration_test
```
```text
example result:
PASS
ok      github.com/dinhquockhanh/togo/internal/pkg/sql/sqlc     0.172s

```

- all tests:

```bash
  make test
```
## Clean dev infrastructure:

```bash
  make down
```

## What do you love about your solution?

- Easy to write unittest, using mock
- Dependency injection make another members can develop without conflict, just need to inject the interface
- Easy to add more feature

## What else do you want us to know about however you do not have enough time to complete?

- Add policy feature
- Write more unit tests
- Write GitHub action
- Push Docker image to AWS ECR
- Deploy
- ...
