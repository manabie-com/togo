## Cadidate: Danh Nguyen Hoang - nguyendanhtravel@gmail.com
### I. My improvement
#### 1. Split `services` layer to `use case` and `transport` layer
#### 2. Change to use `chi-go router`. Why chi-go? Because it lightweight, idiomatic, and composable router for building Go HTTP services. Especially, chi's router is based on `Radix trie`, so it'll handle the request as fast as possible if we have a lot of handlers in the future.
#### 3. We shouldn't use `GET` method for login because it have some problem it will save the user's information url, another user can find id and password in the browser's history,..So I have changed to `POST` method for /login
#### 4.1 Add config (with file - just easy for testing) - after load config file, it will overwrite the variable environment, so this project still abide by 12factor `https://12factor.net/` .P/s: Again the config file just save the infomation for quick run, when in production, use variable environment.
#### 4.2 Variable Environment in this project:
Key | Example value | Description
--- | --- | ---
`APP_ENV` | `dev` | represent the current app environment 
`SERVICE` | `todo` | name of the current service 
`JWT_KEY` | `wqGyEBBfPK9w3Lxw` | JWTKey 
`LOG_LEVEL` | `info` | represent level out the log 
`ADDRESS` | `:5050` | the address or port when the app will deploy
`LDB_PATH` | `./data.db` | the path of sqlite 
`PDB_USERNAME` | `postgres` | the username in postgres
`PDB_PASSWORD` | `secret` | the password with username in postgres
`PDB_HOST` | `127.0.0.1` | the host of postgres
`PDB_PORT` | `5432` | the port of postgres 
`PDB_DBNAME` | `todo` | the dbname of postgres 


#### 5. Added logging for this project with json format, for easily tracing log (error), using zap here (uber).
#### 6. Remove unnecessary pointer variable (replaced by variable) to avoid memory allocate in the runtime to much.
#### 7. Complete main requirement, limit N task per day each user.
#### 8. Add postgres DB to project, add prepared statement for sqlite, use pgx as driver to work with postgres database (included prepared statement). In pq's github, they still encourage use pgx.
#### 9. Store password in database as hashed password (`bcrypt` algorithm), to protect the user:the developer shouldn't know the user's plain password, protect against rainbow table attacks and so on and so forth.
#### 10. New schema for postgres, postgre have uuid type (16 bytes, because it save as binary) to save memory instead use text or varchar (36 bytes) save the uuid. Attention: litesql doesn't have uuid type, so still use text.
#### 11. Added dockerfile, docker-compose with postgres for the app.
```sql
-- users definition

CREATE TABLE users (
	id varchar(50) NOT NULL,
	password varchar(60) NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (id, password, max_todo) VALUES('firstUser', '$2a$10$hBqrcwfOt/HBKLXKxa48tu1SMDn62pSU8iZYWIXTxTCXQ8PoXvvi2', 5);

-- tasks definition

CREATE TABLE tasks (
	id uuid NOT NULL,
	content TEXT NOT NULL,
	user_id varchar(50) NOT NULL,
    created_date TEXT NOT NULL,
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);
```
-----
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

INSERT INTO users (id, password, max_todo) VALUES('firstUser', '$2a$10$hBqrcwfOt/HBKLXKxa48tu1SMDn62pSU8iZYWIXTxTCXQ8PoXvvi2', 5);

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
