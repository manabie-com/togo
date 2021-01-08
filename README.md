### Todos
    - [x] Write Integration Test (usecase, storages layer).
    - [x] Structure code by simple MVC Architecture.
    - [x] Limit Create Task per day can avoid race condition.
    - [x] Split `services` layer to `use case` and `transport` layer.
    - [x] Change from using `SQLite` to both `SQLite` and `Postgres`.
    - [x] Hashing user's password.
    - [x] Error handling.
    - [ ] Write Integration Test (transport layer).
    - [ ] Using docker-compose.

### How to run 
    - Please set env to `config.json` and run `go run main.go`

### How to test 
	- Please set env to `test.config.json` for every layer before run test