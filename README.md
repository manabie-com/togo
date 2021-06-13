### Overview

This is a simple backend for a good old todo service, right now this service can handle login/list/create simple
tasks.  
To make it run:

- `go run main.go`
- Import Postman collection from `docs` to check example

Candidates are invited to implement below requirements but the point is not to resolve everything in a perfect way but
selective what you can do best in a limited time.  
Thus, there is no correct-or-perfect answer, your solutions are way for us to continue the discussion and collaboration.

### Requirements

Right now a user can add many task as they want, we want ability to limit N task per day.

Example: users are limited to create only 5 task only per day, if the daily limit is reached, return 4xx code to client
and ignore the create request.

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

CREATE TABLE users
(
    id       TEXT              NOT NULL,
    password TEXT              NOT NULL,
    max_todo INTEGER DEFAULT 5 NOT NULL,
    CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (id, password, max_todo)
VALUES ('firstUser', 'example', 5);

-- tasks definition

CREATE TABLE tasks
(
    id           TEXT NOT NULL,
    content      TEXT NOT NULL,
    user_id      TEXT NOT NULL,
    created_date TEXT NOT NULL,
    CONSTRAINT tasks_PK PRIMARY KEY (id),
    CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users (id)
);
```

### Sequence diagram

![auth and create tasks request](https://github.com/manabie-com/togo/blob/master/docs/sequence.svg)

## My Solution

### Overview

To make it run:

* Run docker compose file `deployments/docker-compose.yml`
* ```go run main.go```
* `scripts/req.http` contain example http request, i don't like postman :(
  , [click here for more details!](https://www.jetbrains.com/help/idea/exploring-http-syntax.html)
* Open swagger docs with [link](http://localhost:5050/swagger/index.html)
* Open Jaeger UI with [link](http://localhost:16686/)
* `enhancement.md` contain my draft enhancement idea 

In case you don't like `scripts/req.http`, please try with `curl` and replace `Authorization credentials` with `access_token` from login result.

Login:
```shell
curl -X POST --location "http://localhost:5050/login" \
    -H "Accept: application/json" \
    -H "Content-Type: application/json" \
    -d "{
          \"username\": \"firstUser\",
          \"password\": \"example\"
        }"

```
Retrieve all tasks:

```shell
curl -v -X GET --location "http://localhost:5050/tasks?created_date=2021-06-13" \
    -H "Accept: application/json" \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjM1OTQzODIsIm1heF90b2RvIjo1LCJ1c2VyX2lkIjoiNDgyODM5ZGYtZWQ3MC00NjUzLWJhMTEtN2I1ZTJkOTRkZTJhIn0.i3_Oh9WXbrKxppPaBGEw3iWNmttxVm7fxJiwwkhGpog"
```

Create tasks: 
```shell
curl -v -X POST --location "http://localhost:5050/tasks" \
    -H "Accept: application/json" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjM1OTQzODIsIm1heF90b2RvIjo1LCJ1c2VyX2lkIjoiNDgyODM5ZGYtZWQ3MC00NjUzLWJhMTEtN2I1ZTJkOTRkZTJhIn0.i3_Oh9WXbrKxppPaBGEw3iWNmttxVm7fxJiwwkhGpog" \
    -d "{
          \"content\" : \"another content\"
        }"
```

Hope you enjoy!!!

### What's done

* Refactor project structure to multiple layer
* Limit user task per day with rate limiter
* Apply Dependency injection
* Write simple test case for `service` layer
* Refine header authorization, error header
* Encrypt password
* Change from using `SQLite` to `Postgres` with `docker-compose`. I also refactor table `user`, `task` and use ORM with `ent`
* Refactor API `/login`
* Use `chi` for http server and use authentication interceptor
* Add tracing with jaeger
* Add simple swagger (not completed yet)

### What's improvement i want to do?

* Write more unit test and integration test
* Auto retry for database connection
* Make idempotent rest API. I think we should support `Etag`
* Add more api docs
* Support health check, graceful shutdown
* We will encode/decode request/response base on http header (`Content-Type`, `Accept`)
* Make configurable configuration. We will move config parameter (jwt key...) to file config or env depend on which environment we will deploy this app.
* Make Database provider configurable, so we can replace `postgresql` to `mysql` or `sqlite` without difficulty or effort.