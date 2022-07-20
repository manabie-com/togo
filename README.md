# Version 0.1

- Simple http server
- Store data in a single csv file
- 2 APIs: one for insert task with task name, one for get task information

Insert task

```sh
curl -X POST localhost:8080/task -d '{"name": "hello", "user_id": 1}'
```

Get task

```sh
curl -X GET localhost:8080/task
```

## MySQL

```sh
docker run --name mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=gouser -e MYSQL_PASSWORD=gopassword -e MYSQL_DATABASE=godb -p 3306:3306 mysql:latest
```
