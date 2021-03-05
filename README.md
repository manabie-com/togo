# Overview
A simple backend service to create todo tasks, written in Go.
There's only one user, firstUser@example. This user can only create N tasks per day according to their database value. Once reached, the server returns 400 and ignores the request.
# To make it run
- `go run main.go`
- Import Postman collection from `docs` to check example
- For GET `tasks` and POST `tasks`, make sure to add `Authorization = $JWT` token obtained from GET `login` to the request Header
- For GET `tasks` endpoint, try `2021-03-05` in addition to the example date

# Testing
Currently there's only one test for the login flow. Regardless, start the server and run the command below to test `services`
```
 go test github.com/manabie-com/togo/internal/services
```

# Possible improvements
- Make authentication a separate module/ package.
- Write more tests.
- Database structure has many problems:
    - Introduce auto-increment numeric primary keys instead of text-based. This reduces overhead when indexing & querying.
    - Constraint ids (i.e usernames) to be unique. 
    - Use hash and salt for passwords. Storing passwords in plaintext is a very bad idea.