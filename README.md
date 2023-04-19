### How to run the code locally:

* Build the docker image:

```sh
make
```

* Run the service & database using docker compose:

```sh
make run
```

* While developing, run `make restart` to rebuild & reload the service.

### How to test the api:

1. Build and run the service & database:

```sh
make
make run
```

2. Migrate & seed the database:

```sh
make migrate
make seed
```

3. There are two seeded users:
  * Email: `user@example.com`, Password: `gophers`, Maximum daily todo: `1` 
  * Email: `user2@example.com`, Password: `gophers`, Maximum daily todo: `2` 

4. Get authentication token:

```sh
curl -il 'http://localhost:3000/v1/users/auth' \
    -H 'Content-Type: application/json' \
    --data '{ "email": "user@example.com", "password": "gophers" }'
```

5. Create todo:

```sh
export TOKEN="<token from previous step>"

curl -il 'http://localhost:3000/v1/todos' \
    -H 'Content-Type: application/json' \
    -H "Authorization: Bearer ${TOKEN}" \
    --data '{ "title": "test title", "content": "test content" }'
```

### How to run tests locally:

```sh
make test
```

### What do you love about your solution?

* Use some (nice) foundation/infrastucture code from the [service](https://github.com/ardanlabs/service) project.
* Use vertical slice architecture.

### What else do you want us to know about however you do not have enough time to complete?

* Write tests for db queries.
* Add some more unit tests and integration tests, currently only add one each for demonstration purposes.
* Provide some additional APIs like register a new user, list todos, ...
* Provide admin user and authorization, to set up daily maximum todo items for users.
* Add tracing (via OpenTelemetry).
