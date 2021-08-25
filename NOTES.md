# Test submission from ds0nt

Thanks for taking the time to interview me and run my code here.

I hope you enjoy it. Check the diffs to see all of my changes.

Cheers!

### Running

```bash
    go mod tidy
    

    # running tests for sqlite and postgres stores
    docker-compose -f docker-compose-test.yml up # starts test postgres
    go test -v ./internal/storages
    
    # running the service
    docker-compose up # starts postgres
    go run main.go
```

### Testing in Postman

1. Run Login
1. Copy JWT token into headers of other requests
1. CreateTask can run 5 times
1. Change the date to today's date and run List to see all 5 todos.



### Things I Fixed
- todos limited by users.max_todo
- added some new methods to storages/sqlite
- wrote tests for storages/sqlite package
- bcrypt hash passwords
- 400 status code if max_todos reached
- service no longer imports the sql package
- implemented postgres store
- altered storages tests to cover both sqlite and postgres store implementations

### Things To Fix if I had time
- service handlers can use CORS and auth via middlewares, perhaps use labstack/echo framework
- use pkg/errors to wrap errors and make error messages more informative and easier to identify in the code and debug
- either only use the name Todo or Task, don't mix the two words.
- ValidateUser should return an error as well as a bool in case the db has an error