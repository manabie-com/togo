## How to run your code locally?

### Install docker
Follow Docker document and install Docker Desktop: https://docs.docker.com/engine/install/

### Install `sql-migrate`
```bash
go get -v github.com/rubenv/sql-migrate/...
```

### `cd` to the project directory

### Copy environment variables from .env.example to .env - do some edit if you want or just keep sample env
```bash
cp .env-example .env
```

### Start Docker
```bash
docker-compose up
```

### Run migration to create db tables
```bash
make migrate
```

## A sample “curl” command to call your API

- You have to call create new user API first to get user id, then put user id in create new task API
- Due date in create new task API muse be in the future, update it if necessary

### Create a new user
```bash
curl --location --request POST 'http://localhost:8080/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "dailyTaskLimit": 10
}'
```

### Create a new task
```bash
curl --location --request POST 'http://localhost:8080/tasks' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userId": "7f572763-2079-4002-929e-5ae6eeefee9a",
    "title": "Learn Golang Course",
    "description": "Learn Golang Course on Udemy for 1 hour",
    "dueDate": "2022-06-25T20:10:10Z"
}'
```

## How to run your unit tests locally?
- Keep docker running, we need to keep database running
- Run this command to run unit tests
```bash
go test -v -cover ./...
```

## What do you love about your solution?
It is simple and clear. Data is verified carefully when creating new task.
I decided to add a property `dueDate`, so user can create a task that end in the future.

## What else do you want us to know about however you do not have enough time to complete?
- With the scope in requirement, I think my solution is good enough.
- But I want to put some extra features if I have more time, like:
  - Implement authentication when user create task
  - Add attachment to task
  - Add priority to task
  - Repeating task
- Unit tests and integration tests are not fully tested all the cases