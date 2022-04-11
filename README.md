## Requirements had accquired
- [ ] Implement one single API which accepts a todo task and records it
  - There is a maximum **limit of N tasks per user** that can be added **per day**.
  - Different users can have **different** maximum daily limit.
- [ ] Write integration (functional) tests
- [ ] Write unit tests
- [ ] Choose a suitable architecture to make your code simple, organizable, and maintainable
- [ ] How to run app locally

## How to run App
#### Notes when running App:
  - Has docker installed
  - **Run command** `docker network create mana-nw`

### How to run App locally
- At the workspace directory, enter command `make deploy`
When `make deploy` done, it will build 2 images `manabie-test:latest` and `manabie-mssql:latest`
### Sample `curl` command to call my API
- A sample `curl` command to call API with PUT method and endpoint `api/tasks`
```
```
### How to run unit tests locally
There are 2 make commands for testing: `make test-all` and `make test`

**1. Run all tests on testing container**
- Enter command: `make test-all`
  - When running test-all, a MSSQL container will be created for testing (with port 1434).
  - Container run only one time and will be down after testing done.

**2. Run test after editing code on testing container**
- First setup testing enviroment: `make setup-integration-test`
- Option run test both unit-test and integration-test: `make test`
- Option run only unit-test: `make unit-test`
- Option run only integration-test: `make integration-test`

### My solutions

