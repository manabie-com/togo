### Run code on local environment
go run main.go

### Sample Curl to call the API
curl --location --request POST 'http://localhost:9000/user/1/todo' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "To Do 1",
    "detail": "This is a sample to do list",
    "remind_date": "2022-04-03"
}'

### Run all unit tests on local environment
go test ./...

### Solution
A simple solution that is not involve in DB, only using a simple user struct (model) to demonstrate the core logic of the todo task recording, limitation of the maximum daily tasks based on users, and reseting the daily limit after the day has passed.

### Future Implementation
Implement database connection to save users & todo tasks.