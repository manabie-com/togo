
### Things I have used in this sample

- Dot NET 6, EF Core
- MediatR, CleanArchitecture, OpenIddict for OAuth 2 standart
- NUnit

### about Projects:
- BasicIdentityServer: an Identity Provider Server use for granting access token.
- TestingApi: Web Api.
- Using InMemory Database.

### Setup Prepare Resource:
- Make sure you have installed .NET SDK 6 and .NET RUNTIME 6 .
- Run "PublicApp.bat" in the folder solution.
- Run "RunIdentity.bat" and "RunTestingApi.bat".

- Try to setup docker but have an issue with communication between dockers.

### Step Run

- Main flow:
	- GET access_token:

		- Adminitrator account:

			curl --location --request POST 'http://localhost:5000/api/connect/token' \
				--header 'Content-Type: application/x-www-form-urlencoded' \
				--data-urlencode 'grant_type=password' \
				--data-urlencode 'username=administrator@localhost' \
				--data-urlencode 'password=Administrator1!'
			
			It's should return a JWT with role is "administrator" using for access resource application.
			Administrator only can add 5 tasks per day.
			
		- User account:
		
			curl --location --request POST 'http://localhost:5000/api/connect/token' \
				--header 'Content-Type: application/x-www-form-urlencoded' \
				--data-urlencode 'grant_type=password' \
				--data-urlencode 'username=user@localhost' \
				--data-urlencode 'password=User1!'
				
			It's should return a JWT with role is "user" using for access resource application.
			Administrator only can add 3 tasks per day.
			
	- GET Todo:

		- curl  --location --request GET 'http://localhost:7216/Todo/GetTodos' \
				--header 'Authorization: Bearer access_token
				
	- ADD Todo:

		- curl --location --request POST 'http://localhost:7216/Todo/AddTodo' \
			--header 'Authorization: Bearer access_token' \
			--header 'Content-Type: application/json' \
			--data-raw '{
				"Title": "Todo1",
				"Note": "Note1"
			}'
		If Success, you should receive a response as below:
		{
			"data": 1,
			"succeeded": true,
			"errors": [],
			"code": 0
		}
		
		If adding exceeded the limit, it should return somethings like:
		{
			"data": 0,
			"succeeded": false,
			"errors": [
				"Exceeded Limit Per Day"
			],
			"code": 0
		}
		
### Run Test:
- Open Visual Studio 2022 to run integration test.
