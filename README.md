# simple-shopping
This is just a simple demo for demonstration concepts of microservices.

## Setup development environment
- Install docker. Please refer to https://docs.docker.com/engine/install/
- Install docker compose. https://docs.docker.com/compose/install/

## Technical solution
- Containerization with docker
- API and Microservices Framework: Looback Framework. Please refer to https://loopback.io/
- Database with MongoDB
- Messaging with Rabbitmq

## Technical Design
- Entity Relationship Diagram:

  ![ERD](/_docs/assets/design-erd.jpg)

- Flow of getting todo detail:

  ![Todo Detail](/_docs/assets/design-get-todo.jpg)

- Flow of logging http request:

  ![Access Log](/_docs/assets/design-access-log.jpg)

## Test
- open terminal and go to project folder
- run: `bash _devtools/run-n-test.sh`

* To clean up after test
- run `bash _devtools/clean-up-test-env.sh`

## Run locally
- open terminal and go to project folder
- run: `bash _devtools/start-dev.sh`

### Check api using curl
  - first step: create user by executing `bash _devtools/curl/login.sh`
  - create todo: `bash _devtools/curl/create-todo.sh`
  - get todo: `bash _devtools/curl/get-todo.sh`
  - list todo: `bash _devtools/curl/list-todo.sh`
  - create multiple todo: `bash _devtools/curl/create-todo-without-id.sh`

  Notes: all scripts are using auto login. If you want to login manually please call login script separately by invoking `bash _devtools/curl/login.sh` and then copy token and token to ACCESS_TOKEN in other shellscript

## Open API Explorer
### Product API
- open browser
- enter the url: http://localhost:3000

## Clean up dev
- run `bash _devtools/clean-up-dev.sh`