## TOGO
Todo using Go

## Primary Tech stack
* [GoLang](https://go.dev/)
* [MySQL](https://www.mysql.com/)
* [Go SQL Driver](https://github.com/go-sql-driver/mysql)
* [GORM](https://gorm.io/index.html)
* [Gorilla Mux](https://github.com/gorilla/mux)

## Project structure
```bash
├── cmd
│   ├── app
│   │   ├── config
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
│   ├── main.go
├── internal
│   ├── utils
│   │   │── response
│   │   │   |── response.go
│   │   │   |── response_test.go
│   │   │── errors
│   │   │   |── errors.go
│   │   │   |── errors_test.go
│   │   ├── repository
│   │   │   |── repository.go
│   │   │   |── repository_test.go
│   ├── todo
│   │   │   |── constants.go
│   │   │   |── models.go
│   │   │   |── service.go
│   │   │   |── service_test.go
│   ├── user
│   │   │   |── constants.go
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