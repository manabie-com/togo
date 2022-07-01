## Todo API

- Link API document: https://basalt-leech-fa6.notion.site/Todo-API-document-fcfd66100f6743b2b585eeb1cf84f49a
- Link API online: https://todo-api-version1.herokuapp.com/
- Link API after install local: localhost:8000

### :old_key: Prerequisites
Before you start, ensure you meet the following requirements:

 - You have installed Golang version 1.18.1
 - You have installed the Visual Studio Code.
 - You have a basic understand of Golang, CLI.
 
### :page_with_curl: Guide

#### How to run this project locally

Open Git bash
Paste folllowing command:

```
git clone https://github.com/huynhhuuloc129/togo.git
```

#### How to call API using curl

####	Install
Open the folder todo with VSCode and using this command to get all the library necessary:
```
go mod tidy
```

#### Usage
After installing enviroment, using those commands on VSCode to use 
* Note: ***REPLACE ALL THE TEXT IN <> WITH YOUR INFO***

##### :one: Start server
* The server must be keep open all the time you use the app
```
go run server.go
```
##### :two: Register
Open another CLI to call API:
  * To use the API you need to register account first: 
```
curl --location --request POST 'localhost:8000/auth/register' \
--data-raw '{
    "username": "<yourusername>",
    "password": "<yourpassword>"
}'
```
  - The response will be something like this:
```json
{
   "Username":"<yourusername>",
   "Password":"<yourpasswordafterhashed>",
   "LimitTask":10
}
```

##### :three: Login
  * After that you can login using the same password as register to get the token of that account: 
```
curl --location --request POST 'localhost:8000/auth/login' \
--data-raw '{
    "username": "<yourusername>",
    "password": "<yourpassword>"
}'
``` 
  - The response will be something like this:
```json
{
    "Message": "login success",
    "Token": "<yourtoken>"
}
```

##### :four: Using task
* You can check your info at: 
```
curl --location --request GET 'localhost:8000/users/info' \
--header 'token: <yourtoken>'
```

###### Use the token response to you after login to:
* Get all task
```
curl --location --request GET 'localhost:8000/tasks' \
--header 'token: <yourtoken>'
```
* Get one task by task id
```
curl --location --request GET 'localhost:8000/tasks/<taskid>' \
--header 'token: <yourtoken>'
```
* Create new task
```  
curl --location --request POST 'localhost:8000/tasks' \
--header 'token: <yourtoken>' \
--data-raw '{
    "Content": "<yourtaskcontent>"
}' 
```
* Update an existing task
```
curl --location --request PUT 'localhost:8000/tasks/<taskid>' \
--header 'token: <yourtoken>' \
--data-raw '{
    "Content": "<newtaskcontent>"
}'
```
* Delete one task by task id
```
curl --location --request DELETE 'localhost:8000/tasks/<taskid>' \
--header 'token: <yourtoken>' 

```
***Beside that you can run your all of tasks and also users command under admin account***
* Login with
```
curl -X POST -H "Content-Type: application/json" -d '{"username": "admin", "password": "admin"}' "localhost:8000/auth/login"
``` 
* After logging in, you can modify users and your task:
* Get all users
```
curl --location --request GET 'http://127.0.0.1:8000/users' \
--header 'token: <admintoken>'
```
* Get one task by user id
```
curl --location --request GET 'http://127.0.0.1:8000/users/<userid>' \
--header 'token: <admintoken>'   
```
* Create new user
```  
curl --location --request POST 'http://127.0.0.1:8000/users' \
--header 'token: <admintoken>' \
--data-raw '{
    "username": "<newusername>",
    "password": "<newuserpassword>"
}'   
```
* Update an existing user
```
curl --location --request PUT 'http://127.0.0.1:8000/users/<userid>' \
--header 'token: <admintoken>' \
--data-raw '{
    "username": "<newusername>",
    "password": "<newpassword>"
}'   
```
* Delete one user by user id
```
curl --location --request DELETE 'http://127.0.0.1:8000/users/<userid>' \
--header 'token: <admintoken>'
```
#### How to run test locally
* Using this command to test all test in project (include unit test and intergration test):
```
go test -v -cover ./...
```
* If you want to test some folder only, you can go to the folder and using command:
```
cd <yourfolder>
go test . 
```
#### What special about this solution:
* This project using all dependencies in clouds, include database, even the project it self has been host through heroku for using convenient.
* It has an admin account for full admin control over users but can't seen task to keep it private for each user.
* ***Note: Because this project use a cloud database so some action may take longer than usual***
