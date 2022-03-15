## How to run the code locally

To run the code locally, some pre-requisites need to be installed:

* `.NET 6` - https://dotnet.microsoft.com/en-us/download/dotnet/6.0
* Microsoft SQL Server - https://www.microsoft.com/en-us/sql-server/sql-server-downloads

After installing these pre-requisites, make sure both `dotnet` and `sqlcmd` is included in your environment variables or `PATH`.

Finally, running the SQL scripts in the sub-directories of the `db/` directory will create the required stored procedures and tables.

## Sample curl commands to access the API

### POST User

```
curl --request POST \
  --url https://localhost:5001/api/users \
  --header 'Content-Type: application/json' \
  --data '{
    "FirstName": "John",
    "LastName": "Smith",
    "DailyTaskLimit": "5"
}'
```

### GET User

Using the `id` returned when creating a user, supply the `id` as query to get User details.

```
curl --request GET \
  --url 'https://localhost:5001/api/users?id=1'
```

### POST Todo

```
curl --request POST \
  --url https://localhost:5001/api/todos \
  --header 'Content-Type: application/json' \
  --data '{
    "UserId": "1",
    "Name": "Todo Sample Name",
    "Description": "Todo Sample Description"
}'
```

## Running the program

### API

To run the solution, simply run the command `dotnet run` on the `src/Todo/Todo.Host` directory. Or simply `dotnet run --project src/Todo/Todo.Host`.

### Tests

To run the tests, simply run `dotnet test`. This command can be run on either the root directory of the solution, or in the test-specific folders, namely:

* `src/Todo/Todo.Tests.Unit`
* `src/Todo/Todo.Tests.Integration`

## What I love about my solution

The solution incorporates a lot of the important concepts I've learned over the past few years, namely:

* Domain Driven Design (DDD) and Microservice Architecture, which allows for separation of application, model, and infrastracture layers. Additionally, by separating the request objects (`Todo.Contract`) from the actual model objects (`Todo.Domain.Models`) this allows the user to create requests with minimal information, while a lot of the other data that the user doesn't have to know is abstracted (such as the auto-generated `id` retrieved from the database). This design also makes it robust to domain and infrastructure changes.
* In line with this, I also opted for using Screaming Architecture when naming directories and projects, which makes it simple to easily tell at a glance what each directory/project is supposed to do.
* Repository pattern which allows very easy switching of database implementations through the abstraction and centralization of the data layer and domain classes. For example, if the user would want to use a different database other than SQL, the user simply needs to implement the `IRepository` interfaces and follow the `Storage.Contract` classes.
* Using convenient tools such as `Xunit` and `FluentAssertions` for the tests, which make testing easy and readable, even for more multiple parameter tests through functionalities like `[InlineData]`.
* Dependency injection, which makes the classes free of dependencies. All dependencies are injected meaning changing implementations of back-end logic simply require changing a single line in the host, for example, registering the UserManagementService would be as simple as:
  
  ```csharp
  builder.Services.AddTransient<IUserManagementService, UserManagementService>();
  ```

## Others

Regarding other things that I would want the reviewers to know about, but did not have time to complete, would be the following:

* Applying secrets for sensitive data, such as connection strings for better security. To keep it simple, I opted to simply read from the `appsettings.json` file supplied by .NET Core by default.

* Better disposal of created database entries in the integration tests, namely the created user and todo items.

* Better logging. While I did use the default logging framework, if time permitted I would have added more user-friendly error messages and better logging framework alternatives that I'm familiar with such as Serilog and NLog. 

Having said that, I still learned a lot since the .NET version I'm using (NET6) is relatively new and I have no experience with some of the changes. I encountered lots of interesting changes, especially when I was coding the integration tests where the framework and factories have changed significantly from what I'm used to from previous versions.

Lastly, I would have loved to code this project in Go. However, given the time constraints, I felt that I would not be able to give a satisfactory output so I stuck with C# as it is the language I have the most experience with. 


