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
1. Installed and working: Golang, Docker
2. Docker 

### Running the Code Locally
1. Pull the code to your local repository
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

### Running Unit Tests
1. Go to the project path "todo-app" and run the unit test
```
go test -v ./repository
