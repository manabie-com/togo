# SINGLE TODO Api

Single API which accepts a todo task and records it
Contents:
- There is a maximum limit of N tasks per user that can be added per day.
- Different users can have different maximum daily limit.


## Feature
- Create todo task (includes: UT & IT)

## Tech
- Spring Boot
- Spring JPA
- MySQL


## Setup local machine
Install the dependencies and devDependencies and start the server.

```sh
1. Open eclipse or Spring tool suites
2. Import existing maven project
3. Create 3 database: todos and todos_test
4. Run sql file: create_tables_data.sql
5. Create tables for database todo_test before run unit test or integration test
6. Before run local you need change username and password to connect to mysql in 2 file: application.properties ( one for development, one for testing)
7. To run local: right click file "TodoController" > "Run As" > "Java Application"
```

## Curl example
```sh
curl -X POST http://localhost:9090/api/todo/tasks -H 'Content-Type: application/json' -d '{"title":"title_test","description":"ABC","userId":1}'
```
