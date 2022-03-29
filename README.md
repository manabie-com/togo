# Togo

Time start: 2022-03-25 20:04

### Dependencies

- Go 1.16
- PostgreSQL 14.2
- Docker/Podman (for integration test)
- Linux (not tested on Windows or MacOS)

### How to start

1. Setup PostgreSQL instance, create a database named `togo_dev`

1. Create a `.env` file at project root
    ```sh
    SERVER_PORT=3000
    POSTGRES_URL=postgres://postgres@localhost:5432/togo_dev?sslmode=disable
    ```

1. Run migration with
    ```go
    go run ./cmd/migrate
    ```

1. Start the server with
    ```go
    go run ./cmd/server
    ```

1. Create new user
    ```sh
    curl -X POST localhost:3000/api/v1/users -i \
      -H 'Content-Type: application/json' \
      -d '{"taskDailyLimit": 1}'
    ```

    Sample response:
    ```json
    {
      "id": "d8fLZ71wS3K",
      "taskDailyLimit": 1
    }
    ```

1. Create new task
    ```sh
    curl -X POST localhost:3000/api/v1/tasks -i \
      -H 'Content-Type: application/json' \
      -d '{
        "timeZone": "Asia/Ho_Chi_Minh",
        "task": {
          "userId": "d8fLZ71wS3K",
          "message": "resolve todos"
        }
      }'
    ```

    Sample response:
    ```json
    {
      "id": "uEiIZEqsWOe",
      "userId": "d8fLZ71wS3K",
      "message": "resolve todos"
    }
    ```

### Testing

#### Unit testing

```go
go test ./...
```

#### Integration testing

```go
go test -tags=integration ./...
```

### Q&A

> What do you love about your solution?
- Isolate domain logic from application logic
- Isolate unit and integration test
- Use `_test` suffix to test only public interfaces
- Central HTTP error handling
- Use container for testing on real PostgreSQL instance

> What else do you want us to know about however you do not have enough time to complete?
- Dedicated logger, structured request logging
- Transaction for integration test to rollback after each test
- Authentication

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

### Notes

- We're using Golang at Manabie. **However**, we encourage you to use the programming language that you are most comfortable with because we want you to **shine** with all your skills and knowledge.

### How to submit your solution?

- Fork this repo and show us your development progress via a PR

### Interesting facts about Manabie

- Monthly there are about 2 million lines of code changes (inserted/updated/deleted) committed into our GitHub repositories. To avoid **regression bugs**, we write different kinds of **automated tests** (unit/integration (functionality)/end2end) as parts of the definition of done of our assigned tasks.
- We nurture the cultural values: **knowledge sharing** and **good communication**, therefore good written documents and readable, organizable, and maintainable code are in our blood when we build any features to grow our products.
- We have **collaborative** culture at Manabie. Feel free to ask trieu@manabie.com any questions. We are very happy to answer all of them.

Thank you for spending time to read and attempt our take-home assessment. We are looking forward to your submission.
