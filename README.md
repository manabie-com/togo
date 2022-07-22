# Version 0.2

- Simple http server using `gorilla mux`
- Store data in MySQL database
- API contain user_id as path parameter and task name in its body
- Test suite for handler

Insert task

```sh
curl -X POST localhost:8080/user/1/task -d '{"name": "do exercise"}'
```

Get task

```sh
curl -X GET localhost:8080/user/1/task
```

## Local set up MySQL

First time initialize mysql container

```sh
docker run --name mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=gouser -e MYSQL_PASSWORD=gopassword -e MYSQL_DATABASE=godb -p 3306:3306 mysql:latest
```

Access to MySQL

```sh
mysql -h0.0.0.0 -P 3306 -u gouser -p
```

Run [migration](./internal/store/migrations/)

```sql
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

## Mock

Install [golang mock](https://github.com/golang/mock)

Update mocks

```sh
cd [togo path]/internal/store
mockgen --source=./querier.go -destination=./mocks/querier.go -package=mocks
```
