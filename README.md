### Notes
This is a simple backend for a todo service, right now this service can handle login/list/create simple tasks, to make it run:
- `sudo docker-compose up`
- go to http://localhost:5051/login with username `admin@manabie.com` and password `admin` and run queries in `migration_script.sql`
- `go run main.go`
- Postman collections can be imported to run integration testing through two files using Postman Runner
	- collection file: togo.postman_collection.json
	- environment file: local.postman_environment.json

### Non-functional requirements:
- [ ] **A nice README on how to run, what is missing, what else you want to improve but don't have enough time**
- [ ] **Consistency is a MUST**
- [ ] Fork this repo and show us your development progress by a PR
- [ ] Write integration tests for this project
- [ ] Make this code DRY
- [ ] Write unit test for the services layer
- [ ] Change from using SQLite to Postgres with docker-compose
- [ ] This project includes many issues from code to DB structure, feel free to optimize them
- [ ] Write unit test for storages layer
- [ ] Split services layer to use case and transport layer


#### Sequence diagram
![auth and create tasks request](https://github.com/manabie-com/togo/blob/master/docs/sequence.svg)
