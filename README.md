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

- Login to get jwt token:

```shell
curl -X POST \
  -H "Content-Type: application/json" -d '{"username":"daiphuc", "password":"password"}' \
  http://localhost:8080/api/v1/login
```

- Create task with jwt token above:

```shell
curl -X POST \
  -H "Content-Type: application/json" -H "Authorization: Bearer <token>" \
  -d '{"title":"Task 1", "description":"Task 1 description"}' \
  http://localhost:8080/api/v1/tasks
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
