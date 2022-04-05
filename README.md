### Requirements
* Golang version 1.16
* SQLite3

### How to run the code locally
* Update the `absolutePath` variable in `internal/database/init`. the db and db for unit tests are in `databases` directory
* Go to the cmd directory through your command line
* Run `go run main.go`
* Run this on curl
```
curl --request POST \
  --url http://localhost:8080/tasks \
  --header 'Content-Type: application/json' \
  --data '{
	"username": "jmramos",
	"name": "Buying Milk from Grocery Store"
}'
```
* You can use the following usernames for seperate users: `jmramos (limit of 5 todo's / day)`, `test1 5 limit`, `test2 5 limit` `test3 1 limit`

### How to run the tests locally
* Update the `internal/test/util` to your own directory. the db and db for unit tests are in `databases` directory
* Run `go test ./...` in the project directory to run all tests

### What do i like about my solution?
* Using SQLite would not require the reviewers to install databases on their local
* Easy testing of functions using tests
* Seperation of concerns

### What else do you want us to know about however you do not have enough time to complete?
* I could have created an endpoint for creating users.
* Add test on the handler (controller level test)
