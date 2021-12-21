# TO-GO
A todo API written in golang 

### Features

- Accepts a todo task and records it
  - There is a maximum **limit of 5 tasks per user** that can be added **per day**.
  - Initially uses in memory storage but can be changed

## API

### Endpoints
Access the endpoint specification via Swagger

The docs has been pre generated on the repo but incase you want to update the docs
```
# install swag tool
 GO111MODULE=off go get github.com/swaggo/swag/cmd/swag

# generate swagger docs
~/go/bin/swag init
```
go to [swagger index](http://localhost:8080/swagger/index.html) to access specification and try out the api

## Run the API
```bash
PORT=<PORT HERE> go run main.go
```

## Test the API
```bash
go test ./...
```


