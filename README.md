# Manabie

## Introduction

A code challenge from manabie.

## Usage

#### Prerequisite

golang and docker must be installed.

#### How to run the app
- Before run program, must run docker-compose file and run db migration files
  - Run docker-compose
  ```
  docker-compose up -d
  ```
  - Run database migration files (after docker-compose)
  ```
  docker exec -it manabie_postgres bash
  psql -U postgres
  \c manabie
  ```
  copy two files *.sql and run it into console.

- In order to run the app without building the binary file, please run following commands:

  - Run with default env
  ```
  go run main.go
- If you want to run the app by building binary file, please run following commands:

```
go build -o challenge
./challenge
```

#### How to test the app

- Access to golang directory and run command go test:

```
go test ./...
```

## Source code structure explanation

```
|-coltroller/
|-db/
|-form/
|-middleware/
|-model/
|-main.go 
```
## Future

- Create tool to auto-generate SQL simple functions (CRUD)
- Split middleware and add Claim for all requests.
- Use redis to handle case MaxTodo.
- Create tool for log request, sql.
- Add dependencies injection.
- Add apis to update MaxTodo.
