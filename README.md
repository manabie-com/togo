## TOGO: Todo using Go
Demonstrate Project Architecture for Golang API Services -07022022

## Requirements
Primary Tech stack
* [GoLang](https://go.dev/)
* [MySQL](https://www.mysql.com/)
* [Go SQL Driver](https://github.com/go-sql-driver/mysql)
* [GORM](https://gorm.io/index.html)
* [Gorilla Mux](https://github.com/gorilla/mux)

## Project structure
```bash
├── cmd
│   ├── app
│   │   ├── config  # Configuration files
│   │   │   ├── environment
│   │   │   │   |── environment.go
│   │   │   │   |── environment_test.go
│   │   │   ├── database
│   │   │   │   |── database.go
│   │   │   │   |── database_test.go
│   │   │   ├── deployment
│   │   │   │   |── deployment.go
│   │   │   │   |── deployment_test.go
│   │   │   ├── routes
│   │   │   │   |── routes.go
│   │   │   │   |── routes_test.go
│   │   │   ├── other configs like metrics, monitoring..
│   ├── main.go     # Main entry point
├── internal        # Application internal/services files
│   ├── utils       # Application Common utils
│   │   │── response
│   │   │   |── response.go
│   │   │   |── response_test.go
│   │   │── errors
│   │   │   |── errors.go
│   │   │   |── errors_test.go
│   │   ├── repository
│   │   │   |── repository.go
│   │   │   |── repository_test.go
│   ├── todo # Todo Service
│   │   │   |── constants.go  # Optional: create constants in separate file if there's too many
│   │   │   |── models.go
│   │   │   |── service.go
│   │   │   |── service_test.go
│   ├── user # User Service
│   │   │   |── constants.go  # Optional: create constants in separate file if there's too many
│   │   │   |── models.go
│   │   │   |── service.go
│   │   │   |── service_test.go
│   ├── other internal application services
├── test
│   ├── integration
├── .env
├── .env.local
├── README.md
└── .gitignore
└── go.mod
└── go.sum
├── ...
```

## Installation

* Clone this repo

```bash
~$ git clone https://github.com/xrexonx/togo.git
```

* Change Directory

```bash
~$ cd todo
```

* Create `.env` file

```bash
~$ touch .env
```

```bash
~$ cp .env.local .env
```

* Modify `.env` file with your correct database credentials and desired Port

## Usage

To run this application, execute:

```bash
~$ go run cmd/app/main.go
```

You should be able to access this application at `http://127.0.0.1:{portInYourEnvFile}`

>**NOTE:**<br>
>When you run/serve the application, there is a migration script configured already to create the DB and tables,
>as well as seeds sample users for testing purposes, see database.go line 76 triggered from line 51 on the same file

## Sample request
I've configured a sample health check endpoint
```bash
$ curl http://127.0.0.1:{yourPort}/api/v1/healthCheck -v
```
Response
```bash
{"Status":"Health Checked","Message":"API is running","Code":200,"Data":null}
```