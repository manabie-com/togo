### Functional requirement:
Right now a user can add as many tasks as they want, we want the ability to **limit N tasks per day**.  
For example, users are limited to create only 5 tasks only per day, if the daily limit is reached, return 4xx code to the client and ignore the create request.

### Non-functional requirements:
- [X] **A nice README on how to run, what is missing, what else you want to improve but don't have enough time**
- [X] **Consistency is a MUST**
- [X] Fork this repo and show us your development progress by a PR
- [ ] Write integration tests for this project
- [X] Make this code DRY
- [ ] Write unit test for the services layer
- [ ] Change from using SQLite to Postgres with docker-compose
- [X] This project includes many issues from code to DB structure, feel free to optimize them
- [ ] Write unit test for storages layer
- [X] Split services layer to use case and transport layer => ***use Hexagonal structure***

### Technical requirements:
- [X] Warm up: init database if it does not exists
- [X] Use [Uber Dig](https://github.com/uber-go/dig) as a dependency injection
- [X] Update method (change GET to POST) and use body instead of param as data of request of Login API
- [X] Use [Go Gin](https://github.com/gin-gonic/gin) to setup RestAPI server
- [X] Use username and password to login instead of userid and password
- [X] Run sql queries in transaction (create database package)

### Missing
- Use [mockery](https://github.com/vektra/mockery) to test service layer
- Use [sqlmock](https://github.com/DATA-DOG/go-sqlmock) to test storages layer
- Use [dockertest](https://github.com/ory/dockertest) for integration tests

### Inprove but don't have enough time
- Hash password with salt (use bcrypt)
- Get config from file or environment variables (ex: server address, database address, jwt key, ...)
- Build a log package

### How to run
To start the server, run 2 commands:
```golang
go mod tidy
go run ./cmd/main.go
```

To login, use default account
```
Username: tridm
Password: 123456789
```


#### DB Schema
```sql
-- users definition

CREATE TABLE IF NOT EXISTS user (
	id TEXT NOT NULL,
	username TEXT NOT NULL,
	password TEXT NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT user_PK PRIMARY KEY (id),
	CONSTRAINT user_UK UNIQUE (username)
);

INSERT INTO user
(id, username, password)
VALUES
("a5858792-ff00-475d-b9be-a6864f15eb28", "tridm", "123456789");

-- tasks definition

CREATE TABLE IF NOT EXISTS task (
	id TEXT NOT NULL,
	content TEXT NOT NULL,
	user_id TEXT NOT NULL,
	created_date TEXT NOT NULL,
	CONSTRAINT task_PK PRIMARY KEY (id),
	CONSTRAINT task_FK FOREIGN KEY (user_id) REFERENCES user(id)
);
```

#### Sequence diagram
![auth and create tasks request](https://github.com/manabie-com/togo/blob/master/docs/sequence.svg)
