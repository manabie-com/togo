### Q&A

#### What do you love about your solution?
- The reasons:
  - Independent Layer (Transport, Repository, Storage)
  - Flexible to change as Storage (MySQL, MongoDB, ...) and Transport (Rest, GRPC)
  - Easy to maintain
  - Easy to reusable
  - Cause Independent Layer so it's easy to unit test 

### How to run project

#### Setup Database
- Install MySQL
```sh
docker run --name mysql -e MYSQL_ROOT_PASSWORD=password123 bitnami/mysql:latest
``` 
- Set MySQL's URI (MYSQL_URI) in .env file

#### Build
```sh
go build
```

#### Auto migration (Create Relevant Table To Adding Data)
```sh
./togo auto-migration
```

#### Start server
```sh
./togo server
```

#### Run unit-test
```sh
cd ./test/${FuncName}
go test
```

### Curl

- Create User:
```sh
curl --location --request POST 'http://localhost:8080/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "User Test 2",
    "email": "usertest2@gmail.com"
}'
```

- Update User Config:
```sh
curl --location --request PUT 'http://localhost:8080/users/6/tasks/update' \
--header 'Content-Type: application/json' \
--data-raw '{
    "max_task": 7
}'
```

- Create Task for User:
```sh
curl --location --request POST 'http://localhost:8080/users/tasks/add' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": 6,
    "tasks": [
        {
            "name": "Task 1"
        },
        {
            "name": "Task 2"
        },
        {
            "name": "Task 3"
        }
    ]
}'
```
