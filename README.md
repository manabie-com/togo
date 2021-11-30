### Overview

This is a simple backend for good old todo service, right now this service can handle login/list/create simple tasks.

Candidates are invited to implement below requirements but the point is not to resolve everything in a perfect way but selective what you can do best in a limited time.  
Thus, there is no correct-or-perfect answer, your solutions are way for us to continue the discussion and collaborating.

### Requirements

Right now a user can add many task as they want, we want ability to limit N task per day.

Example: users are limited to create only 5 task only per day, if imit reached, return 4xx code to client and ignore the create request.

#### What I have done

- Limit N task per day.
- Change from using `SQLite` to `Postgres` with `docker-compose`
- Split `services` layer to `use case` and `transport` layer
- Write unit test for `storages` layer
- Write unit test for `service` layer
- Write unit test for `transport` layer

#### Run the testing

```bash
$ `go test -v -cover -covermode=atomic ./...`
```

#### Run the application

```bash
#Download dependencies
$ go get -v
or
download individual dependency having in go.mod file using go get [dependency path]

#Build postgres image
$ docker-compose up --build -d

#Run application
$ go run main.go
```

#### Test the application

- Import Postman collection from `docs` to check example

#### Features need add or improve

- Add middleware to limit adding task instead query database
- Add authentication/logging middleware
- Add model layer to scale and maintain easily
- Improve more variety of test case
