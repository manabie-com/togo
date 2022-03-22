### To run code locally

- Run sql script
- Download and install .NET SDK (link below).
	- https://dotnet.microsoft.com/en-us/download/dotnet/5.0
- Clone source code.
- Change the connection string in the appSetting.json.
- Open terminal in folder contains WebAPI.csproj and run the command: dotnet run .\WebAPI.csproj

### Sample “curl” command to call API
- register api
	- curl --location --request POST 'http://localhost:5000/api/users/register' \
			--header 'Content-Type: application/json' \
			--data-raw '{
				"Username": "Admin",
				"Password": "12345678",
				"TaskPerDay": 10
			}'

- login api
	- curl --location --request POST 'http://localhost:5000/api/users/login' \
			--header 'Content-Type: application/json' \
			--data-raw '{
				"Username": "Admin",
				"Password": "12345678"
			}'

- create task api
	- curl --location --request POST 'http://localhost:5000/api/tasks/create' \
			--header 'Content-Type: application/json' \
			--header 'Authorization: Bearer <accessToken from login api>' \
			--data-raw '{
				"Content": "Task Content"
			}'
			
### Run unit tests locally
- Open terminal in folder contains UnitTests.csproj and run the command: dotnet test .\UnitTests.csproj

### What do i love about my solution?
- The first time, I build a project from scratch.
- Generate RSA access token.
- The structure of the project.

### What else do you want us to know about however you do not have enough time to complete?
- I will make a better response (Error Code, Message,...)
- Seperate security code into another service.
- Update security code to generate refresh token and more information.
- Design a better database with more information.
- Use an api gateway (ocelot) to implement microservice.