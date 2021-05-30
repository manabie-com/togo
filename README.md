### Overview

This is a simple backend for a good old todo service, right now this service can handle login/list/create simple tasks.  

#### To make it run

- `go run main.go`
- Import Postman collection from `docs` to check example

#### To check unit test coverage 

`go test ./... --coverprofile cover.out && go tool cover -html=cover.out -o cover.html && open cover.html`
#### What have been done

- Implement daily limit for creating tasks ( base requirement)
- Write some integration tests for the project
- Write complete unit test for addTask function
- Refactor
    - extract `responseOK()` and `responseError()` to reduce duplication
    - extract interface for storages.LiteDB to loosen coupling and easy to write UT
    - some other small refactors

#### What missing

- Change from using `SQLite` to `Postgres` with `docker-compose`
- Unit test for db layer

#### What to improve if having enough time

- Add `controller` layer, responsible for routing and parsing body, `service` layer should only care about business logics
- Make authenticate into a separate middleware ( or filter), applied to all endpoints except `/login`
- Separate `ServeHTTP()` from `TodoService` to a standalone struct (maybe named `MainServer`), there are too many responsibilities for `TodoService`
- Get JWT key from a separated config file
- Separate `users` table to `users` and `accounts` table. `account` is for authentication and authorization (if any), and `users` is for basic information of
  user. There must be something wrong when we put `password` next to `max_todo` lol
- Write more tests to make it 100% test coverage
- Of course, make `created_date` type to `datetime`, if we plan to use any DB other than Sqlite 

### DB Schema
```sql
-- users definition

CREATE TABLE users (
	id TEXT NOT NULL,
	password TEXT NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (id, password, max_todo) VALUES('firstUser', 'example', 5);

-- tasks definition

CREATE TABLE tasks (
	id TEXT NOT NULL,
	content TEXT NOT NULL,
	user_id TEXT NOT NULL,
    created_date TEXT NOT NULL,
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);
```

### Sequence diagram

![auth and create tasks request](https://github.com/manabie-com/togo/blob/master/docs/sequence.svg)
