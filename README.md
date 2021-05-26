### Overview
This is a simple backend for a good old todo service, right now this service can handle login/list/create simple tasks.  
To make it run:
- `docker-compose up db`
- `go run main.go`
- Import Postman collection from `docs` to check example

#### What I have done
- Add rule to check limit of task before system add a new one (based on user's `max_todo`) 
- Change from using `SQLite` to `Postgres` with `docker-compose`

#### What is missing
- Missing unit test and integration test.

#### Note
- You can run `docker-compose up adminer` to turn up adminer tool to monitor database at `localhost:8080`
  - Credential to log in adminer: `postgres/MyPassword1`
