# Manabie Togo

[![Go|GoDoc](https://godoc.org/github.com/hellofresh/health-go?status.svg)](https://godoc.org/github.com/hellofresh/health-go)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

## Project Structure

```sh
.
├── cmd/
│   └── srv/        
│   │   └── main.go           #  main func 
│   └── migrate                
│   |   └── main.go           #  main func for migrate
├── configs/
│   ├── config.go             # contain config variable
├── internal/ 
│   ├── transport/            # contain transport layer
│   │   └── ...
|   ├── domain/               # contain bussiness layer 
│   │   └── ...
│   └── model/                # contain model and repository/store layer
│   │   └── postgres/         
│   │   |   └── ...
│   │   └── redis/         
│   │   |   └── ...
│   │   └── entities.go       #  contain models and entities
│   │   └── task.go           #  contain interface  of task store
│   │   └── user.go           #  contain interface of user store
├── docs/ 
├── common/ 
│   └── errors/
│   |   └── ... 
│   └── constants/
│   |    └── storages/  
├── test/                     # e2e test 
├── └── ... 
├── utils/ 
├── Dockerfile
├── docker-compose.yml
├── docker-compose.dev.yml
├── .gitignore                # specifies intentionally untracked files to ignore
├── .env
├── docker.env
├──.dockerignore
├── go.mod 
├── go.sum
```

## Installation

Make sure you have Go installed ([download](https://golang.org/dl/)). Version `1.16` or higher is required.

Install make for start the server.

For Linux:

```sh
$sudo apt install make
```

For Macos:

```sh
$brew install make
```

## Start server

First of all, you must copy .env.example to .env:

```sh
$cp .env.example .env
```

Start server with cmd/terminal:

```sh
$make docker-dev     # start docker with dev environment
$make migrate
$make start
```

## Start server

Register CURL command:

```sh
$curl --location --request POST 'localhost:8080/register' --header 'Content-Type:application/json' --data-raw '{"user_id": "1","password":"123456", "max_todo":5}'
```

Login CURL command:

```sh
$curl --location --request POST 'localhost:8080/login' --header 'Content-Type:application/json' --data-raw '{"user_id": "1","password":"123456"}'
```

Get List Task CURL command:

```sh
$curl --location --request GET 'localhost:8080/tasks?created_date=2022-06-24 -H "Accept: application/json" -H "Authorization: {token}'
```


Create Task CURL command:

```sh
$curl --location --request POST 'localhost:8080/tasks' \--header {token}
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": "1",
    "content": "task description"
}'
```

Your app should now be running on [localhost:8080](http://localhost:8080/).

## Unit test

```sh
$chmod +x ./test.sh
$make unit-test
```


## What do you love about your solution?
Project architecture is clear, following the `clean-architecture` standard.
Integrate test to check and fix errors.
Golang's speed is very fast.

## What else do you want us to know about however you do not have enough time to complete?
Some functions I want to improve:
1. set the deadline for the task.
2. Set time for each task.
3. Add status to the task.
4. Test and fix bugs more thoroughly, write more unit tests