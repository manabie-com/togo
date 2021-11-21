# express-mongo
This is a togo project using the following technologies:
- [Express](http://expressjs.com/) for RESTful API
- [MongoDB](https://www.mongodb.com/) for database

# Requirements

- Implement one single API which accepts a todo task and records it
- There is a maximum limit of N tasks per user that can be added per day.

# Non-functional requirements:
- Fork this repo and show us your development progress via a PR
- Use another type of db you are most comfortable with instead of sqlite (MySql, Postgree, Oracle).
- Split into class constraints (service layer, use case layer, storage layer)
- Write unit tests

## DB Schema
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
## Sample statement to execute
- User Login

``` POST {"username": "firstUser", "password": "example" } /api/v1/auth/login ```

- GET todos or by id

``` GET /api/v1/todo ```
``` GET /api/v1/todo/:id ```

- POST todo

``` POST {"content": "example"} /api/v1/todo ```

- PATCH todo or by id

``` PATCH /api/v1/todo ```
``` PATCH /api/v1/todo/:id ```

- DELETE todo or by id

``` DELETE /api/v1/todo ```
``` DELETE /api/v1/todo/:id ```

- User Logout
``` GET /api/v1/auth/logout ```

## Running
Install Dependencies
```npm install```

Start Server
```npm start```

Run Tests
```npm test```

#### Sequence diagram
![auth and create tasks request](https://github.com/cesc1802/go_training/blob/master/docs/sequence.svg)