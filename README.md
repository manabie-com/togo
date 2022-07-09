### Requirements
  - Implement one single API which accepts a todo task and records it
  - There is a maximum **limit of N tasks per user** that can be added **per day**.
  - Different users can have **different** maximum daily limit.

### Choose a suitable architecture to make your code simple, organizable, and maintainable

## Ben Johnson purposes 4 principles to structure our code.

  1. Root Package is for domain types
  2. Group subpackages by dependency
  3. Use a shared mock subpackage
  4. Main package ties together dependencies

- Ref links: https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1

## How to run your code locally?

  1. Download and Install PgAdmin4 tools (postgresql): https://www.postgresql.org/download/ 
  2. Open queries.sql file, then Excute from "--> START" to "--< END" to create "users", "todos" tables
  3. Download/Update go modules
    ```bash
    go mod tidy
    ```
  4. Start webservice
    ```bash
    go run cmd/webservice/main.go
    ```
  5. Send request with curl or Postman
    ```bash
    curl -XPOST 'http://localhost:3000/api/lawtrann/todos' -H 'Content-Type: application/json' -d '{"description":"todo something"}'
    ```

## A sample “curl” command to call your API
    ```bash
    curl -XPOST 'http://localhost:3000/api/lawtrann/todos' -H 'Content-Type: application/json' -d '{"description":"todo something"}'
    ```

## How to run your unit tests locally?
- Unit Test
 ```bash
 go test -v ./services/
 ```

- Test cases
 - Testcase#1: Create NewUser with Description:"Todo 1"
    ```bash
    curl -XPOST 'http://localhost:3000/api/newuser/todos' -H 'Content-Type: application/json' -d '{"description":"Todo 1"}'
    ```
 - Testcase#2: add Description:"Todo 2" for NewUser on testcase#1
    ```bash
    curl -XPOST 'http://localhost:3000/api/newuser/todos' -H 'Content-Type: application/json' -d '{"description":"Todo 2"}'
    ```
 - Testcase#3: add bunch of Description:"Todo n" until getting "you have reached the limit of adding todo task per day" message
    ```bash
    curl -XPOST 'http://localhost:3000/api/newuser/todos' -H 'Content-Type: application/json' -d '{"description":"todo n"}'
    ```

## What do you love about your solution?
- Testability
  - With such a pluggable system, we can test the functionality of each layer separately by injecting a mock version of the dependent layers
- Clear separation between layers
  - In our domain package we can see an interface for each layer in our application. This helps us to have a clear boundary between each layer.
  
## What else do you want us to know about however you do not have enough time to complete?
- I haven't dealt with the system's logging yet, nor have I used middleware in this project.
- Using route open source like Chi to optimize router.
- Using docker to run postgredb instead of creating manually through queries.sql file.
