### Requirements

- Implement one single API which accepts a todo task and records it
  - There is a maximum **limit of N tasks per user** that can be added **per day**.
  - Different users can have **different** maximum daily limit.
- Write integration (functional) tests
- Write unit tests
- Choose a suitable architecture to make your code simple, organizable, and maintainable
- Write a concise README
  - How to run your code locally?
  - A sample “curl” command to call your API
  - How to run your unit tests locally?
  - What do you love about your solution?
  - What else do you want us to know about however you do not have enough time to complete?

### Things I have used in this sample

- Dot NET 6, EF Core
- MediatR, CleanArchitecture, OpenIddict for OAuth 2 standart

### about Projects:
- BasicIdentityServer: an Identity Provider Server use for granting access token.
- TestingApi: Web Api.
- Using InMemory Database.

### Step to run

- Make sure that you have installed .Net 6 .Net Runtime and .Net SDK

- Main flow:
	- GET access_token:

		- Adminitrator account:

			curl --location --request POST 'https://localhost:7173/api/connect/token' \
				--header 'Content-Type: application/x-www-form-urlencoded' \
				--data-urlencode 'grant_type=password' \
				--data-urlencode 'username=administrator@localhost' \
				--data-urlencode 'password=Administrator1!'
			
			It's should return a JWT with role is "administrator" using for access resource application.
			Administrator only can add 5 tasks per day.
			
		- User account:
		
			curl --location --request POST 'https://localhost:7173/api/connect/token' \
				--header 'Content-Type: application/x-www-form-urlencoded' \
				--data-urlencode 'grant_type=password' \
				--data-urlencode 'username=user@localhost' \
				--data-urlencode 'password=User1!'
				
			It's should return a JWT with role is "user" using for access resource application.
			Administrator only can add 3 tasks per day.
			
	- GET Todo:

		- curl  --location --request GET 'https://localhost:xxxx/Todo/GetTodos' \
				--header 'Authorization: Bearer access_token
				
	- ADD Todo:

		- curl --location --request POST 'https://localhost:xxxx/Todo/AddTodo' \
			--header 'Authorization: Bearer access_token' \
			--header 'Content-Type: application/json' \
			--data-raw '{
				"Title": "Todo1",
				"Note": "Note1"
			}'
		If Success, you should recive a response as below:
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
