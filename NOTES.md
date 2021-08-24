
### Running

```bash
    go mod tidy
    
    # test
    go test -v ./internal/storages/sqlite/
    
    # running
    go run main.go
```

### Things I Fixed
- passwords should be hashed
- service package imports sql and makes use of sql.NullString

### Things To Fix if I had time
- service handlers can use CORS and auth via middlewares, perhaps use labstack/echo framework
- use pkg/errors to wrap errors and make error messages more informative and easier to identify in the code and debug
- either only use the name Todo or Task, don't mix the two words. Fix sqllite vs sqlite
- ValidateUser should return an error as well as a bool in case the db has an error
- make a Store interface, and implement it in postgres as well.