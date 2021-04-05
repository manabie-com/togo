# Manabie-togo

## How to run?
1. To run the program: **docker-compose --env-file ./configs/app.dev.env up dev**
2. To run integration_test: **docker-compose --env-file ./configs/app.test.env up integration_test**
3. To run unit test of postgres in storages: **docker-compose --env-file ./configs/app.test.env up storages_test**
4. To run unit test of usecase: **docker-compose --env-file ./configs/app.test.env up usecase_test**

## Explain what I do.
### Structure:

```
├── configs // variable environment with stage
├── docker-compose.yaml // run program with docker
├── scripts // scripts for init db
├── internal
│   ├── logs // store logs into file for tracing.
│   ├── usecase // logic between transport and storages
│   ├── util // contain file config, function common and random
│   ├── transport
│   │   ├── integration_test.go // run integration_test
│   │   ├── rate_limit.go // logic use for rate limit this application
│   │   ├── server.go // serving, forward request and return response
│   │   └── todo.go // validate, logic and call todoUsecase
│   ├── storages
│   │   ├── entities // folder define entity
│   │   ├── postgres // logic interact with database of postgres, unit test
│   │   ├── sqlite // logic interact with database of sqllite
│   │   └── store.go // define interface for database
├── main.go
```
        
### What did I do on this project?
1. Use sha256 for hash password
2. Use index for query.
3. Use offset, limit for get list
4. Transaction when create task
5. Strategy pattern
6. Create file for variable environment
7. Logging into file for tracing
8. Use Gin framework for serving
9. Change login become method POST
10. Graceful shutdown and ratelimit
11. unit test and integration test


### What is I want to improve?
* I want to use Kong gateway for ratelimit.
* Use another framework to interact with postgres, such as: sqlx, gorm.
