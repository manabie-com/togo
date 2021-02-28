# tokoin

## Introduction

A code challenge from manabie.

## Usage

#### Prerequisite

golang must be installed.

#### How to run the app

- In order to run the app without building the binary file, please run following commands:

    - Run with default config
  ```
  go run main.go
  ```
    - Run with config file
  ```
  go run main.go -config-file ./config/togo-server.yaml
  ```  

- If you want to run the app by building binary file, please run following commands:

```
go build -o challenge
./challenge
```

### Get sample config
```
go run main.go -example
```

#### How to test the app

- Access to golang directory and run command go test:

```
go test ./...
```

## Source code structure explanation

```
|-config/
|-db/
|-docs/
|-internal/
|-services/
|-storages/
    |-task/
    |-user/
|-pkg/
    |-config/
    |-cmsql/
    |-crytpo/
    |-xerros/
|-up/
|-main.go 
```

The app is splitted into serveral packages:

- **config**: It contains ```config.go``` to store config struct and default config. Without any external config file, the program still be runnable.

- **db**: It contains migration files for database.

- **docs**: It contains Postman sample requests collection.

- **internal**: It contains ```services, storages``` packages:
    - services: Provides functions to handle requests and return responses.
    - storages: Contains models and handle jobs with database.

- **pkg**: It contains ```config, dot``` packages:
    - config: Provides functions to load config from file or default.
    - cmsql: Provides functions for mock data tests and config Postgres.
    - crypto: Hash and Check password.
    - xerrors: Define errors and map its with httpstatus.

- **up**: Define structure responses and service interfaces.

- **main.go**: It load configs and run server.

## Future

- Create tool to auto-generate SQL simple functions (CRUD)
- Split middleware and add Claim for all requests.
- Use redis to handle case MaxTodo.
- Create tool for log request, sql.
- Add dependencies injection.
- Add apis to update MaxTodo.