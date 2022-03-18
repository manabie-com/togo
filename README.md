# Manabie

## Introduction

A to-do project from manabie challenge

## Usage

#### Prerequisite

golang and docker must be installed.

#### How to run code locally
- After running program by docker-compose file, automatically run database mirgration files
  - Run docker-compose
  ```
  docker-compose up -d
  ```
- In order to run the app without building the binary file, please run following commands:

  - Run without binary file
  ```
  go run ./cmd/manabie_togo/main.go
  ```  
  - If you want to run the app by building binary file, please run following commands:
  ```
  go build -o main ./cmd/manabie_togo/main.go
  ./main
  ```

#### How to run unit tests locally

- Access to golang directory and run command go test:
  ```
  go test ./...
  ```
## Sample “curl” command to call API

- Create-User CURL command:

    ```
    curl --request POST \
      --url http://localhost:8080/api/user \
      --header 'Content-Type: application/json' \
      --data '{
        "username": "example",
        "password": "123456",
        "limit_task": 3
    }'
    ```

- Create-Task CURL command:

    ```
    curl --request POST \
        --url http://localhost:8080/api/task \
        --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2NDM1NjU2Nzd9.ESVv1A-nMACuVAfl2mKb4MwFaXw-FMn2ZqS_HpjwqYc' \
        --header 'Content-Type: application/json' \
        --data '{
        "content": "Quét nhà"
    }'
    ```

## What i love

- Good, quiet clean, maintainable architecture
- Meet the requirements
- Having useful tools (ex: migration tool, api request logging, ...)
- Nice Reademe file

## Improve

- Improve unit/integration tests
- Add swagger feature
- Add APIs about admin user domain
