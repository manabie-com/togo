  # Togo Service
There are 3 projects as below:
  - **TogoService.API**: It's main application for WEB API. It's a lightweight web API. Business logic is in Controller class as well.
  - **TogoService.UnitTest**: Unit test for API. I only care about the method AddTasksForUser in UserController class. It contains all business logic.
  - **TogoService.IntegrationTest**: Integration Test for API.

**Language**: C#.
**Framework**: .NET Core 3.1

## Prerequisite
- .NET Core 3.1 SDK ([download](https://dotnet.microsoft.com/en-us/download/dotnet/3.1 "download"))
- Editor: any editor you familiar with. 
- [ SQLite Studio](https://sqlitestudio.pl/ " SQLite Studio") for viewing data 

## Run locally
**PLEASE UNINSTALL ALL .NET SDK IN YOUR LOCAL MACHINE IF YOU ALREADY INSTALLED ANY .NET SDK VERSION > 3.1**
**INSTAL CORRECT VERSION OF .NET SDK. .NET Core SDK 3.1 WITH THE LINK ABOVE.**
**YOU CAN INSTALL ANY .NET SDK VERSIONS HIGHER THAN 3.1 AFTER YOU INSTALL .NET SDK VERSION 3.1**
**DO NOT install any .NET SDK versions higher than 3.1 before install version 3.1.**
- Clone code base.
- Navigate to root code base folder. Run `dotnet build`. It will restore and build solution (3 projects).
- Navigate to TogoService.API folder. Run `dotnet run` to run it. Open browser at https://localhost:5001/swagger for API document and try it.
  - You can also use postman to try it.
  - Using CURL: 
> curl -X 'POST' \
  'https://localhost:5001/api/users/{userId}/tasks' \
  -H 'accept: text/plain' \
  -H 'Content-Type: application/json' \
  -d '{
  "Date": "2022-07-05T03:54:32.843Z",
  "Tasks": [
    {
      "Name": "task 1",
      "Description": "task thứ nhất"
    }
  ]
}'

## Run locally with Docker
- Clone code base
- Navigate to root code base folder. Run `docker build -t togoservice .` for building image.
- Run `docker run -d -p 5001:80 --name togoapi togoservice` to start container.
- Open browser at http://localhost:5001/swagger for API document and try it or you can use CURL with above information.

## Note: 2 users id for run locally testing:
- id =BA4688FB-B389-40A7-A94C-FE43F61A2BCE - max daily tasks = 10
- id =437A4196-7ADF-441F-8352-05F126C469B2 - max daily tasks = 0

## Unit Test - Run locally
- Clone code base.
- Navigate to root code base folder. Run `dotnet build`. It will restore and build solution (3 projects).
- Navigate to TogoService.UnitTest. Run `dotnet test`.
You can do the same with integration test.

## My solution
- It is lightweight web API project.
- I already dockerized but did not test it.
- Unit test coverage is 100% and integration test coverage is 80%.
![test coverage_unit test](https://user-images.githubusercontent.com/4899325/177812275-8d8f2b01-78e7-4fa5-9450-cfbc2f120129.PNG)

![test coverage_integration test](https://user-images.githubusercontent.com/4899325/177812394-e5b6d475-aa96-4563-952a-2a452f3b351a.PNG)


## More
If I have more time, I think there are some things we can improve:
- Dockerization with careful testing.
- Using environment variables for various environments (dev, qa, staging, production...).
- Authentication and authorization for endpoints.
- I am using SQLite without any security (credentials...). Need to pay attention here.
