# togo

- this is service for node writing by typescript

## Version specification

- node: v12.22.0
- TypeScript v4.5.5
- MongoDB: latest
- Kafka: latest
- Redis: latest

## IDE recommendations and setup

- VSCode IDE
- Eslint vscode plugin
- Prettier vscode plugin for linting warnings (auto fix on save)
- Add the following setting in vscode settings.json

```json
"eslint.autoFixOnSave": true
```

## Dev setup

- Install all the dependencies using `npm install`.
- If your env does not have services (MongoDB, Kafka, Redis), you can run `make mongo-up`, `make kafka-up`, `make redis-up` to start service or `make services-up` for starting all services.
- To run the server with watch use `npm run start:dev`.

## Test

- Unit Test: We are using Jest for assertion and mocking.
- To run the test cases use `npm run test`.
- To get the test coverage use `npm run test:cov`.
- To run the integration test, you should setup dev (see ## Dev setup section), then run `npm run start:dev` for starting the server first. Then you should run `npm run test:integration`.

## Git Hooks

The seed uses `husky` to enable commit hook.

### Pre commit

Whenever there is a commit, there will be check on lint, on failure commit fails.

## ENV variables

- create .env file for all config.

```none
#SERVER CONFIG
SERVICE_NAME=togo
HOST=localhost
PORT=3000
LOG_LEVEL=info
MONGO_URI=mongodb://localhost:27018/togo-db
MONGO_DB_NAME=togo-db
MONGO_USER=local-togo
MONGO_PASS=root
KAFKA_URL=localhost:39092
KAFKA_GROUP_ID=kafka-group
REDIS_URL=redis://localhost:6379/

# ZOOKEEPER CONFIG
ZOOKEEPER_VERSION=latest
ZOOKEEPER_PORT=2182

# KAFKA CONFIG
KAFKA_VERSION=latest
KAFKA_PORT=39092
KAFKA_TOPIC=task-consumer

#MONGO CONFIG
MONGO_ROOT_USER=root
MONGO_ROOT_PASS=root
MONGO_PORT=27018
MONGO_INITDB_DATABASE=togo-db

#REDIS CONFIG
REDIS_PORT=6379

```

## Misc

- Swagger API is at <http://localhost:3000/documentation>

### Register User

- This is an API for register new user

- API will return `userId` for the client.

```none
POST /user
BODY
{
    "username": "username"
    "password": "password"
}
Response: 201
{
    "data": {
      "userId": "_userId"
    }
}
```

### Create Task

- This is an API for creating new task.

- The task created will have `PENDING` status, it need to be processed to mark it done for creating.

- When it done, it will produce message to the kafka to process continue step.

```none
POST /user/{userId}/task
BODY
{
    "name": "_name"
}
Response: 201
```

### Get Tasks

- This is an API for getting tasks

- We can add filter `?status=PENDING/FAILED/DONE/ALL`. Default the status is `DONE`.

```none
GET /user/{userId}/tasks
BODY
Response: 200
{
    "data": [
      {
        "id": "string",
        "userId": "string",
        "name": "string",
        "status": "PENDING/DONE/FAILED",
        "reason": {
          "errorCode": "string",
          "message": "string"
        }
  }
    ]
}
```

## Overview

- We use caching database to count the limit of tasks per user that can be added per day.

- We use message queue (kafka) to avoid multi request create tasks at the same time for one user.

- We can improve the performance by separating the service creating task and consuming message for checking possible to be added per day.

- We can improve the performance by scaling out the service with increase the partition of topics and the instances server. Then we can consume more messages to run.

## Need to improve

- Should add more unit test to coverage all codes.

- Write validation for consuming message

- Should write Dockerfile for building docker image for the application

- Find another solution for writing integration test
