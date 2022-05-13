### Prerequisite:
- .Net SDK (.Net 5) 
- Visual Studio 2019-2022

### Project Structure:
- ToDoApp.API: API Controller 
- ToDoApp.DTO: Database setup, entity structure
  - Using In memory DB of Entity Framework
- ToDoApp.Test: Unit and Integration Test

### Run code Locally:
## Option 1: Via .Net CLI
- After install .NET SDK, please verify that .Net SDK has been installed successfully by open Command Line Tool (Windows) or Terminal (MAC) and Run 
```
dotnet --version
```

- Locate **togo\ToDoAPI** and run command
```
dotnet build
```

- Locate **togo\ToDoAPI\ToDoApp.API**  and run command
```
dotnet run -c Debug
```

- The project will be built and using port 5001

I have integrate Swagger into this project that we can view and execute API via UI - Please open localhost:5001

## Option 2: Via Visual Studio
- Open Solution and Start Debug 

### Run Test Locally

## Option 1: Via CLI

- Locate **togo\ToDoAPI** and run command
```
dotnet test
```
## Option 2: Via Visual Studio
- Open Solution and open Test Explorer then Run test


With this project, i love at that it has ability to write unit and integration test

If there is more time for me, i would like to work about the authentication when execute API and API logging.