### Requirements
- Docker Desktop and Docker Compose installed
- Run `docker-compose --env-file app.env up` **before** running unit tests or issuing `curl` commands

### Running Locally
- `go run .`
- Sample `curl` command `curl -d '{"content":"some content", "user_id":1}' -H 'Content-Type: application/json' -X POST http://localhost:5000/todo`
- Use included simple DB admin viewer to inspect changes in database after each request
  - Go to `http://localhost:8080`
  - Select PostgreSQL as DB System and use the following credentials to login
    - Server: `localhost:5432` or `127.0.0.1:5432`
    - Username: **postgres**
    - Password: **postgres**
    - DB Name: **postgres**
### Running Tests
- Run test with `go test ./tests/ -v`