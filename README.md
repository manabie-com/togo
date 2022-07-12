
# How to run
#### Installing
- SQL Server 2019
- Visual Studio 2019
## Start API
- Open & Build solution
- Review Database Connection string in appsettings.json file in both "MyTodo.Data.EntityFramework" and "MyTodo.BackendApi"
- Set default Assemply to "MyTodo.Data.EntityFramework" in Package Manager Console
- Excute commands:
	- Update-Database
- Set Start up project to "MyTodo.BackendApi"
- Start project (F5)
- Open Swagger
## Run tests
- Open solution
- Menu Test => Run All Tests
# Todos
- Testing bussiness rules of tasks managment as Task Limit exceeded, Assign permissions,...
- User management, Roles management
- Logging
- Caching
- etc.