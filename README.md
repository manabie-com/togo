### Notes

This is a simple backend for a todo service, right now this service can handle login/list/create simple tasks, to make
it run:

- `go run main.go`
- Import Postman collection from docs to check example

## togo

* Create task, get task by date
* Login user with username and password return access token(JWT)

## Libraries

- Golang
- This project use [Echo](https://echo.labstack.com/). Please check it for more information.
- [direnv](https://direnv.net/) load env in current directory
- [mockery](https://github.com/vektra/mockery) golang mock interface for testing
- [sqlmock](https://github.com/DATA-DOG/go-sqlmock) mock driver for golang to test database interactions
- [migrate](https://github.com/golang-migrate/migrate) database migrations. CLI and Golang library

---

## How to run

- Install `direnv` and run `direnv allow .` to load current directory enviroment
- Start docker-compose(contain postgresql)

```
make start-docker-compose
```

- Run database migartion

```
make migrate-up
```

- Run the server

```
make run
```

## Migrations

All migrate store in `migrations` folder

To create new migrate, run:

```
make create-migration
```

To migrate down, run:

```
make migrate-down
```

## Test

Run all test

```
make test
```

This project use `mockery` library to generate interface mocking for testing

Generate all interface, run:

```
make generate-mocks
```

### Functional requirement:

Right now a user can add as many tasks as they want, we want the ability to **limit N tasks per day**.  
For example, users are limited to create only 5 tasks only per day, if the daily limit is reached, return 4xx code to
the client and ignore the create request.

### Non-functional requirements:

- [x] **A nice README on how to run, what is missing, what else you want to improve but don't have enough time**
- [x] **Consistency is a MUST**
- [x] Fork this repo and show us your development progress by a PR
- [x] Write integration tests for this project
- [x] Make this code DRY
- [x] Write unit test for the services layer
- [x] Change from using SQLite to Postgres with docker-compose
- [ ] This project includes many issues from code to DB structure, feel free to optimize them
- [x] Write unit test for storages layer
- [x] Split services layer to use case and transport layer

#### DB Schema

```sql
-- user table
CREATE TABLE users
(
    username        VARCHAR PRIMARY KEY,
    hashed_password VARCHAR   NOT NULL,
    max_todo        INTEGER            DEFAULT 5 NOT NULL,
    created_at      TIMESTAMP NOT NULL DEFAULT (NOW()),
    updated_at      TIMESTAMP NOT NULL DEFAULT (NOW())
);

-- seed one user, hashed_password = "example"
INSERT INTO users (username, hashed_password, max_todo)
VALUES ('firstUser', '$2a$10$3jEtynoYdZJlw2fTUMjuCeGxHEjvc8a23gXMaidDW3yKPjMWFbb4W', 5);

-- tasks definition table

CREATE TABLE tasks
(
    id         SERIAL PRIMARY KEY,
    content    TEXT      NOT NULL,
    username   VARCHAR   NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (NOW()),
    updated_at TIMESTAMP NOT NULL DEFAULT (NOW()),
    CONSTRAINT tasks_FK FOREIGN KEY (username) REFERENCES users (username)
);
```

#### Sequence diagram

![auth and create tasks request](https://github.com/manabie-com/togo/blob/master/docs/sequence.svg)
