# Manabie-com/togo

## What I got
- this project using Golang latest (v1.16)
- use cobra to split into 2 app: web app and migrate cli
- update schema database
- add user_signup
- update user_login with POST
- CRUD tasks with some filter
- split service into transport, domain, handler, repo and storage
- use `sqlc` to generate storage code
- write unit test for storage layer
- change from using `SQLite` to `Postgres`
- add `Redis` to help caching and descrease number of Postgres query
- migration with `goose`
- build a optimize dockerfile ( the final image only have 28MB)
- setup docker-compose with postgres, redis, nginx and web app server
- setup minimize nginx for load balancing to 3 containers web app server

## Project Structure
```
.
├── Makefile
├── cmd
│   ├── app
│   │   ├── cmd
│   │   │   ├── handlers.go
│   │   │   └── root.go
│   │   └── main.go
│   ├── internal
│   │   ├── postgresql.go
│   │   └── redis.go
│   └── migrate
│       ├── cmd
│       │   └── root.go
│       └── main.go
├── common
│   └── errors.go
├── config
│   ├── config.go
│   └── serverConfig.go
├── docker-compose.yml
├── internal
│   ├── domain
│   ├── dto
│   ├── entity
│   ├── handler
│   ├── middleware
│   ├── postgresql
│   ├── redix
│   ├── repository
│   └── transport
├── migrations
├── nginx
└── utils
    ├── password.go
    └── validator
        └── validator.go

```

## DB Schema
```sql
CREATE TABLE users
(
    id         serial PRIMARY KEY,
    username   TEXT      NOT NULL,
    password   TEXT      NOT NULL,
    max_todo   INTEGER            DEFAULT 5 NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX users_username_index ON users (username);

CREATE TABLE tasks
(
    id           serial PRIMARY KEY,
    content      TEXT      NOT NULL,
    user_id      int4      NOT NULL,
    created_date DATE      NOT NULL DEFAULT CURRENT_DATE,
    is_done      BOOLEAN   NOT NULL DEFAULT FALSE,
    created_at   TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT tasks_fk FOREIGN KEY (user_id) REFERENCES users (id)
);
```

## Install
duplicate file `.env.example` and rename to `.env`, add your enviroment variables to this file, then source it

```bash
source .env
```

create a database with name match with `$DB_NAME`

migrate database

```bash
make migrate
```

start server

```bash
make run
```
run the test

```bash
make test
```

We have another way to start is using `docker`, make sure docker is running in your computer, then build and start docker-compose

```bash
make docker-build
make docker-up
```

## Postman

please add the postman collection into your postman app, then create a env with
```
url=localhost:5050
jwt=xxx (get jwt after user login/signup)
```

## Ideas for the next step
- Write unit test for the transport, domain, handler, repo layer
- write intergation tests for this project
- add ElastichSearch to fast query task wiht multi column
- Deploy to AWS, using ECR, ECS and jenkins/github action
- support pagination
- add detail code with error number and http status
