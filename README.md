# Manabie's interview test solution

## How to run your code locally?

You need to install Go version 1.18 and an IDE/editor such as Goland or VSCode to run the project.

```bash
cp .env.template .env # create `.env` file
make run # or `go run main.go`
```

## A sample “curl” command to call your API

- Register a new user

```bash
curl --location --request POST 'localhost:8000/api/v1/register' \
--form 'email="<user-email>"' \
--form 'password="<user-password>"'
```

- The server uses JWT for authentication, so you need to log in with `email` and `password` to get a JWT token.

```bash
curl --location --request POST 'localhost:8000/api/v1/login' \
--form 'email="<user-email>"' \
--form 'password="<user-password>"'
```

- Use the token you get from API above to create a task with the command below.

```bash
curl --location --request POST 'localhost:8000/api/v1/tasks' \
--header 'Authorization: Bearer <your-access-token>' \
--form 'title="Task 1"' \
--form 'description="squat 100 kg"'
```

## How to run your unit tests locally?

```bash
make test # or `go test --cover ./...`
```

## What do you love about your solution?

### Isolation and technology independence

The clean architecture isolates the code then we can use any lib we want without worrying breaking the current codebase.

For instance, if we want to use MongoDB instead of MySQL, we only need to update model folder or if we want to user `fiber` over `gin`, we only need to update the transport folder.

The API flow: Transport layer -> Business layer -> Repository layer -> Storage layer.

- Transport layer: parse data from request/socket
- Business layer: handle business logic
- Repository layer:
- Storage layer: integrate with DB

### Testing

Because each layer is independence, so we can easily create mock input for each layer.

### Security

Currently, the server has 5 APIs.

- POST `/register`: register a new user
- POST `/login`: user logs in to get JWT token
- POST `/tasks`: create a new task for current logged-in user

The server generates fake uid (a string) to hide real integer id in MySQL. That trick will enhance security.
These 2 APIs are used for development environment only.

- POST `/encode-uid`:
  - Encode uid receives real id and database type (an integer constant that assigned for each SQL table), then return fake uid.
  - E.g: id: 16, db_type: 2 -> fakeId: 3w5rMJ8raFkfXS
- POSt `/decode-uid`:
  - Decode uid receives fake uid then return real id and database type
  - E.g: fakeId: 3w5rMJ8raFkfXS -> id: 16, db_type: 2

## What else do you want us to know about however you do not have enough time to complete?

- [ ] Increase the test coverage up to 80%
- [ ] Write integration test
- [ ] Handle run some APIs in development environment only
- [ ] Handle CI/CD
