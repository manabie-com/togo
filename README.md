### Requirements
- Docker Desktop and Docker Compose installed
- Run `docker-compose --env-file app.env up` **before** running unit tests or issuing `curl` commands

### Running Locally
- `go run .`
- Sample `curl` command `curl -d '{"content":"some content", "user_id":1}' -H 'Content-Type: application/json' -X POST http://localhost:5000/todo`
### Running Tests
- Run test with `go test ./tests/ -v`