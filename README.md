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
![Sequence](https://raw.githubusercontent.com/DuongThanhTin/togo/master/document/Flow-Login.svg)

- Feature Create User
![Sequence](https://raw.githubusercontent.com/DuongThanhTin/togo/master/document/Flow-CreateUser.svg)

- Feature Add Task
![Sequence](https://raw.githubusercontent.com/DuongThanhTin/togo/master/document/Flow-AddTask.svg)

2. #### ERD Diagram

![ERD](https://raw.githubusercontent.com/DuongThanhTin/togo/master/document/ERD.svg)

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
      |- task --> API for task
      |- user --> API for user
    |- routes --> Make create route for API
  |- driver --> Config connection to database
  |- models --> Application models
  |- pkg --> Make create function to use common
    |- id --> Make create uuid
    |- responses --> Make create many response data
    |- repositories --> Repositoryies will action to database (CRUD)
      |- auth --> Repository for auth
      |- task --> Repository for task
      |- user --> Repository for user
  |- usecases --> Usecases to implement action for application
    |- auth --> Usecase for auth
      |- task --> Usecase for task
      |- user --> Usecase for user
```

### How to start

#### Start local
1. Clone repository
```bash
git clone https://github.com/DuongThanhTin/togo.git
```
2. Run main.go
```bash
cd togo
go run cmd/main.go
```
