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

### Notes
- I decided to code in Go straight away upon knowing that in the interview stage, (should I be considered for the position that is :p) it will be a test of my knowledge in the Go programming language. Sadly it resulted to my output being incomplete as I raced with time learning the language and learning TDD in Go. However, it's the part that I enjoyed the most. Learning a new language that I loved a lot along the way. I'll definitely pursue this line of work wherever it may lead me. Trying out in this opportunity by Akaru/Manabie is alreay a big win for me that it helped me find something that I'll be passionate about.
- Given more time, I would've covered more use cases from the coding challenge requirements and increase the unit tests and actually create integration tests for the project. I would've employed authentication for users and api call quota.