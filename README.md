## Technical Assignment for Backend Engineer
### Overview
This is a simple backend for a good old todo service, right now this service can handle login/list/create simple tasks. \
To make it run: 
- `go run main.go`
- Import Postman collection from `docs` to check example \
- Run testcase: `go test -v .\test\...`
### What I do ?
- Separate the layer according to the follow structure: Router -> Delivery -> Usecase (Service) -> Respository (Storage) -> Model
```
.
├── delivery
├── model
├── respository
├── usecase
├── utils
```
- Make this code more DRY
- Write unit test for service and respotory layer
### What is missing ?
- Change from using `SQLite` to `Postgres` with `docker-compose`
- Integration tests
### Things I want to improve
- More clean code with Clean Architecture of Uncle Bob
- Research more about testing for unit test and integration test in golang
- Handle error \
...
