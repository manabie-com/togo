### Notes
This is a simple backend for a todo service, right now this service can handle login/list/create simple tasks, to make it run:
- Clone source
- `cd togo`
- `sh run.sh`
- Import PostManTodo.postman_collection from docs to check example

### Functional requirement:
Right now a user can add as many tasks as they want, we want the ability to **limit N tasks per day**.  
For example, users are limited to create only 5 tasks only per day, if the daily limit is reached, return 4xx code to the client and ignore the create request.

### Non-functional requirements:
- [x] **A nice README on how to run, what is missing, what else you want to improve but don't have enough time**
- [x] **Consistency is a MUST**
- [x] Fork this repo and show us your development progress by a PR
- [x] Write integration tests for this project
- [x] Make this code DRY
- [x] Write unit test for the services layer
- [ ] Change from using SQLite to Postgres with docker-compose
- [x] This project includes many issues from code to DB structure, feel free to optimize them
- [ ] Write unit test for storages layer
- [ ] Split services layer to use case and transport layer


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
