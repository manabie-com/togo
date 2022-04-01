### Clone
- Clone project

>  https://github.com/leducphucdev/togo.git

> cd togo

### Run Docker
- Copy .env.example to .env
- Download docker and run command

> docker-compose up

- Use curl call api create todo

> curl -X 'POST' 'http://localhost:3000/todo' -H 'Content-Type: application/json' -d '{ "task": "test" }'; 

- Attach container for run test

> docker exec -it manabie_app bash 

- Run command test

>  yarn test:e2e

### Not run Docker
- Copy .env.example to .env
- If you're not running with Docker then install the package postgres with version 13 and node 12
- Then create database postgres:

  - user: admin
  - password: 123456
  - databse: manabie
  - port: 5432

- Then write .env. Example:

  DB_DIALECT=postgres

  DB_USER=admin

  DB_PASSWORD=123456

  DB_NAME=manabie

  DB_HOST=127.0.0.1

  DB_PORT=5432

- Then run command

> yarn install

> yarn start:dev

- Use curl call api create todo

> curl -X 'POST' 'http://localhost:3000/todo' -H 'Content-Type: application/json' -d '{ "task": "test" }'; 

- Run command test

> yarn test:e2e

### Structure

- Use DDD(Domain Driven Design) to can open for microservice
- Use layer service Application Layer and Domain Layer in DDD for Business
- Use Decotor pattern in Nest for Aspect Oriented Programming. Used to separate the processing stream into separate objects from each other. As validation, cache, log ...