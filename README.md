### About this project
This is a web service that provides APIs that make it possible for anyone to manage your own work, and the system can set a limit on the number of tasks that can be added to the system per day.

Technical use in this project: Nodejs, express, Sequelize, postgress sql, joi, log4js, bluebird, jsonwebtoken,...

database schema
![data base chema](./images/database-schema.png)

### details about the APIs on the system
Currently the APIs are developing so there are a few bugs, we will improve in the next release. you can test the APIs on the production server [here](https://togotask.herokuapp.com/)  or run locally.
1.  **Login:** user login to system
* path: {{host}}/api/accounts/login
* method: post
* example params: 
```
{
    "username": "free",
     "password": "free"
}
```
* curl
```
curl --location --request POST 'https://togotask.herokuapp.com/api/accounts/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "free",
     "password": "free"
}'
```
* error code

| code | message|
| -------- | -------- |
| 400     | child "username" fails because ["username" is required]     | 
| 400     |child "username" fails because ["username" must be a string]    | 
| 400     | child "username" fails because ["username" length must be at least 3 characters long]    | 
| 400     | child "username" fails because ["username" length must be less than or equal to 50 characters long]    | 
| 400     | child "password" fails because ["password" is required]     | 
| 400     | child "password" fails because ["password" must be a string]     | 
| 401    | Username or password incorrect     | 


3.  **refresh token:** get new access token when access token expire
* path: {{host}}/api/accounts/refreshtoken
* method: post
* example params: 
```
{
    "token": "refreshToken"
}
```
* curl
```
curl --location --request POST 'https://togotask.herokuapp.com/api/accounts/refreshtoken' \
--header 'Content-Type: application/json' \
--data-raw '{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6ImZyZWUiLCJ1c2VySWQiOjEsImlhdCI6MTY1MTE2Mjg2Nn0.Otre_7q1xHxBRzxiTc80GGr9bcJpLd--YQBnyAEiYnE"
}'
```
* error code

| code | message|
| -------- | -------- |
| 400     |child "token" fails because ["token" is required]     | 
| 400     |child "token" fails because ["token" length must be at least 10 characters long]  | 
| 403    |    | 

5.  **add task:** add new task for current user
* path: {{host}}/api/tasks
* method: post
* example params: 
```
{
    "name": "accc"
}
```
* curl
```
curl --location --request POST 'https://togotask.herokuapp.com/api/tasks' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6ImZyZWUiLCJ1c2VySWQiOjEsImlhdCI6MTY1MTE2Mjg2OSwiZXhwIjoxNjUxMTY0MDY5fQ.pXRaFrpWV-My60H9hBcs5OegWBNQ2_5PT7fQ5E5lIO0' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "accc"
}'
```
* error code

| code | message|
| -------- | -------- |
| 400     | child "name" fails because ["name" is required]    | 
| 400     | child "name" fails because ["name" must be a string]  | 
| 400     | child "name" fails because ["name" is not allowed to be empty]  | 
| 400     | child "name" fails because ["name" length must be less than or equal to 100 characters long]  | 
| 400     | User's tasks are limited  | 
| 401    |    | 
| 403    |    | 

### How to run locally
currently, the project configures to connect to the Heroku database server. you can configure the database locally in config/local.json.

Clone source code and run locally:
```
git clone https://github.com/manabie-com/togo.git
cd togo
npm install
npm start
```
in addition, you can run unit tests locally
```
npm run test
```

### API test
i am use post man collection to implement api testing. you can import postment collect from https://www.getpostman.com/collections/9d93c1b60ad16231b0f4 and run test.

** Note: You can update the server variable to change the server 

#### What do you love about your solution?

*  using jsonwebtoken to authentication and make sure user logged in an account on system.
*  each peoplen just congire timezone one time, every time add new task system automatic convert time to timezone of current user (if system just use local you do need care about this because by default system get timezone of server).
*  user do need sent local time to server.
*  sever save correct time in spite of client side and server side in different timezone.
*  use mechanism refresh token. user do not login again many times.

-----

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

### Notes

- We're using Golang at Manabie. **However**, we encourage you to use the programming language that you are most comfortable with because we want you to **shine** with all your skills and knowledge.

### How to submit your solution?

- Fork this repo and show us your development progress via a PR

### Interesting facts about Manabie

- Monthly there are about 2 million lines of code changes (inserted/updated/deleted) committed into our GitHub repositories. To avoid **regression bugs**, we write different kinds of **automated tests** (unit/integration (functionality)/end2end) as parts of the definition of done of our assigned tasks.
- We nurture the cultural values: **knowledge sharing** and **good communication**, therefore good written documents and readable, organizable, and maintainable code are in our blood when we build any features to grow our products.
- We have **collaborative** culture at Manabie. Feel free to ask trieu@manabie.com any questions. We are very happy to answer all of them.

Thank you for spending time to read and attempt our take-home assessment. We are looking forward to your submission.
