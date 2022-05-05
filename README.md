## How to run the project:
```
make run
```

## Routes:
#### Register
```
curl -X POST 'http://localhost:8080/register' -d 'username=sample_user' -d 'password=sample_password' -d 'rate_limit=0'
```
Note : if rate limit was set to 0 then the user does not apply any limit per day
#### Login
```
curl -X POST 'http://localhost:8080/login' -d 'username=sample_user' -d 'password=sample_password'
```
Note : Login to get the token need for each request
#### Create task
```
curl -X POST -H "Authorization: Bearer <place your token here>" 'http://localhost:8080/create_todo' -d 'name=sample task'
```

## how to run test locally
To run my test locally you can just call on `make test`  or `go test -v ./...`

## Notes :
I used sqlite as my database which needs cgo enaled.
