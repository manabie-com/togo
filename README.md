TOGO
==========

A simple home assessment for apply position Backend Developer at Manabie

Table of Contents
-----------------

  * [Requirements](#requirements)
  * [Technical](#technical)
  * [Usage](#usage)
  * [Comment](#comment)

Requirements
------------

Togo requires the following to run:

  * [Golang](https://go.dev/doc/install) 1.18+
  * [Docker](https://docs.docker.com/get-docker/)
  * [Docker Compose](https://docs.docker.com/compose/install/)

Technical
------------
- Backend implemented in Go
- Database i used PostgreSQL, cause Manabie use PostgreSQL right?

Usage
-----
Run code locally
- First install local database
```sh
make docker-compose
```
- Then create table
```sh
make migrate
```
- Run server
```sh
make run
```

A sample “curl” command to call your API
- First we need to create user, limit_count is **limit of N tasks per user** as requirement
```
curl --location --request POST 'http://localhost:8081/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "name",
    "limit_count": 1
}'

Sample Response:
{
    "data": {
        "id": 1,
        "limit_count": 1,
        "name": "name"
    },
    "status": 200
}
```

- Then main api is create task, we get user_id from previous request we make, then put in url /users/:user_id/tasks 
```
curl --location --request POST 'http://localhost:8081/users/1/tasks' \
--header 'Content-Type: application/json' \
--data-raw '{
    "description": "12",
    "ended_at": "2022-07-27T11:55:37+07:00"
}'

Sample Response:
{
    "data": {
        "id": 1,
        "user_id": 1,
        "description": "12",
        "ended_at": "2022-07-27T11:55:37+07:00"
    },
    "status": 200
}

If user get limit task, response:
{
    "data": "limit_max",
    "status": 400
}
```

- I have another api to check how many task user have
```sh
curl --location --request GET 'http://localhost:8081/users/1/tasks'

Response:
{
    "data": [
        {
            "id": 1,
            "user_id": 1,
            "description": "12",
            "ended_at": "2022-07-27T11:55:37+07:00"
        }
    ],
    "status": 200
}
```

Run unit test
```sh
make unit-test 
```

Run integration-test, remember create local database before
```sh
make docker-compose # if did not create local database
make integration-test
```

Comment
------------
Why i love my solution?
- Well i keep everything simple
    - First i think about create 3 table users, tasks and user_setting but our system very simple, so i add limit_count to table users, if our system go scale, we can migrate database later
    - Only create file if needed, no Dockerfile because i dont deploy this project, only run command
- I choose echo library , for fast implement, minimalist
- I code follow Domain Driven Design, no Layered Architecture, i have task and user domain, even it very simple but easy to separate, move this domain to another project, i means another service in micro-service
- Manage migration with ```gormigrate```, which is easy to migration database with no sql script

If i have enough time to complete:
- Caching each api get, yes absolutely
- API document with swagger
- API login to get access_token, so that we can authentication user
- Optimized the database, use the aggregated column for checking exceeded daily limit instead of `count` query. Need to pay attention if there are many records
- Manage error with error code, not return errors.New() like i did in this project