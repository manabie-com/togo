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

### Step to run

- Make sure that you have installed .Net 6 .Net Runtime and .Net SDK

- Start 2 Project:
	- BasicIdentityServer: an Identity Provider Server use for granting access token.
	- TestingApi: Web Api.

- Main flow:
	- GET access_token:

		- Adminitrator account:

			curl --location --request POST 'https://localhost:7173/api/connect/token' \
				--header 'Content-Type: application/x-www-form-urlencoded' \
				--data-urlencode 'grant_type=password' \
				--data-urlencode 'username=administrator@localhost' \
				--data-urlencode 'password=Administrator1!'
			
		- User account:
		
			curl --location --request POST 'https://localhost:7173/api/connect/token' \
				--header 'Content-Type: application/x-www-form-urlencoded' \
				--data-urlencode 'grant_type=password' \
				--data-urlencode 'username=user@localhost' \
				--data-urlencode 'password=User1!'
			
	- GET Todo:

		- curl  --location --request GET 'https://localhost:7215/Todo/GetTodos' \
				--header 'Authorization: Bearer access_token
				
	- ADD Todo:

		- curl --location --request POST 'https://localhost:7215/Todo/AddTodo' \
			--header 'Authorization: Bearer access_token' \
			--header 'Content-Type: application/json' \
			--data-raw '{
				"Title": "Todo1",
				"Note": "Note1"
			}'
			
#Run Test:
