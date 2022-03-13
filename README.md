
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

### "Curl” command to call your API
>```
>./curl_test.sh
>```

### Run unit tests
>```go
>go test ./..
>```

## Question
- What do you love about your solution?
- What else do you want us to know about however you do not have enough time to complete?

>```
> See directory ./docs
>```