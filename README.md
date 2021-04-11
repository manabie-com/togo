### Overview

To make it run:

- `docker-compose up` in a terminal
- `go run main.go` in an other terminal
- Import Postman collection from `docs` to check example

> With PostgreSQL implemented, pgx is used so these may be called:
>
> `go get github.com/jackc/pgconn`
>
> `go get github.com/jackc/pgx/v4`

#### Unit tests

Test endpoints with a mock DB: `go test -v ./...`

#### Integration tests

Basically just unit tests with PostgreSQL up and running: `go test -v ./... --tags=integration`

> I did not handle setup and teardown in integration tests. Later attempt to should be cached, but in case it's not, you one should call `docker-compose down` and up it again every time a testing process is done.

### Checklist

- [x] Right now a user can add many task as they want, we want ability to limit N task per day
> Status code 429 Too Many Requests is now returned when daily limit is reached.
- [x] Write integration tests for this project
> I hope tests with PostgreSQL count.
- [x] Make this code DRY
> I tried to make it DRY to a degree in the tests. But I admit it could be DRY even more. I did not mess around much with existing code.
- [x] Write unit test for `services` layer
- [x] Change from using `SQLite` to `Postgres` with `docker-compose`
- [ ] Write unit test for `storages` layer
- [ ] Split `services` layer to `use case` and `transport` layer