## Apply exercises for Manabie Backend Engineer

- Name: Le Su Truong Giang
- Applying position: Backend Engineer


## Quick start
- Step 1: Setup dependencies and tools
```
make setup
```

- Step 2: Start database with docker 
```
make docker/up
```

- Step 3: Migrate database to newly created docker container
```
make db/up
```

- Step 4: Run the server
```
make run
```
or traditionally
```
go run main.go
```

- Step 5: Open postman, import the previously provided postman collection and
test the API

- Step 6 (Optional): Run unit test
```
make test
```

- Step 7 (Optional): Run linting
```
make lint
```


## What I have done
- Refactor single services layer into handlers layer (business/application layer) and services layers
- Additionally, added DTO and DAO as transport layer
- Refactor server into clearer structured server object with middlewares, injection, handler types (custom auth handler and RESTful handler)
- Implement fully testable code that can be mocked with interface
- Implement mock object which can be generated with code
- Provide database-first transport layer, generate from postgres database schema with customizable templates and generated unit test code, and with super fast access
instead of popular ORM plugins
- Implemented integration test for login handler, with enough case so far
- Implemented unit test for some utilities, services and handlers
- Added linting also, implement everything following lint rules
- Easy to run bash scripts which help to setup everything

## What I have not yet done 
- I have not yet implemented unit tests for every packages, but the implementation of unit tests of 
other packages shall be quite similar to implemented one 
- I have not yet implemented integration test for all handlers, but the implementation of other integration tests shall be quite similar to implemented one
- Integrate easyjson or json efficient marshalling package that help to increase performance of json marshalling and unmarshalling speed
- Implement better solution for max_todo validation on create new task

## Old README
### Overview
This is a simple backend for a good old todo service, right now this service can handle login/list/create simple tasks.  
To make it run:
- `go run main.go`
- Import Postman collection from `docs` to check example

Candidates are invited to implement below requirements but the point is not to resolve everything in a perfect way but selective what you can do best in a limited time.  
Thus, there is no correct-or-perfect answer, your solutions are way for us to continue the discussion and collaboration.
 
### Requirements
Right now a user can add many task as they want, we want ability to limit N task per day.

Example: users are limited to create only 5 task only per day, if the daily limit is reached, return 4xx code to client and ignore the create request.
#### Backend requirements
- A nice README on how to run, what is missing, what else you want to improve but don't have enough time
- Fork this repo and show us your development progess by a PR.
- Write integration tests for this project
- Make this code DRY
- Write unit test for `services` layer
- Change from using `SQLite` to `Postgres` with `docker-compose`
- This project include many issues from code to DB strucutre, feel free to optimize them.
#### Frontend requirements
- A nice README on how to run, what is missing, what else you want to improve but don't have enough time
- https://github.com/manabie-com/mana-do
- Fork the above repo and show us your development progess by a PR.
#### Optional requirements
- Write unit test for `storages` layer
- Split `services` layer to `use case` and `transport` layer

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
