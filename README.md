
# Bee's Todo
Backend APIs for simple todo app

## How to run it locally?
### Requirements
- Golang 1.7+
- PostgreSQL (in ./env/postgres)

### Setting config.json file
>```
>./config.json
>```

### Run Script Create or Update table
>```go
>go run .\env\postgres\create_db_scripts\main.go 
>```

### Run
>```go
>go run main.go
>```

## Test
### "Curl” command to call your API
>```
>./curl_test.cmd
>```

### Run unit test
>```go
>go test .\src\biz\biz_todo_test.go -v
>```

### Run integration test
>```go
>go test -tags=integration .\integration_test -v -count=1
>```

## Question
- What do you love about your solution?
=> Everything is clean, easy to extend.

- What else do you want us to know about however you do not have enough time to complete?
=> Smooth things out, optimize performance, and write documents

## Docs
>```
> See directory ./docs
>```