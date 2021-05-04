### Getting started 


### Overview
This is a simple backend for a good old todo service, right now this service can handle login/list/create simple tasks.  

If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer. The kit requires Go 1.4/X-version) or above.

Install Postgresql database. You need to get into /docker and run following commands:

```
docker-compose up --d
docker exec -ti docker_postgres_1 bin/bash
su postgres
psql
CREATE DATABASE manabie;
CREATE USER manabie WITH PASSWORD 'manabie';
ALTER ROLE manabie SET client_encoding TO 'utf8';
ALTER ROLE manabie SET default_transaction_isolation TO 'read committed';
ALTER ROLE manabie SET timezone TO 'UTC';
GRANT ALL PRIVILEGES ON DATABASE manabie TO manabie;
```

Next, create a PostgreSQL database named `manabie` and execute the SQL statements given in the file `sql/db.sql`.
The project uses the following default database connection information:
```
database_host : localhost
database_username : manabie
database_password : manabie
database_database_name : manabie
database_max_connection : 10
```

Now you can build and run the application by running the following command:
- `go run main.go`
- Import Postman collection from `docs` to check example

The application runs as an HTTP server at port 5050. It provides the following RESTful endpoints:

### TODO


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
