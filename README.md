## Requirements
- Go 1.17

## How to run the code locally

### Set the env variables in a .env file
This is an example of the needed environment variables
```
APPS_PORT=8080
DB_MONGO_URI=mongodb://localhost:27017
DB_MONGODB_NAME=test
```

### Run this commands
##### Clone the repository
####
```bash
git clone https://github.com/kier1021/togo.git
```
##### Change the directory 
####
```bash
cd togo
``` 
##### Run the go run command to run the server
####
```bash
go run main.go
```

## Sample "curl" commands
### Create User
```
curl -X POST http://localhost:8080/user \
    -H "Content-Type: application/json" \
    -d "{ \"user_name\": \"Test\", \"max_tasks\": 3 }"
```

### Get Users
```
curl -X GET http://localhost:8080/users 
```

### Add Task To User
```
curl -X PUT http://localhost:8080/user/task \
    -H "Content-Type: application/json" \
    -d "{ \"user_name\": \"Test\", \"title\": \"Task 1\", \"description\": \"My first task\" }"
```

### Get Task Of User
##### With given ins_day
#### 
```
curl -X GET "http://localhost:8080/user/task?user_name=Test&ins_day=2022-02-18"
```

##### Without given ins_day, the default is date today
#### 
```
curl -X GET "http://localhost:8080/user/task?user_name=Test"
```

## How to run unit test

```
go test -v ./api/services
```

## What do I love about my solution
- The separation of concerns of each functionalities
    - Controller for handling the http requests
    - Service for handling the business logic
    - Repository for handling the database transactions
    - Models for representing the entities
- Adding a repository interface makes the unit testing easy using mock data
- Dependencies were injected to unit test the services
- I used the table-driven pattern in unit test
- Used some libraries that can be useful for other future projects

## What to improve?
- Add an integration test
- The modelling of entities may be improved
