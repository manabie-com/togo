## Requirements had accquired
- [x] Implement one single API which accepts a todo task and records it
  - There is a maximum **limit of N tasks per user** that can be added **per day**.
  - Different users can have **different** maximum daily limit.
- [x] Write integration (functional) and unit tests
  - Implemented integration test and unit test by making command: `make integration-test` or `make unit-test`
  - Implemented `make test-all` command which will run both integration and unit test (remove container mssql-test when finished)
- [x] Choose a suitable architecture to make your code simple, organizable, and maintainable
  - Project is using a reference of 3-tier model (just Models package as Data Access Layer and Handlers package as Controller Layer)
- [x] How to run app locally
  - [A brief description about how to run App and its deployment diagram](https://github.com/huynhnhattu/huynhtu/#how-to-run-app)
## How to run App
#### Notes when running App:
  - Has Docker and Go installed
  - **Run command** `docker network create mana-nw` to create a new docker network
  - Project uses `mcr.microsoft.com/mssql/server:2019-latest` as Database container, which not support for MAC M1
### How to run App locally
- Local deployment Diagram

![Local deployment](https://github.com/huynhnhattu/huynhtu/blob/master/local_deployment.png)
- At the workspace directory, enter command `make deploy`
When `make deploy` done, it will build 2 images `manabie-test:latest` and `manabie-mssql:latest`
### Sample `curl` command to call my API
- A sample `curl` command to call API with PUT method and endpoint `api/tasks`
- Note that `"userId"` must not be empty.
- Using `curl` command:
```
curl -X PUT http://localhost:8080/api/tasks -H 'Content-Type: application/json' -d '{"userId":"user_id1","maxDailyLimit":8,"task":"Finish Manabie testing project"}'
```
- Sample request PUT body in JSON format for POSTMAN or ThunderClient vscode's extension
```
{
    "userId":"user_id1",
    "maxDailyLimit":8,
    "task":"Finish Manabie testing project"
}
```
### How to run unit tests locally
There are 2 make commands for testing: **`make test-all`** and **`make test`**

**1. Run all tests on testing container**
- Enter command: `make test-all`
  - When running test-all, a MSSQL container will be created for testing (with port 1434).
  - Container run only one time and will be down after testing done.

**2. Run test after editing code on testing container**
- First setup testing enviroment: `make setup-integration-test`
- Option run test both unit-test and integration-test: `make test`
- Option run only unit-test: `make unit-test`
- Option run only integration-test: `make integration-test`
- Option run only test-coverage: `make test-coverage`
- After testing done run command to shutdown testing container: `docker-compose -f docker-compose-test.yaml down`
### My solutions
- Created a single API which handle PUT request from client. Refer to function handler diagram:
![User's task function handler Diagram](https://github.com/huynhnhattu/huynhtu/blob/master/user_task_handler.png)

### Things that I can't finish because of time limitation
1. Not support for macOS M1 because of using unsupported images MSSQL for macOS m1. I will find a solution for deploying on many platforms.
2. About testing, a problem that when I try to change system day to simulate adding tasks on a new day (tasks limit will be reset after new day), because docker container which I use to run application cannot change system time so I have to wait till next day. But, if I go run at local (not in container), I still could simulate change system time. I will build application from linux image which support `date` command.
3. It was my less experiences in Golang, so I took a lot of time in coding.
