# Change Log - 2021-01-28
- All command is showed on `Makefile`

### Added
- Unit test for `tasks.go` and `db.go`
- Migration database
- `db.go`: add two functions
  - `GetUser`: get a user by id
  - `CountTasks`: count tasks that match userID AND created date
- `User` struct: add `MaxTodo`
- main.env
- Config folder

### Changed
- Database: `sqlite` to `postgresql`
- `addTask` function: support max task per day

### Run
- Create file config environment
    - Follow command:
    - `cp main.env.example main.env`
    - `make postgresql`
    - `make createdb`
    - `make migrateup`

- Run server:
    - `make server`
- Using postman create API