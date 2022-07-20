# Version 0.1

- Simple http server
- Store data in a single csv file
- 2 APIs: one for insert task with task name, one for get task information

Insert task

```sh
curl -X POST localhost:8080/task -d '{"name": "hello"}'
```

Get task

```sh
curl -X GET localhost:8080/task
```
