## TOGO: Todo using Go
Demonstrate Project Architecture for Golang API Services -07022022

## Requirements
Primary Tech stack
* [GoLang](https://go.dev/)
* [MySQL](https://www.mysql.com/)
* [Go SQL Driver](https://github.com/go-sql-driver/mysql)
* [GORM](https://gorm.io/index.html)
* [Gorilla Mux](https://github.com/gorilla/mux)
* [Ginkgo](https://github.com/onsi/ginkgo)

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
git clone https://github.com/xrexonx/togo.git
```

* Change Directory

```bash
cd todo
```

* Create `.env` file

```bash
touch .env
```

```bash
cp .env.local .env
```

* Modify `.env` file with your correct database credentials and desired Port

## Usage

To run this application, execute:

```bash
go run cmd/app/main.go
```

You should be able to access this application at `http://127.0.0.1:{portInYourEnvFile}`

>**NOTE:**<br>
>When you run/serve the application, there is a migration script configured already to create the DB and tables,
>as well as seeds sample users for testing purposes, see main.go line 22

## Sample request
I've configured a sample health check endpoint
```bash
curl http://127.0.0.1:{yourPort}/api/v1/healthCheck -v
```
Response:
```bash
{"Status":"Health Checked","Message":"API is running","Code":200,"Data":null,"Date":"2022-07-03T11:59:07.901816+08:00"}
```

Using the pre-configured user
```bash
curl -d '{"name":"new todo", "description":"this is todo", "userID": "1"}' -H "Content-Type: application/json" -X POST http://{yourHost}/api/v1/todo
```
Response:
```bash
{"Status":"Ok","Message":"Successfully created todo","Code":200,"Data":{"ID":20,"CreatedAt":"2022-07-03T20:14:41.724+08:00","UpdatedAt":"2022-07-03T20:14:41.724+08:00","DeletedAt":null,"name":"Take a bath","description":"value2","completed":false,"userId":"1"},"Date":"2022-07-03T20:14:41.734186+08:00"}
```

## Testing
I'm using Goland to run test, but you can run test on terminal using this command:
```bash
go test ./... -coverprofile cover.out
```

>**Additional note:**<br>
>This project is using Go's new features [Generics](https://tip.golang.org/doc/go1.18#generics) or [Type parameters](https://go.googlesource.com/proposal/+/master/design/15292/2013-12-type-params.md) and requires go version 1.18 or higher.

Go to github [Development branch](https://github.com/xrexonx/togo/tree/mvp)

##TODO
- [ ] Add more Unit test
- [ ] Add more data validation
- [ ] Make repository an interface and reuse to all models

###Questions:
####What do you love about your solution?
The project structure/architecture, it's simple and maintainable and I'd like to improve it more.
####What else do you want us to know about however you do not have enough time to complete?
I can implement this using Ruby/NodeJS/Java-Springboot
