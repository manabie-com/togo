# ToGo

### Requirements

- Implement one single API which accepts a todo task and records it
  - There is a maximum **limit of N tasks per user** that can be added **per day**.
  - Different users can have **different** maximum daily limit.
- Write integration (functional) tests
- Write unit tests
- Choose a suitable architecture to make your code simple, organizable, and maintainable
- Write a concise README
  - How to run your code locally?
  - A sample “curl” command to call your API
  - How to run your unit tests locally?
  - What do you love about your solution?
  - What else do you want us to know about however you do not have enough time to complete?

### Features

- [x] Login
- [x] Create task

### Main technologies & Libraries

- [x] Docker.
- [x] PostgreSQL.
- [x] Redis.
- [x] Gin.
- [x] Gorm.
- [x] [testify](https://github.com/stretchr/testify).
- [x] [gomock](https://github.com/golang/mock).

### Project structure

```
|- app: Main app directory
|  |- config: load env variables into config struct
|  |- constant: app constants
|  |- domain
|     |- <domain_name>: Domain package
|        |- <domain_model>.go: Domain model(s).
|        |- <domain_model>_repository.go: Domain repository interface.
|  |- errs: app-specific libraries for handling error
|  |- gingo: gin-specific libraries
|  |- migration: migration files
|  |- pkg: shared packages
|  |- test: test package
|  |- wire: dependency injection code
|     |- wire.go: Wire DI file.
|     |- wire_gen.go: Wire DI generated file.
|  |- app.env.example: Template file for `app.env(.{environment})`.
|  |- app.go: Main app
|  |- <domain_name>: Domain package (app-specific)
|     |- controller: (Delivery layer in Clean Architecture)
|        |- <protocol>: http/amqp/...
|           |- <domain_name>_controller.go: Controller
|     |- dto: Data Transfer Objects
|     |- mock: contains Go mock generated files
|     |- repository:
|        |- <domain_name>_repository.go: Implementation for repository
|     |- service: (Use case layer in Clean Architecture)
|        |- <domain_name>_service.go: Interfaces and implementations for service
|  |- Makefile: Makefile (use for development).
|  |- Makefile_test: Makefile (use for testing).
|- .husky.yaml: Husky configs.
```

Notes:

- For large-scale projects, you can split service interface into 2 parts: domain service & app service.
- Because this is only an assignment project at this time, there are something might be missed.

### Instructions

#### Coding rules

- Use singular form for names. Ex: `model`, `user`, `task`.
- Interface name for Repository and Service starts with `I`. Ex: `IUserRepository`, `ITaskService`.
- Receiver var for repository is `r`.
- Receiver var for service is `s`.
- Receiver var for controller is `ctrl`.

#### How to run this project?

1. Clone/Pull
2. Install prerequisite tools.

```shell
cd ./app && make prepare
```

##### Development

> **Prerequisite:** Run steps 1 & 2.

3. Open `app.env` and update values. Example:

```
APP_NAME=togo

LOG_LEVEL=debug

HOST=localhost
PORT=8080

DB_DRIVER=postgres
DB_HOST=127.0.0.1
DB_PORT=5433
DB_NAME=todo
DB_USERNAME=postgres
DB_PASSWORD=postgres

REDIS_HOST=127.0.0.1
REDIS_PORT=6380
REDIS_PASSWORD=p4ssw0rd

# Second
TOKEN_TTL=3600
```

4. Init docker containers

```
make docker_up
```

5. Apply migrations

```
make migup
```

6. Start development server

```
make dev
```

##### Unit Test && Integration Test

> **Prerequisite:** Run steps 1 & 2.

8. Run

```shell
make -f ./Makefile_test prepare
```

or

```shell
cd ./app && make -f ./Makefile_test prepare
```

then open `app.env.test` and update values. Example:

```
APP_NAME=togo

LOG_LEVEL=debug

HOST=localhost
PORT=8080

DB_DRIVER=postgres
DB_HOST=127.0.0.1
DB_PORT=5432
DB_NAME=todo
DB_USERNAME=postgres
DB_PASSWORD=postgres

REDIS_HOST=127.0.0.1
REDIS_PORT=6379
REDIS_PASSWORD=p4ssw0rd

# Second
TOKEN_TTL=300
```

8. Run tests

```shell
make -f Makefile_test test
```

9. Run tests with coverage

```shell
make -f Makefile_test coverage
```

Important notes for step 3 and 7: **Change value of `DB_PORT` and `REDIS_PORT` if there is any port conflict**.

##### Acceptance Test

> **Prerequisite:** Run steps from 1 to 6.

10. Run
```shell
make execdb
```

11. Run SQL command in the PostgreSQL shell, change value of username if duplicates.
```sql
INSERT INTO "user" (username, password, max_daily_task, created_at) VALUES ('ansidev', '$2a$12$iQwI4Neo.AuEaJaVnhpZPebczadhmDx0rYhex06is9dT5yz.LaD8G', 1, CURRENT_TIMESTAMP);
```

**Note:** Password is
```
5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8
```

12. Run SQL command in the PostgreSQL shell to check if the user was created
```sql
SELECT id, username, password, max_daily_task, created_at, updated_at from "user" WHERE username='ansidev';
```

13. Exit the PostgreSQL shell
```sql
exit
```

14. Login via API

```shell
curl --location --request POST 'http://localhost:8080/auth/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "ansidev",
    "password": "5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8"
}'
```

You should see the output looks like

```
{"token":"b6581984-a2be-4baa-bafe-c53b176243f1"}
```

Copy the token value for the next step.

###### Test cases

**Note:** Replace `<token>` (if exists) with your token from the above step.

1. Create task via API without token should return HTTP status code `401`.

```shell
curl -i --location --request POST 'http://localhost:8080/task/v1/tasks' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Sample task"
}'
```

2. Create task via API with invalid authorization header should return HTTP status code `400`.

```shell
curl -i --location --request POST 'http://localhost:8080/task/v1/tasks' \
--header 'Authorization: <token>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Sample task"
}'
```

3. Create task via API with an expired/fake token should return HTTP status code `401`.

```shell
curl -i --location --request POST 'http://localhost:8080/task/v1/tasks' \
--header 'Authorization: Bearer invalid_token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Sample task"
}'
```

4. Create task via API with body request missing field `title` should return HTTP status code `400`.

```shell
curl -i --location --request POST 'http://localhost:8080/task/v1/tasks' \
--header 'Authorization: Bearer <token>' \
--header 'Content-Type: application/json' \
--data-raw '{}'
```

5. Create task via API with valid body request should return HTTP status code `201`.

```shell
curl -i --location --request POST 'http://localhost:8080/task/v1/tasks' \
--header 'Authorization: Bearer <token>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Sample task"
}'
```

Run `make execdb`, then run SQL command to check if the task is exists.

```shell
make execdb
```

```sql
SELECT id, title, user_id, created_at, updated_at from "task";
```

Exit PostgreSQL shell
```
exit
```

### My thoughts

I love this solution because it is a working solution at least. I applied almost all of my related knowledge to do this assignment. The project structure is inspired by the Clean Architecture and uses the template from one of my recent public projects [gin-starter-project](https://github.com/ansidev/gin-starter-project). During the implementation of this solution, I realized some mistakes that I can apply back to the starter project.

Known issues:
- Some test cases for the Auth API might be missed because this is not the main goal of this assignment.
- The error codes are declared using `iota` to make sure it is unique between services and less code. However, the drawback is that the code can be changed by changing the order of a line (a mistake from another developer who doesn't know my idea) instead of appending new lines. We should have test cases for this issue or a better solution.
- The return message is a snake_case string because the implementation is not complete.
  - One of my previous use cases is support i18n for the API:
    - The client-side submits the expected language via the HTTP Header `Accept-Language` or the query parameter `lang`.
    - The server extracts the language using middleware.
    - For building the response: receive the message key and language and use a translator library to get and return the i18n string.
- Production build files are not included because it does not the main goal of this assignment so I skipped adding them.
- If I have more time, I want to add more test cases and CI/CD definition files for automated tasks.
