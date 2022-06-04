# Togo

Golang application which accepts a todo task and records it if the user has not yet reached the limited number of tasks per day.

## Usage

### I. Run server

Make sure you have Docker installed, otherwise you may download it from this [link](https://www.docker.com/products/docker-desktop/).

If you already have Docker, simply follow the steps to deploy the code.

1. Clone the repository

```Shell
$ git clone git@github.com:jrpespinas/togo.git
```

2. Change directory

```Shell
$ cd togo
```

3. Run docker compose to deploy the application

```Shell
$ docker compose up
```

By now you must be able to see Docker running the containers after building the images.

NOTE: For the sake of this exam, I included the `.env` file in my commits for you to observe successful results when testing and making a simple post request to the application. However, it is bad practice to commit the environment variables to the code repository.

### II. Sample Request

1. You need to register a new user to create tasks.

```Shell
$ curl -X POST http://localhost:8080/registration -H 'Content-Type: application/json' -d '{"email":"admin@gmail.com","password":"password"}'
```

2. Login to start your session

```Shell
$ curl -X POST http://localhost:8080/login -H 'Content-Type: application/json' -d '{"email":"admin@gmail.com","password":"password"}'
```

Finally, make a simple post request to create a task.

```Shell
$ curl -X POST http://localhost:8080/tasks -H 'Content-Type: application/json' -d '{"title":"sample title","description":"sample description"}'
```

You should have received a response such as this:

```json
{
  "status": "Success",
  "code": 200,
  "message": {
    "id": "cabhte81hrh6mgum9d7g",
    "title": "sample title",
    "description": "sample description",
    "created_at": "2022-06-01T08:09:29.0564148Z"
  }
}
```

## Solution

### Deployment: Docker

This approach deploys the Golang Backend via **Docker**. Subsequently, I am able to utilize `docker compose` to run a MongoDB container which resolves compatibility issues.

### Design Pattern: Repository Pattern

This approach uses the **Repository Pattern**.

> Repositories are classes or components that encapsulate the logic required to access data sources. They centralize common data access functionality, providing better maintainability and decoupling the infrastructure or technology used to access databases from the domain model layer.

\- [Microsoft Documentation](https://docs.microsoft.com/en-us/dotnet/architecture/microservices/microservice-ddd-cqrs-patterns/infrastructure-persistence-layer-design)

This pattern is usually made up of three main layers: **Repository**, **Service**, and **Controller**.

The **Controller** layer is responsible for handling the request and for returning the response.

The **Service** layer is responsible for the business logic. This is the layer in which you manipulate the data. In this case, this is where the id and the JWT token is generated. This is also where the validation of tasks and users occur.

The **Repository** layer is responsible for the data access or the interaction with the database.

The main benefit of this approach is to divide the application by layers to encourage long-term maintainability of the codebase. By abstracting each layer from the other, you will be able to easily test and/or to easily refactor the code. For example, since the Repository layer is abstracted from the Service layer, the Service layer does not have to worry which database I use--in this case, MongoDB.

## Things to improve

- Follow file structure best practices
- Learn about proper logging
- implement CRUD

#### Total Working Time: 22 hours
