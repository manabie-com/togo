### Prerequisites

- docker, docker-compose
- go 1.17
- mysql(client)

### Setup
To build and start the server:

```docker-compose up --build togo```

To setup the database: 

```mysql -h localhost -P3001 --protocol=tcp -utogo --password=togo togo  < togo.sql```

### Sample API 

```curl -v http://localhost:8000/users/1/tasks -X "POST" -d "{\"name\":\"tests\"}"```

### Running tests

```go test ./...```


### Notes

For quick setup, docker and docker-compose is used to handle dependencies.

For this project, I setup the structure to follow MVC principles. The structure is also loosely based on https://github.com/katzien/go-structure-examples/tree/master/domain-hex. 

The idea is to be able to replace the API layer(implemented JSON here, but can be swapped to SOAP, gRPC, etc), and database layer(implemented MySQL and mock test here) any time, without affecting the core logic.

For tests I used golang's native test framework for both unit tests and integration tests. The integration test verifies the API result and the MySQL database entries. 

Since the requirement is to only have 1 API endpoint, the user resource is created automatically if the user doesn't exist yet. Otherwise I would have created a separate controller/logic for the user resource. A few values also hard coded to skip config management. Some errors aren't checked because they are assumed to not fail, this isn't production grade and otherwise would've been checked for completeness.
