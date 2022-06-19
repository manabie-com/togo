### How to run your code locally?
  - Install Redis, MongoDB locally
  - Add your redis, mongo to `storage/const.go` and `cache/const.go`.
  - From root, run `go run main.go`

### A sample “curl” command to call your API
```
curl --location --request POST 'localhost:8080/api/v1/task/record' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": "3",
    "task": "todo"
}'
```

### How to run your unit tests locally?

- TBU

### What do you love about your solution?

- Use go clean architect so that it makes my code structure readable
- Use Redis cache to control limit user, it is an interesting use case of redis
- Use MongoDB to log user data -> it is fast compared to normal RDS (read, write)

### What else do you want us to know about however you do not have enough time to complete?

- I got sick during the test time, but I still tried by best to complete it