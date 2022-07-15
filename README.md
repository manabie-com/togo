
# How to run
## OS
 - Window 8, 10 (x64)
## Download & Install inv
- .NET Core sdk 2.0 (x64)
	https://dotnet.microsoft.com/en-us/download/dotnet/thank-you/sdk-2.0.0-windows-x64-installer
- .NET Core sdk 3.1 (x64)
	https://dotnet.microsoft.com/en-us/download/dotnet/thank-you/sdk-3.1.100-windows-x64-installer
- SQL Server Engine 2019
	https://go.microsoft.com/fwlink/p/?linkid=866662
- Visual Studio 2019 Professional
	https://visualstudio.microsoft.com/thank-you-downloading-visual-studio/?sku=Professional&rel=16&src=myvs&utm_medium=microsoft&utm_source=my.visualstudio.com&utm_campaign=download&utm_content=vs+professional+2019
## Start API
- Open "MyTodo.sln" file with Visual Studio 2019
- Press Ctrl + Shift + B combination to build solution & install nuget package
- Right click on "MyTodo.BackendApi" project -> Choose "Set as Startup Project"
![image](https://user-images.githubusercontent.com/20410120/179136454-71242327-5b1a-4f43-bfc9-751647bf9fdc.png)

- Visual Studio -> Tools menu -> Nuget Package Manger => Package Manager Console
![image](https://user-images.githubusercontent.com/20410120/179136580-03cc311f-fe89-4a98-b8ee-5abb9b34a9cd.png)

	- In default project combobox -> set "MyTodo.Data.EntityFramework" as default
	![image](https://user-images.githubusercontent.com/20410120/179135272-74431277-e7f0-49f2-aa88-9a0549c63100.png)

	- In Package Manager Console -> run "Update-Database" command to generate database
	![image](https://user-images.githubusercontent.com/20410120/179135396-c18848fb-5820-48e6-b806-a02bc946799a.png)

- Start project (Press F5)
![image](https://user-images.githubusercontent.com/20410120/179135505-6145358f-d743-41c9-83d1-9d2161310736.png)

- OpenAPI browser
![image](https://user-images.githubusercontent.com/20410120/179135708-120bcae0-b2af-4290-8f0d-559df82714fd.png)
## Test API
- Login
	user: admin/pwd: 123456
	![image](https://user-images.githubusercontent.com/20410120/179136055-32591850-03e1-45d6-ad83-ada62a0dbe47.png)
	![image](https://user-images.githubusercontent.com/20410120/179136247-cb88b61d-4e3b-4cc2-ab41-5c3103eb6388.png)
	
## Run tests
- Solution level: Test menu => Run All Tests
- Project level: Right on project name -> Run Tests
- Function level: Open function impl code -> Run Test(s)
# Todos
- Testing bussiness rules of tasks managment as Task Limit exceeded, Assign permissions,...
- User management, Roles management
- Logging
- Caching
- etc.
