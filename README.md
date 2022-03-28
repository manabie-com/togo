## Features
- API auth and handle task for each user.
- Split services layer to domain and transport layer.
- Write unit test for the domain layer.
- Write unit test for the store layer.

## Project Structure

```sh
.
├── cmd/
│   └── main.go                #  main func 
│   │          
│   └── migrate                
│   |   └── main.go           #  main func for migrate
├── configs/
│   ├── config.go             # contain config variable
├── internal/ 
│   ├── transport/            # contain transport layer
│   │   └── ...
|   ├── domain/               # contain bussiness layer 
│   │   └── ...
│   └── storages/             # contain model and repository/store layer
│   │   └── postgres/         
│   │   |   └── ...
│   │   └── entities.go       #  contain models and entities
│   │   └── task.go           #  contain interface  of task store
│   │   └── user.go           #  contain interface of user store
├── common/ 
│   └── errors/
│   |   └── ... 
│   └── constants/
│   |    └── storages/  
├── utils/ 
├── Dockerfile
├── docker-compose.yml
├── .gitignore                # specifies intentionally untracked files to ignore
├── .env
├── docker.env
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

## Start project


Start server with docker:

```sh
$make docker-start
```

Your app should now be running on [localhost:5050](http://localhost:5050/).

## Curl sample 

curl --location --request POST 'localhost:5050/task' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDA0OTUzMjUsInVzZXJfaWQiOiJoYW9wcm8xIn0.PrNVQtJ9in3rGaCCsdbl4e07zwdlZzu57pOzyNx7otw' \
--header 'Content-Type: application/json' \
--data-raw '{
    "content":"noi dung 6"
}'

## Test

```sh
$make test
```

## Some love for my solution 
I think it is the way that I implement a global var for avoiding race condition about number task each user
