# Trình Đại Phúc

## Requirements

- Golang Version: 1.16+
- Docker Version: 17.06+
- Docker Compose Version: 1.21+

## Run Project:

```shell
make dev-up
```

## Sample curl command:

- Sign up:

```shell
curl -X POST \
  -H "Content-Type: application/json" -d '{"username":"daiphuc", "password":"123456", "taskLimit":5}' \
  http://localhost:8080/api/v1/signup
```

- Login to get jwt token:

```shell
curl -X POST \
  -H "Content-Type: application/json" -d '{"username":"daiphuc", "password":"123456"}' \
  http://localhost:8080/api/v1/login
```

- Create task with jwt token login:

```shell
curl -X POST \
  -H "Content-Type: application/json" -H "Authorization: Bearer <token>" \
  -d '{"title":"Task 1", "description":"Task 1 description"}' \
  http://localhost:8080/api/v1/tasks
```

- Get task with jwt token login:

```shell
curl -X GET \
  -H "Authorization: Bearer <token>" \
  http://localhost:8080/api/v1/tasks/1
```

- Get all tasks with jwt token login:

```shell
curl -X GET \
  -H "Authorization: Bearer <token>" \
  http://localhost:8080/api/v1/tasks?page=1&limit=10
```

- Update task with jwt token login:

```shell
curl -X PUT \
  -H "Content-Type: application/json" -H "Authorization: Bearer <token>" \
  -d '{"name":"Task 2", "content":"Task 1 description"}' \
  http://localhost:8080/api/v1/tasks/2
```

- Delete task with jwt token login:

```shell
curl -X DELETE \
  -H "Authorization: Bearer <token>" \
  http://localhost:8080/api/v1/tasks/1
```

## Testing:

- Run test:

```shell
make test
```

- Run coverage:

```shell
make coverage
```
