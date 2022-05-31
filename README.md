### Evironments
- .NET 6.0
- Database : MongoDB Server

### How to test
- In the root folder (containing the ToDo.sln file), run the following in the terminal:
dotnet test 

### How to run
- In the root folder (containing the ToDo.sln file), run the following in the terminal:
dotnet run --project .\ToDo\ToDo.Api.csproj

### View Swagger Doc
- While running, open a browser and navigate to:
https://localhost:7161/swagger/index.html

### Sample Curl
- Creat User : 
curl -X 'POST' \
  'https://localhost:7161/api/users' \
  -H 'accept: */*' \
  -H 'Content-Type: application/json' \
  -d '{
  "fullName": "Le Van Binh",
  "dailyTaskLimit": 10
}'

- Get User
curl -X 'GET' \
  'https://localhost:7161/api/users?userId=e7e1ff8a-e045-4188-974b-24f750a9bdb8' \
  -H 'accept: */*'
  
- Create task
curl -X 'POST' \
  'https://localhost:7161/api/todos' \
  -H 'accept: */*' \
  -H 'Content-Type: application/json' \
  -d '{
  "userId": "e7e1ff8a-e045-4188-974b-24f750a9bdb8",
  "status": 0,
  "todoName": "implement todo project",
  "todoDescription": "do somethings"
}'

- Get task
curl -X 'GET' \
  'https://localhost:7161/api/todos?ToDoId=c8a73ed4-4098-4850-8a2e-a7a70da6f36b' \
  -H 'accept: */*'

### What I love about my solution
- I choose MongoDb to be easily extensible when the client has requirements to change in the future: for example adding new properties in user or task.
- Support Swagger.
- DI very easy to use, hence testing is easy.

### What else do you want us to know about however you do not have enough time to complete?
- To be honest, I just started writing unit tests and have never written integration tests. I tried to have high coverage for the main API. I have learned about integration tests but haven't had enough time to do them.
