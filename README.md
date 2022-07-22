# README

## How to run your code locally?

### Set up MySQL

First time initialize mysql container

```sh
docker run --name mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=gouser -e MYSQL_PASSWORD=gopassword -e MYSQL_DATABASE=godb -p 3306:3306 mysql:latest
```

Access to MySQL

```sh
mysql -h0.0.0.0 -P 3306 -ugouser -pgopassword
```

Run [migration](./internal/store/migrations/)

```sql
USE godb;

CREATE TABLE IF NOT EXISTS `todo_tasks` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `user_id` bigint(20) UNSIGNED NOT NULL,
    `task_name` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_todo_tasks_user_id` (`user_id`)
);
```

### Run server

```sh
go run main.go
```

## A sample `curl` command to call your API

API which accepts a todo task and records it

Method: `POST`

Path: `/user/:user_id/task`

Parameter:

- `user_id`: Uint64

Body:

```javascript
{
    name: string
}
```

Example:

```sh
curl -X POST localhost:8080/user/1/task -d '{"name": "do exercise"}'
```

API for getting the all task of a user

Method: `GET`

Path: `/user/:user_id/task`

Parameter:

- `user_id`: Uint64

Response:

```javascript
{
    message: string,
    data: [
        {
            id: numnber,
            user_id: number,
            task_name: string
        }
    ]
}
```

Example:

```sh
curl -X GET localhost:8080/user/1/task
```

## How to run your unit tests locally?

```sh
go test ./...
```

## What do you love about your solution?

I started the project from scratch and slowly picked the suitable tools to solve the problem. I structured the project to be a candidate for a boilerplate project. I correct things I felt wrong with all previous templates that I used. So I love this solution.

## What else do you want us to know about however you do not have enough time to complete?

The configuration is hardcoded. So when run the project as a container it cannot communicate with MySQL container. I underestimate the importance of the config module.

## How to use the project

### Update store - database handler module

- Test your query in the real database.
- Write the query (using query parameters) in [queries](./internal/store/queries/)
- Update model (using CREATE TABLE) in [schemas](./internal/store/schemas/)

Install [sqlc](https://docs.sqlc.dev/en/stable/) then run command

```sh
cd [togo path]/internal/store
sqlc generate
```

### Build and Run docker

```sh
docker build --tag togo .
```

```sh
docker run --publish 8080:8080 todo
```

### Mock

Install [golang mock](https://github.com/golang/mock)

Update store mocks

```sh
cd [togo path]/internal/store
mockgen --source=./querier.go -destination=./mocks/querier.go -package=mocks
```
