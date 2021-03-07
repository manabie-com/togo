### Overview
This is a simple backend for a good old todo service, right now this service can handle login/list/create simple tasks.  
To make it run:
- `go run main.go`
- Import Postman collection from `docs` to check example
 
### What I do
- Make code more DRY
- Separate the layer according to the following structure: Router -> Transport -> Service-> Storage -> Model
- Write unit test for user in service and storage layer
- Setup integration test

### What  miss:I
- Write integration test
- Change DB to PostgreSQL
