
# How to run
## Installing
- .NET Core 2.0 x64
	https://dotnet.microsoft.com/en-us/download/dotnet/thank-you/sdk-2.0.0-windows-x64-installer
- .NET Core 3.1 x64
	https://dotnet.microsoft.com/en-us/download/dotnet/thank-you/sdk-3.1.100-windows-x64-installer
- SQL Server Engine 2019
	https://go.microsoft.com/fwlink/p/?linkid=866662
- Visual Studio 2019 Professional
	https://visualstudio.microsoft.com/thank-you-downloading-visual-studio/?sku=Professional&rel=16&src=myvs&utm_medium=microsoft&utm_source=my.visualstudio.com&utm_campaign=download&utm_content=vs+professional+2019
## Start API
- Open "MyTodo.sln" file with "Visual Studio 2019"
- Press Ctrl + Shift + B to build & install nuget package
- Right click on "MyTodo.BackendApi" project -> Choose "Set as Startup Project"
- Tools->Nuget Package Manger=> Package Manager Console
	- Select default project "MyTodo.Data.EntityFramework"
	- Run command "Update-Database" to generate database
- Start project (Press F5)
- Open swagger
## Run tests
- Open solution
- Menu Test => Run All Tests
# Todos
- Testing bussiness rules of tasks managment as Task Limit exceeded, Assign permissions,...
- User management, Roles management
- Logging
- Caching
- etc.