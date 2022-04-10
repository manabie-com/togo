Assume that we already have Golang installed

### 1. How to run your unit tests locally?
`go test -v $(go list ./...)`


### 2. How to run your code locally?
`go run main.go`

### 3. A sample “curl” command to call your API
#### Add task
`
curl --location --request POST 'http://localhost:8080/tasks/' 
--header 'Content-Type: application/json' 
--data-raw '{"email": "email@domain.com", "task": "task detail"}'
`
#### Set Limit
`
curl --location --request POST 'http://localhost:8080/config/'
--header 'Content-Type: application/json'
--data-raw '{"email": "email@domain.com", "limit": 3 }'
`

### 4. What do you love about your solution?
- Using golang and sqlite so we can run it easily without external installation
- Can set limit task per day of each user and this limit can be different for each day
- If we don't set limit before add task, config will be created automatically with default value
- Using Interfaces for handlers and services help us generate mock easier

### 5. What else do you want us to know about however you do not have enough time to complete?
- I only wrote unittest for some basic usage, the others are defined but skipped with reason 
- I don't check uniqueness of task per user per day
- Return concise response for add new task when exceed limit per day