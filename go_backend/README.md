### Dependency
- Nodemon
- postgresql

### Setup psql
- Setup psql account
  - CREATE ROLE testuser CREATEDB WITH LOGIN PASSWORD 'P@55w0rd'; 
  - CREATE DATABASE testdb;
- Run “create_table.sql” script included in folder

### How to run locally
- go to “go_backend” directory
- run via “make all”
  - for running test cases: “go test main_test.go -v”

### Sample cURL
- GET | All Task
 - curl --location --request GET 'localhost:3000/api/tasks/list'

- GET | Task Details by `ID`
 - curl --location --request GET 'localhost:3000/api/tasks/?id=e3ffa001-8e42-415a-b92b-e363d6f8f5cb'

- POST | Assign Task to User
 - curl --location --request POST 'localhost:3000/api/tasks/' \
    --header 'Content-Type: application/json' \
    --data-raw '{
      "assigned_to": "USER_1",
      "title": "Assignment",
      "description": "this is your first task"
    }'


- I made a simple approach to complete the exam
- Completed the code around 2-3 hours due to hectic work schedule