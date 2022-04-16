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

### System Requirements
Installed and working: Golang, Docker 

### Architecture/Design
Code was written in a REST API architecture. The code structure may not be industry standard as this is the first time I have created a REST API from scratch and first time using the Go language and PostgreSQL. The controller kind of acts as both a controller and the service itself.

### What do you love about your solution?
1. I love how each code function is separated through files according to their purpose, a practice that I prefer in coding. This allows an easier to maintain and read as contrary to the one single file. 
2. My solution also features a full CRUD of the todo task
3. Task limit per user is configurable in the database. Based from my experience as an application support, simple configurations such as the task limit per user and configurations that are expected to change every now and then, is better implemented in the database to be easily configured as requested by the business/client.
4. I now appreciate what a Test Driven Development means and the benefit associated with it. Should I make changes to my code, I would be able to test its functions immediately and see if the expected output is still achieved/attained.

### What else do you want us to know about however you do not have enough time to complete?
1. As someone who's out of the practice in IT development, I am proud that I was able to create a REST API in a limited time (~a week).  
3. I think there are still a lot of things to improve on how the base code and the tests are implemented, in which I hope to learn more a lot should I get hired. The things to improve that I thought of as of now are Golang coding standards, and more use of structs and interfaces. 
4. Another thing is the environment variables of the PostgreSQL connection string be separated into username, password, hostname, port, etc. While the password, ideally, should be stored and fetched from a password vault application like CyberArk as best practice and should not be stored in a simple environment variable.
5. TODO: Implement an user authentication mechanism on the API that uses JWT as the authentication token. Another one is to implement a password encryption/decryption mechanism to be used along side the authentication mechanism wherein an encrypted password will be stored instead in the database.

### Running the Code Locally
1. Pull the code from the repository to your local machine
```
git clone https://github.com/qgdomingo/todo-app.git
```
2. TODO: setup docker compose postgres container
3. Setup environment variable POSTGRES_DB_URL, no need to change on the link
```
POSTGRES_DB_URL=postgres://todo_user:secret@localhost:5432/tododb
```
4. Go the project path "todo-app" and run the code
```
go run main.go
```
5. API will run in port 8080

### Running Integration Test
Go to the project path "todo-app" and run the test. PostgreSQL docker must be up and running
```
go test -v ./repository
```

### Running Unit Test
Go to the project path "todo-app" and run the test.
```
go test -v ./controller
```

### Running Sample curl Commands
1. POST request - Create a task
```bash
curl -X POST -H "Content-Type: application/json" \
    -d '{"task_title": "Test Task from Curl", "task_desc": "Test Task Description from Curl", "created_by": "qgdomingo"}' \
    localhost:8080/todo/create
```
2. GET request - Fetch all tasks 
```bash
curl -X GET localhost:8080/todo/fetch
```
3. GET request - Fetch tasks from specific username 
```bash
curl -X GET localhost:8080/todo/fetch/usertask/qgdomingo
```
4. GET request - Fetch specific task using id (get the task id of the "Test Task from Curl" task and replace on the link on the curl command below)
Sample: localhost:8080/todo/fetch/3
```bash
curl -X GET localhost:8080/todo/fetch/id
```
5. PUT request - Update a task (use the same id previously)
Sample: localhost:8080/todo/update/3
```bash
curl -X PUT -H "Content-Type: application/json" \
    -d '{"task_title": "Test Task Update from Curl", "task_desc": "Test Task Description Update from Curl", "created_by": "qgdomingo"}' \
    localhost:8080/todo/update/id
```
6. DELETE requerst - Delete a task (use the same id previously)
Sample: localhost:8080/todo/delete/3
```bash
curl -X DELETE localhost:8080/todo/delete/id
```
