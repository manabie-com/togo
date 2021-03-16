### Overview

Just another To-do API refactored.

### Setting up development environment

- Prequisites: Must have `docker` and `docker-compose` installed in the system
- Run the following commands to have the system running:

  ```sh
  $ make run # Run the system with docker-compose
  $ make migrate # Migrate database
  $ make migrate_test # Migrate database for integration test
  $ make seed # Seed default data
  ```

- To-do API can be accessed at `http://localhost:5050` and PgAdmin can be accessed at `http://localhost:5080`.
- Test can be run using `make test` command. But it requires the system already running with docker.
- Default credentials seeded is `linh / linhdeptrai`. The exported Postman collection contains the latest changes.

### Refactored Project Structure

```sh
- build # Dockerfile
- cmd # Entrypoint to build app
  - todo # Main app the API
  - todomigrate # Migration 
  - todoseed # Seed
- db # Migrations and seeds sql
- deploy # docker-compose file and config needed to deploy the system
- internal
  - pkg # Packages that will be used across apps
  - todo # Main todo app codes
    - domain # Hold domain objects, interfaces
    - handler # HTTP handler, main app entry
    - service # Business logic
    - repository # DB access layer
    - mocks # Mockery generated mocks
    - tests # Integration tests
- scripts # Scripts for easier development
```

### What's done

- Easier local development setup:
  - Only docker and docker-compose needed.
  - Live reload configured.
  - Using Postgres with PgAdmin included.
  - Migrate and seed command included.
- Refactor code to have a clearer structure and better DRY.
- Fix database's structures
- Unit test for `service` and `repository` layer.
- Integration tests with real database.
- Makefile with necessary commands.
- Using Postgresql to handle concurrent inserting tasks based on count condition

### What to improve

- More thorough integration test cases.
- Better error handling and tracing.
- Define request body and response structure for each entity
- Better validation mechanism (validator object?)
- Instead of using Postgresql to handle concurrent inserting tasks based on count condition, use Redis or other key-value DB with some kind of key locking mechanism for better performance.
- Using Redis or other key-value DB for saving and invalidating jwtToken, so we can have Logout feature.
- Implement refresh token feature.
- Better timezone handling (currently can only handle 1 timezone)