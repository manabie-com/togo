### Notes
This is a simple backend for a todo service, right now this service can handle login/list/create simple tasks, to make it run:
- The version of Golang is 1.16
- Setup Postgres database with the sql script as below
- Update configuration information on config.yml file such as Postgres database information
- `go run main.go`
- Import Postman collection from docs to check example
- Missing items that have not been done:
	+ Write test cases
	+ Create Docker-Compose script
#### DB Schema
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

#### Sequence diagram
![auth and create tasks request](https://github.com/manabie-com/togo/blob/master/docs/sequence.svg)
