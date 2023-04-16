# Manabie-com/togo

## What I got
- this project using Golang v1.19
- split into migrate application and server application
- update schema database
- add register/login user
- authen with JWT
- add user_signup
- update number limit task per day of user
- CRUD tasks with some filter
- split service into transport, domain, handler, repo and storage
- use `sqlc` to generate storage code
- write unit test for handler layer
- use Postgres for primary database
- use Redis for caching and handle count today task of user
- migration with `goose`
- development code with docker, docker-compose and taskfile
- build a optimize dockerfile
- setup docker-compose with postgres, redis, server

## Project Structure
```
├── Taskfile.yaml
├── backend
│   ├── cmd
│   │   └── server #init server
│   ├── common
│   ├── modules
│   │   ├── task # module task split into transport, handler, repo, storage and model
│   │   └── user # module user split into transport, handler, storage and model
│   └── plugin
├── docker-compose.yml
├── libs # libs
└── migrate # migrate app

```

## DB Schema
```sql
CREATE TABLE users
(
    id         serial PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password   TEXT      NOT NULL,
    salt       TEXT      NOT NULL,
    limit_task   INTEGER            DEFAULT 5 NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX users_email_index ON users (email);

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

install taskfile at taskfile.dev

start database and redis
```
docker-compose up redis-db postgres-db
```

start migrate
```
task migrate -- up
```

start server
```
task server
```

## Curl help

register
```bash
curl --location 'localhost:4000/api/users/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "aaa@abc.com",
    "password": "123123123"
}'
```

login return jwt token
```bash
curl --location 'localhost:4000/api/users/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "aaa@abc.com",
    "password": "123123123"
}'
```

create task
```bash
curl --location 'localhost:4000/api/tasks?is_done=true' \
--header 'Authorization: Bearer ${token}' \
--header 'Content-Type: application/json' \
--data '{
    "content": "content"
}'
```

update task
```bash
curl --location --request PATCH 'localhost:4000/api/tasks/3' \
--header 'Authorization: Bearer ${token}' \
--header 'Content-Type: application/json' \
--data '{
    "is_done": true
}'
```

list task
```bash
curl --location 'localhost:4000/api/tasks?created_date=2023-04-14&is_done=true' \
--header 'Authorization: Bearer ${token}'
```

delete task
```bash
curl --location --request DELETE 'localhost:4000/api/tasks/3' \
--header 'Authorization: Bearer ${token}'
```
## Postman

please add the postman collection into your postman app, then create a env with
```
url=localhost:4000
token=xxx (get jwt after user login)
```

## Ideas for the next step
- Write unit test for the transport, handler, repo and storage layer
- write intergation tests for this project
- Store jwt in redis to check expire time and whitelist token
- caching user in redis to descrease number call to database
- Deploy to AWS, using ECR, ECS and jenkins/github action
