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

#### How to run

`Simply run the main.exe in build folder`

#### Test user

Username: `dung`

Password: `dung1234`

#### Sample CURL:

`curl --location --request POST 'localhost:8800/api/togo/v1/task' \
--header 'Content-Type: application/json' \
--data-raw '{
"content": "Task 1",
"user_id": 1,
"date": "11-12-2022"
}'`

#### "date" must in type dd-mm-yyyy

#### How to run test

simple change dir to api_test folder and edit some date to make sure if it not duplicated and then

`go test -v -run TestCreateTask `

or

`go test`

#### My solution is using gorm with support a lot of types database, so it is much easier to get data which relation

#### A lot of things to do to enhance this (like cache database to avoid data reading time, doing pagination)
