### Requirements
* Golang version 1.16
* SQLite

### How to run the code locally
* Update the `absolutePath` variable in `internal/database/init`. the db and db for unit tests are in `databases` directory
* Go to the cmd directory through your command line
* Run `go run main.go`

### How to run the tests locally
* Update the `internal/test/util` to your own directory. the db and db for unit tests are in `databases` directory
* Run `go test ./...` in the project directory to run all tests
