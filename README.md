# TOGO

Backend Engineer Coding Challenge

- Language: GoLang
- Database: PostgreSQL

# A. How to run this code locally?

## Install Go

Follow this link to install golang https://golang.org/doc/install

## Setup Golang $GOPATH directory

Follow this link to setup $GOPATH env https://golang.org/doc/gopath_code.html

## Clone this repository in $GOPATH/src directory

1. Navigate to gopath directory by executing command

```
$ cd $GOPATH/src
```

2. Clone the repository by executing command

```bash
$ git clone [this repository]
```

## Create `.env` file from the root folder with variables:

```
DATABASE_HOST=<database_host>
DATABASE_PORT=5432
DATABASE_USERNAME=<database_user>
DATABASE_PASSWORD=<database_password>
SSL_MODE=<database_ssl_mode>
```

## Install the database migration tool, [goose](https://github.com/pressly/goose) locally

```bash
$ go install github.com/pressly/goose/v3/cmd/goose@latest
```

For macOS users goose is available as a Homebrew Formulae:

```bash
$ brew install goose
```

## Run the database migration scripts

1. Create a database in your server with the name `psql_togo_db`

2. Navigate to the `migrations` directory

```bash
$ cd $GOPATH/src/togo/migrations
```

3. Check the migration script status

```bash
$ goose postgres "user=[your user] password=[your password] dbname=psql_togo_db host=[your host] sslmode=[your sslmode]" status
```

4. Run the migration sript to apply the changes in the database

```bash
$ goose postgres "user=[your user] password=[your password] dbname=psql_togo_db host=[your host] sslmode=[your sslmode]" up
```

## Run the API

### From the root directory `$GOPATH/src/togo/`, run without creating executable

```bash
$ cd $GOPATH/src/togo/
$ go run server.go
```

### Create executable file

1. Run the command `go build` to create an executable file
2. Execute the generate file by specifying the name e.g. `./togo`

# B. Sample “curl” command to call the API

1. Create User

```bash
curl -X POST http://localhost:8080/api/user -d '{"username":"readme","taskDailyLimit":2}'
```

2. Update User Task Daily Limit

```bash
curl -X PATCH http://localhost:8080/api/user -d '{"username":"readme","taskDailyLimit":1}'
```

3. Create Task

```bash
curl -X POST http://localhost:8080/api/task -d '{"username":"readme","title":"Sample title","description":"Sample description"}'
```

4. Delete User And Created Tasks

```bash
curl -X DELETE http://localhost:8080/api/user -d '{"username":"readme"}'
```

# C. How to run your unit tests locally?

1. Navigate to `rest` directory

```bash
$ cd $GOPATH/src/togo/rest
```

2. Run the command to execute the test

```bash
$ go test -v
```

# D. What do I love about my solution?

Deciding to make simple APIs that enabled us to create a user, update a user daily task limit differently, create task and, delete a user that completes the cycle of integration testing based on the requirements.

# E. What else do you want us to know about however you do not have enough time to complete?

- Completing the go testing coverage
- I can also create a GraphQL API endpoint with schema first approach written in Golang
