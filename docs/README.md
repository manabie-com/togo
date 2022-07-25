# TOGO

## Technical Specs
- Backend implemented in Go
- Storage Backend is PostgreSQL
- Use Sliding Window as Rate Limiter Algorithm

## Getting Started

### Prerequisites
- [Golang](https://go.dev/doc/install)
    - Since TOGO is written in Go, so before you get started, you need to install Go. As the dependencies are managed by
      Go Module, the lowest Go version is supported is 1.11, though we recommend using the 1.18.1 for development and
      production.
- Docker Compose
- cmake
### Build

#### Bring up/down the dev environment

`make` and `docker-compose` is your friend. All you need is:

```shell
╰> make docker-compose 
```

#### Give it a try!
##### Create a Task
```shell
curl --request POST \
  --url http://0.0.0.0:9000/tasks \
  --header 'Content-Type: application/json' \
  --data '{
	"user_id": "c7cd294c-627f-452a-8c46-33b5dbfca47f",
	"title": "1st task",
	"note": "should have a better doc"
}'
```

##### Run Test
```shell
make unit-test
make integration-test
```

If you'd love to run a single unit test instead of the whole suite (time, you know)
```shell
go test -run CreateTasks ./ nternal/server/handler
```





