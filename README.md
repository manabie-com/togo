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

### Curl Commands

- The routes will require token authorization to proceed.

To start, create a sample user.

    curl -c cookie.txt -X POST http://localhost:5000/register -H "Content-Type: application/json" -d '{"name": "sample_user", "password": "sample_password", "limit_per_day":1}'

- Login test

Login to retrieve a token for authorization.

    curl --user sample_user:sample_password http://127.0.0.1:5000/login

- Create Todo

Create a new todo.

    curl -b cookie.txt POST http://localhost:5000/todo -H "Content-Type: application/json" -H "x-access-token:<user_id>" -d '{"todo": "sample_task"}'



### Running Tests

- Tests are done using pytest
  In order to run the test, you must be at the root folder and run the following command:

  python3 -m pytest

### Things to improve

Due to the lack of time, it would be better if the user id is stored in the cookie, since the flask login is 
causing issue when unit testing (login per endpoint is required before execution since current_user is being
anonymous). Curl commands are also affected since in the terminal, sessions are not being stored. Although, things
are working fine in postman.