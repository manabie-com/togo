# How to run (Docker)
1. Clone the repository
2. At the root directory
```
docker-compose build
```
```
docker-compose up
```
3. Wait for docker compose
4. You can launch service as below url
```sh
http://localhost:8687/swagger/index.html
```
5. UAT

- Request token:
	- admin@mytodo.com/123456

# How to run (Local)
## OS
- Window 8, 10 (x64)
## Download & Install
- .NET Core SDK 2.0 (x64)
	https://dotnet.microsoft.com/en-us/download/dotnet/thank-you/sdk-2.0.0-windows-x64-installer
- .NET Core SDK 3.1 (x64)
	https://dotnet.microsoft.com/en-us/download/dotnet/thank-you/sdk-3.1.100-windows-x64-installer
- SQL Server Engine 2019
	https://go.microsoft.com/fwlink/p/?linkid=866662
- Visual Studio 2019 Professional
	https://visualstudio.microsoft.com/thank-you-downloading-visual-studio/?sku=Professional&rel=16&src=myvs&utm_medium=microsoft&utm_source=my.visualstudio.com&utm_campaign=download&utm_content=vs+professional+2019
## Start API
- Open Solution
![image](https://user-images.githubusercontent.com/20410120/179144735-1d3d5614-b0b0-49ea-86a6-c6f9beb0c46a.png)

- Build solution & restore package
	![image](https://user-images.githubusercontent.com/20410120/179146449-2b3c52ec-f1fe-4483-8c8b-34253f5b3f0f.png)
 
- Set Start up project
	![image](https://user-images.githubusercontent.com/20410120/179145052-ebf7e339-65c0-4907-b3ae-a180681c2b77.png)
- Migration database
	- Tools-> Nuget Package Manger => Package Manager Console
	![image](https://user-images.githubusercontent.com/20410120/179145132-393e6592-7369-4f33-abcc-3f9a3ec541c0.png)

	- Select default project "MyTodo.Data.EntityFramework"
	![image](https://user-images.githubusercontent.com/20410120/179145207-9658e3b3-0a3d-4ebe-8b9d-10d53ad04b3b.png)

	- Run "Update-Database" to generate database
	![image](https://user-images.githubusercontent.com/20410120/179145310-57d9f0a6-b937-478a-b5e6-a7c63cb2d80a.png)

- Start project (Press F5)
![image](https://user-images.githubusercontent.com/20410120/179145505-19383655-9666-4cc6-9413-a39bdd3ca623.png)

# Run tests
- Solution level
![image](https://user-images.githubusercontent.com/20410120/179145780-bfdf8266-ecb9-482a-86d9-682d8846b39e.png)
![image](https://user-images.githubusercontent.com/20410120/179145896-f39a7572-7263-46c7-aeab-16a9101332dd.png)

- Project level
![image](https://user-images.githubusercontent.com/20410120/179146019-d2aac915-04ef-4c2e-ad01-99a20a46b8b5.png)
![image](https://user-images.githubusercontent.com/20410120/179146223-51e04a21-cc55-470b-b4b8-59c682df18c6.png)

- Function level
![image](https://user-images.githubusercontent.com/20410120/179146107-61b8b6d7-0e54-4e0a-8229-05cd6f82df8e.png)
![image](https://user-images.githubusercontent.com/20410120/179146153-41012960-9ae7-4fd6-85a9-fcfb4e791c44.png)

# Todos
- Testing bussiness rules of tasks managment as Task Limit exceeded, Assign permissions,...
- User management, Roles management
- Logging
- Caching
- etc.
