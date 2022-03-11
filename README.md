### Requirements

- .NET 5 Runtime: https://dotnet.microsoft.com/en-us/download/dotnet/5.0
  - Download the version for your OS.
- Database : Sql Server: https://www.microsoft.com/en-us/sql-server/sql-server-downloads
  - Download express for quick-setup local development / testing.

### Migrations
- Requires dotnet ef CLI Tools : https://docs.microsoft.com/en-us/ef/core/cli/dotnet
- Install using terminal (requires .NET runtime)
```
dotnet tool install --global dotnet-ef
dotnet tool update --global dotnet-ef
```
- In the root folder (containing the .sln file), run the following in the terminal: 
```
dotnet ef database update -s ./TODO.Api -p ./TODO.Repositories
```
### How to test
- In the root folder (containing the .sln file), run the following in the terminal: 
```
dotnet test 
```

### How to run
- In the root folder (containing the .sln file), run the following in the terminal:
```
dotnet run --project ./TODO.Api
```
### View Swagger Doc
- While running, open a browser and navigate to https://localhost:5001/swagger

### Sample Curl
- Get User/s
```
curl -X GET "https://localhost:5001/api/users" -H  "accept: */*"
```
```
curl -X GET "https://localhost:5001/api/users?userId=1" -H  "accept: */*"
```
- Create Todo
```
curl -X POST "https://localhost:5001/api/todos" -H  "accept: */*" -H  "Content-Type: application/json" -d "{\"userId\":1,\"statusId\":1,\"todoName\":\"Submit Challenge to Manabie\",\"todoDescription\":\"Screening challenge for Manabie!\",\"priority\":0}"
```
### What I love about my solution
- I love my solution because:
  - The repository pattern is great for small scale projects, and can scale well even if the project grows.
  - The API layer is separated from service/repository layer, which means that another API project can reside in the same solution and can be deployed concurrently alongside the     existing one (both consume the same service layer).
  - ASP.NET WebAPI templates make DI very easy to use, hence testing is easy!
  - ASP.NET WebAPI come with OpenAPI (swagger) support, making documentation very easy!
  - EF Core code-first migrations make DB creation and deployment easier and more portable!

### What else do you want us to know about however you do not have enough time to complete?
- I would have loved to create this mini-project in Go, but I did not have time to upskill, so I proceeded with C#.NET, as it is my comfort language.