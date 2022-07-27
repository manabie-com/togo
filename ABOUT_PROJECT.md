### How to run
- Start postgres database
```
make start-db
```
- Run database migration
```
make migrate-up
```
- Start server 
```
make run
```

### How to run test locally:
```
make test
```

### A sample "curl" command to call my API:
```
curl --request POST \
  --url http://localhost:8080/todo \
  --header 'Content-Type: application/json' \
  --data '{
  "userID":2,
  "name":"todo1",
  "content":"content"
  }'
```

### About my solution
I use Clean Architecture. It has some advantages:
- Agnostic to the outside world
- Independent with external services
- Easier to test in isolation
- The Ports and Adapters are replaceable
- Separation of the different rates of change
- High Maintainability

So I think this is my favorite part of the solution.

### What else do I want to improve
- Use test container instead of `sqlmock` is better?
- Increase test case coverage 