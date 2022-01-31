# Togo api

## Infomations
- Framwork: Fiber
- Database: Mysql
- Orm: Gorm

## Config
- create config.json file in root dir
- config.example.json file is an example
```json
{
	"Port":   "<change me: port number>",
	"MySqlUri": "<change me: mysql url>"
}
```

## Run in local
- Create config.json
- command: `go mod download`
- command: `go run main.go`

## Run in docker-compose
- install docker version: 19.03.13
- install docker compose: 1.17.1
- command: `docker-compose up`

## Unit test
- command: `do test ./tests/...`

## Testting using postman file `togo.postman_collection.json`
- Postman version: 2.1
- Change variable "DOMAIN" to your domain like `http://localhost:5000`
- Run request `Create user` to create a user
- Run request `Login` to get token
- Change variable "AUTH_TOKEN" = token which you have got
- Run request `Create one` to create one task
- Run request `Get one` to get one task
