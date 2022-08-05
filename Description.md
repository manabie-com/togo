# Assignment Description

## How to run code locally?

### Prerequisite

- Docker Desktop

### Run the web API

The default port for the API is `8080`. Make sure that the port `8080` is available on your local environment. If not, please modify the `togo.api` service in `services` section in `docker-compose.yml` file to expose another port.

```
docker compose -f docker-compose.yml up --build
```

The database migration will be run automatically and an admin user will be seeded. The default credentials for the admin user is 

```
admin / Abcd@1234
```

#### Postman collection

There is a Postman collection which resides in the repository root directory named `Togo.Api.postman_collection.json`.

#### cURL

Login as admin

```
curl --location --request POST 'http://localhost:8080/api/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
  "UserName": "admin",
  "Password": "Abcd@1234"
}'
```

Create a user

```
curl --location --request POST 'http://localhost:8080/api/user' \
--header 'Authorization: Bearer ADMIN_BEARER_TOKEN' \
--header 'Content-Type: application/json' \
--data-raw '{
  "UserName": "user1",
  "Password": "Abcd@1234",
  "MaxTasksPerDay": 10
}'
```

Create a task

```
curl --location --request POST 'http://localhost:8080/api/tasks' \
--header 'Authorization: Bearer USER_BEARER_TOKEN' \
--header 'Content-Type: application/json' \
--data-raw '{
  "Title": "Discuss with Kelven about the next project"
}'
```

### Run the integration tests

```
docker compose -f docker-compose-integration-tests.yml up --build --exit-code-from togo.api.integration_tests
```

## What do you love about your solution?

I personally love building flexible, reusable project template. Everytime I build a new project structure, I learn new things. This time, what I have learned are:

- Separating the infrastructure (technology related) from the core business logic. The only issue is that I did not fully separate the data access from the core business project, I still use the `DbSet` class which belongs to EntityFramework Core. This could be solved by implementing our own repository and unit of work. 
- Dockerising the integration tests solution

## Why I do not have enough of time to finish?

I spent a lot of time in the first couple days to find a right way to setup the integration tests, custom host creation, custom `Startup` class, custom dependency injection,... After getting stuck at that stuff for a long time, I decided to do the Docker stuff (writing Dockerfile, docker compose for all executable projects) first and go back to the integration tests when I have enough time.

It turned out that when I properly dockerised the integration tests, my problem has been resolved as well. At the first stage, I should think about the examiners situation that they might not have .NET environment set up on their local machine and everything should be run on Docker containers, so I should not really need to care about isolating the integration tests environment. Thanks God.
