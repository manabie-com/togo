# Duong Thanh Tin - Test Manabie

## Togo Respository

### Features
The repository have a few main features
```
- Login
- Create user
- Add task by user and create date
```

### Diagrams

1. #### Sequence Diagram For Flow Feature
- Feature Login
![Sequence](https://raw.githubusercontent.com/DuongThanhTin/togo/master/documents/Flow-Login.svg)

- Feature Create User
![Sequence](https://raw.githubusercontent.com/DuongThanhTin/togo/master/documents/Flow-CreateUser.svg)

- Feature Add Task
![Sequence](https://raw.githubusercontent.com/DuongThanhTin/togo/master/documents/Flow-AddTask.svg)

2. #### ERD Diagram

![ERD](https://raw.githubusercontent.com/DuongThanhTin/togo/master/documents/ERD.svg)

###  Structure Project

```
- cmd --> Main applications for this project.
  |- middlewares --> Middlewares for this project
- constants --> Contain variable common
- db --> Data for migrations
  |- migrations -> You can migration up or down data
- documents --> Contain detail documents for this project
- integrationtest --> Run integration test
- internal
  |- api --> You can create different output commands like Api rest, web, GRPC or any other technology.
    |- handlers --> Contain API
      |- common --> API for common
      |- tasks --> API for task
      |- users --> API for user
    |- routes --> Make create route for API
  |- driver --> Config connection to database
  |- models --> Application models
  |- pkg --> Make create function to use common
    |- id --> Make create uuid
    |- responses --> Make create many response data
    |- repositories --> Repositoryies will action to database (CRUD)
      |- authorization --> Repository for auth
      |- task --> Repository for task
      |- user --> Repository for user
  |- usecases --> Usecases to implement action for application
    |- authorization --> Usecase for auth
    |- task --> Usecase for task
    |- user --> Usecase for user
```

### How to start

#### Prepare

- Install golang [golang](https://go.dev/doc/install)
- Install postgres [postgreSQL](https://www.postgresql.org/download)
- Install migrate If you want to migration [migration](https://github.com/golang-migrate/migrate)

  . MacOS

  ```bash
  $ brew install golang-migrate
  ```

  . Windows

  Using [scoop](https://scoop.sh/)

  ```bash
  $ scoop install migrate
  ```

  . Linux

  ```bash
  $ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
  $ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
  $ apt-get update
  $ apt-get install -y migrate
  ```
#### Start local
1. Clone repository
```bash
git clone https://github.com/DuongThanhTin/togo.git
cd togo
cp .env.example .env
```

2. Create Database postgres with your config

3. Change data variable in file .env with your config on step 2

4. Migration database with terminal (If you want to migrate)
```bash
make migration-up
```
5. Run main.go
```bash
go run cmd/main.go
```

#### Install
```bash
go install ./...
```

#### If you want to integration test after run projection
```bash
go test -tags=integration ./integrationtest -v -count=1
```

#### If you want to unit test test after run projection
```bash
go test ./internal/...
```

#### If you want to remove table on database
```bash
make migration-down
```

### Flow

You can see all flow in folder `documents`. I draw three pictures for call API and one picture for ERD.

#### Endpoints

I make create three endpoints.
```bash
POST /logins -> Login
POST /tasks  -> Create task
POST /users  -> Create user
```

#### Authentication token
I will create token and set data of for token in `Cookie`. You must login first if you want to create task with endpoint `POST /tasks`. The token will be generated and set in `Cookie` with key `token` and have data `user_id` and `max_task_per_day`. The task will create with data user login.

#### Call API
1. You can create user with endpoint `POST /users`. In addtion, I have validate `username`, `password` not empty in request. If `username` is exists, you can't create user.
Request JSON:
```bash
{
	"username":"user-manabie-1",
	"password": "123456",
	"max_task_per_day": 10
}
```

2. You can login with with endpoint `POST /login`. I have validate `username` and `password` not empty in request. First of all, you must create user if you want to login with user as above.
After login, Token will be generated and set in `Cookie` with key `token` and have data `user_id` and `max_task_per_day`.
If you run terminal with `make migration-up`. I create example user for you.
User example: `id: firstUser, username: manabie, password: example, max_task_per_day: 5`
Request JSON:
```bash
{
	"username":"manabie",
	"password": "example",
}
```
I will this response which have token for you. You can easily set this token in `Cookie` with key `token` if you call API with `curl`.

3. You can create task with endpoint `POST /tasks`. First of all, you must login if you want to create task. I have create token and set it in `Cookie`. This token have the information of `user` such as:
`user_id` and `max_task_per_day`. After that, you call API `POST /tasks`, i will parse this token to get data from this token so i need you login. I have validate number task of day if the number task of day is more than `max_task_per_day` of user, you can't create task.
Request JSON:
```bash
{
	"content":"New task 1",
}
```

More detail:
- Field `userID` in `task` get from token
- Field `createDate` in `task` get from date call API

#### Collection postman
I make create collection postman folder `collection_postman`. You can easily see the endpoint and request json for each API.

#### CURL Example
1. Create User
```bash
curl -XPOST -d '{ "username":"manabie-user-3", "password": "123456", "max_task_per_day": 10 }' 'http://localhost:8000/users'
```

Response:
```bash
{"data":{"username":"manabie-user-3","password":"123456","max_task_per_day":10},"message":"Success","status":201}
```

2. Login
```bash
curl -XPOST -d '{ "username":"manabie", "password":"example" }' 'http://localhost:8000/login'
```
Response:
```bash
{"data":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTY1NTE4NDcsIm1heF90YXNrX3Blcl9kYXkiOiI1IiwidXNlcl9pZCI6ImZpcnN0VXNlciJ9.qwHLe5Nd1lxUJlHPh3LtJUsX68ML2foMv_yjD4x5VJY","message":"Success","status":200}
```

3. Create Task
```bash
curl -XPOST -b 'token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTY1MjIzMzYsIm1heF90YXNrX3Blcl9kYXkiOiI1IiwidXNlcl9pZCI6ImZpcnN0VXNlciJ9.RvmCCNF5vOloXQmyZEqAUcZtxQN4lN9_qhkSm4vByOE' -d '{ "content":"New task 2" }' 'http://localhost:8000/tasks'
```
Response:
```bash
{"data":null,"message":"Success","status":201}
```