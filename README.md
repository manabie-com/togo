# TOGO PROJECT

## OVERVIEW
Implement one single API which accepts a todo task and records it.<br>
There is a maximum limit of N tasks per user that can be added per day.<br>
Different users can have different maximum daily limit.

## INDEX

- [Required](#Required)
- [Local development environment](#local-development-environment)
- [Sample curl API](#sample-curl-api)
- [Unit test](#unit-test)
- [About my solution](#about-my-solution)
- [Development](#development)


## Required
- VSCode  
  https://code.visualstudio.com/download
- Docker  
  https://docs.docker.com/get-docker/

## Local development environment

Develop on `Docker Container` using `VSCODE Remote Container`.
https://code.visualstudio.com/docs/remote/containers

Use `Colima` to run Docker Containers

Install Colima:
```
brew install colima
```

Then run colima, and every docker command should work as before:
```
colima start
```

To Stop container, run below command:
```
colima stop
```

If this is the first time, please setting environment variable and package:
```
cp example.env .env
go mod download
```

Start server:
```
go run main.go
```

## Sample curl API
Sample curl API to assign `task ids = 2,3` for `user id = 1`:
```
curl --location --request POST 'http://localhost:8080/users/1/tasks' \
--header 'Content-Type: application/json' \
--data-raw '{
    "task_ids": [2, 3]
}'
```

## Unit test
If this is the first time, please setting environment variable:
```
cp example_test.env test.env
```
Run test:
```
go test -coverprofile=coverage.out ./... 
```
Coverage:
```
go tool cover -html=coverage.out
```

## About my solution
My solution is nothing special, it's simple, easy to understand, easy to expand and maintain.

## Development
If have more time, I will create common function for reusable functions, create middleware, develop necessary features (CRUD users, tasks) and expand features, as said at PR description:
- A task has many subtasks.
- A task will be complete in more than 1 day.
- Assign task on one day but do it on another day.
- Who has the role to assign/do tasks?
- Task will have more information to save such as status, comments...