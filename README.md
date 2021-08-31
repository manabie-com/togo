### Notes
This is a simple backend for a todo service, right now this service can handle signup/login/list/create simple tasks:

### Prerequisites
You have to setup on your machine:
- [Golang](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/engine/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- Prepare Make tool by yourself, it depends on which environment you do, Mac or Linux

Your dependencies have been installed yet? Boot up database first:
- `make docker.start.components`
In case you want to take database down:
- `make docker.stop`

### Boot up server
To make it run:
- `make`
- `make run`
- Import Postman collection from docs to check example

### Running tests
In order to execute test cases, the project has provided integration and unit tests.
Performing integration tests as below:
- `make test.integration`
Performing unit tests as below:
- `make test.unit`

---

### Functional requirement:
Right now a user can add as **limit N tasks per day** <-- done.  
For example, users are limited to create only 5 tasks only per day, if the daily limit is reached, return 4xx code to the client and ignore the create request.

### Non-functional requirements:
- [ ] Write integration tests for this project <-- partial done (1 endpoint for user, 1 endpoint for task)
- [ ] Make this code DRY <-- tried my best
- [ ] Write unit test for the services layer <-- done
- [ ] Change from using SQLite to Postgres with docker-compose <-- done
- [ ] This project includes many issues from code to DB structure, feel free to optimize them <-- tried my best
- [ ] Write unit test for storages layer <-- done
- [ ] Split services layer to use case and transport layer <-- tried my best

### Some words
Even time is not too much, though. I would take it simple as much as possible in this test, from README, files and directory tree, to SQL scripts for easy booting up database.
Especially, split services and transport layer, I would love to make it more abstract if I have more time, in case of adding more services without editing on transport layer.
For examples, beside user and task services, if business extended with 'team' services, developers have to declare the services in transport layer -> need to improve.

#### Sequence diagram
![auth and create tasks request](https://github.com/manabie-com/togo/blob/master/docs/sequence.svg)
