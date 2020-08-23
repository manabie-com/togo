### Overview

A simple backend for good old todo service, right now this service can handle login/list/create simple tasks.

### Operations

#### UNIX environment

First please install these packages before proceed these steps.

- go
- make
- docker
- docker-compose
- github.com/rubenv/sql-migrate

1. init - preparing develop environment

- `make init`

2. run - start running app

- `make dev`
- Import Postman collection from `docs` to check example

3. test - check app unit tests and integration test

- `make test`

### Development Progress

- Updated README on how to run, what is missing, what else I want to improve but don't have enough time.
- Updated development progress
- Writed integration tests for this project
- Writed unit test for `services` layer
- Changed from using `SQLite` to `Postgres` with `docker-compose`
- Hashed login `password` in db
- Hided sensitive environment variable

### What is missing

- Make this code DRY
- Optimize app
- Write unit test for `storages` layer
- Split `services` layer to `use case` and `transport` layer
- Frontend requirements

### Sequence diagram

![auth and create tasks request](https://github.com/phuwn/togo/blob/master/docs/sequence.svg)
