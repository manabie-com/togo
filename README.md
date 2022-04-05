# How to run
make sure you have 'make' to run scripts

```sh
cp .env.example .env
```

set your ip to DB_HOST variable

```sh
make build-db build-app
```

```sql
CREATE TABLE users (
	id serial PRIMARY KEY,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	password VARCHAR ( 255 ) NOT NULL,
    limit_todo SMALLINT NOT NULL DEFAULT 5,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP
);

INSERT INTO users (email, password, created_at) VALUES ('test@gmail.com', 'jZae727K08KaOmKSgOaGzww/XVqGr/PKEgIMkjrcbJI=', now());

CREATE TABLE todos (
	id serial PRIMARY KEY,
	user_id BIGINT,
	task VARCHAR ( 255 ) UNIQUE NOT NULL,
	due_date TIMESTAMP,
	status SMALLINT NOT NULL DEFAULT 0,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP,
	UNIQUE (user_id, task),
	CONSTRAINT fk_user
	FOREIGN KEY(user_id)
	REFERENCES users(id)
);
```
# CURLs

Login:
```sh
curl -H "Content-Type: application/json" \
-d '{"email":"test@gmail.com","password":"123456"}' \
http://localhost:5000/api/login
```

Create task:

```sh
curl -H "Content-Type: application/json" \
-H "Authorization: Bearer <token>" \
-d '{"task":"task1","due_date":"2022-04-05 15:00"}' \
http://localhost:5000/api/todo/create
```

# Test

Integration test:
```sh
make build-integration-test
```

Unit-test:
```sh
make unit-test
```

# Architect, layout

- Clean architecture
- layout: https://github.com/golang-standards/project-layout

Why i love this solution?
- Database Independent
- Highly Testable

There are a few more things I want to do: write more test cases, create docs api, implement Logger package
