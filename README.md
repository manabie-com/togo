# Overview
This http server implements one single API which accepts a todo task and records it.
 - There is a maximum **limit of N tasks per user** that can be added **per day**.
 - Different users can have **different** maximum daily limit.
 - The server is written in Go, using service pattern and repository pattern to create a clean architecture 
to make it simple for organizing and maintaining.
 - The project includes unit tests for repository layer, using SQLite for database.

# Usage
## How to run locally:
- Clone this repo
- From the repo directory, run:
```shell
go build -o ./build/togo ./cmd/server && ./build/togo
```

## How to run test with coverage:
- From the repo directory, run:
```shell
go test -v -cover ./...
```

## Example “curl” command:
```shell
curl -X POST 'http://localhost:8080/task?user_id=1&task=homework'
```
Available data for testing:
- user_id = 1, daily_limit = 1
- user_id = 2, daily_limit = 2
- user_id = 3, daily_limit = 3

## What I love about the solution:
The project follows "clean architecture" concept, make use of service pattern and repository patten,
which make it seamlessly easy to organize and maintain.

## Todo:
- User authentication
- Unit test coverage could be improved
- Integration testing with docker
