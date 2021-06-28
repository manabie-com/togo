# Introduction

A simple backend service which helps user manage their tasks.

### Authentication

Using username/password of user to exchange JWT token. Using this JWT token to authenticate to others endpoints to manage `task`.
Jwt token will be expired in 15 minutes.

#### Features

- Manage users
- Create task
- List Tasks

## Architecture Layer

##### 1. Config

- Manage our application's configuration env variables. 
Our application/service will read env variable from `run time` environment.

##### 2. Routers (or handlers)

- Routing request. Request will be routed to `delivery/rest`

##### 3. Delivery (or controller)

- Receive request and extract the necessary data from request, passing it to the service layer
- Build response request and return to clients.

##### 4. Services

- Represents the Use Case layer
- Implementing business tasks.

##### 5. Repositories

- Work (read/write/update ...) with the data storages (PostgreSQL, Redis ...)
- Mapping data from the storage format to business objects.

#### Data/request Flow

[routers] <---> [Delivery] <---> [Service] <---> [Repositories]

## How to run

#### Setup PostgreSQL databases

- Using `docker-compose` or install on `run machine`

Database info:

`database name`: `manabie_togo`
`password`: `ad34a$dg`
`user`: `togo`

Example `docker-compose.yml` file

```
version: '3.1'
services:
  db:
    container_name: pg_container
    image: postgres
    restart: always
    environment:
      POSTGRES_USER: root
      POSTGRES_PASSWORD: root
      POSTGRES_DB: test_db
    ports:
      - "5432:5432"
```

Or created database and user by using below commands:

```
CREATE DATABASE manabie_togo;
CREATE USER togo WITH PASSWORD 'ad34a$dg';

ALTER ROLE togo LOGIN ;
ALTER ROLE togo SET client_encoding TO 'utf8';
ALTER ROLE togo SET default_transaction_isolation TO 'read committed';
ALTER ROLE togo SET timezone TO 'UTC';

GRANT ALL PRIVILEGES ON DATABASE manabie_togo TO togo;
```

And Create tables:

https://github.com/swdream/togo/blob/feature/THANHNT-Todo-Service-submission/internal/scripts/01_users_tasks_create_tables.sql
```

CREATE TABLE IF NOT EXISTS users (
	id VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (id, password, max_todo) VALUES('firstUser', 'example', 5);

CREATE TABLE IF NOT EXISTS tasks (
	id VARCHAR(255) NOT NULL,
	content VARCHAR(255) NOT NULL,
	user_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);
```


Note: this info is configured in DB env configuration by default
https://github.com/swdream/togo/blob/feature/THANHNT-Todo-Service-submission/internal/pkgs/clients/postgres.go#L11
``envconfig:"POSTGRES_DSN" required:"true" default:"host=localhost user=togo password=ad34a$dg dbname=manabie_togo port=5432 sslmode=disable timezone=UTC"``

#### Setup Redis

I used redis to set `rate limit` for the number of times user can add `task` per day.
Default Redis URL is set here
https://github.com/swdream/togo/blob/feature/THANHNT-Todo-Service-submission/internal/pkgs/clients/redis.go#L10
`envconfig:"REDIS_URL" required:"true" default:"redis://localhost:6379"`

#### Run service

- `go run main.go`
- Import Postman collection from docs to check requested. I updated something.


#### what is missing, what else you want to improve but don't have enough time? 

- Write Integration Test. In our company, we use `Postman` to do integration test
- Format Logger
- Format Response of Request
- Validate Request data/params
- Unittest: I just write some test case for `success` and `failed` scenarios. It is not too detail
 and I used `https://github.com/vektra/mockery` for  mocking.
- We can expose more APIs to manage user (create/update, list ...) and manage tasks.
