1. How to run your code locally?
- To make sure that there are no errors that relate to NodeJS version happens on local machine, I created a docker-compose.yaml file which contains a both web-app and postgres-database components. So, just type that command "docker-compose up -d" to run the code.

2. A sample “curl” command to call your API?
- There are 3 API enpoints:
+ Signin: curl -X POST http://localhost:3000/api/v1/signin -d '{"email": "admin@gmail.com", "password": "admin123"}' -H "Content-Type: application/json" -> This user information already added into database during app starting time.

+ Create user: curl -X POST http://localhost:3000/api/v1/users -d '{"email": "steven@gmail.com", "password": "123456", "maxTasks": 6 }' -H "Content-Type: application/json" -H "Authorization: Bearer token"

+ Create task: curl -X POST http://localhost:3000/api/v1/tasks -d '{ "content": "Task One" }' -H "Content-Type: application/json" -H "Authorization: Bearer token"

* Note: 
+ To call `create user` API, you must login with admin account {"email": "admin@gmail.com", "password": "admin123"}.

3. How to run your unit tests locally?
- At the root folder, just type "yarn run:test" command.

4. What do you love about your solution?
- Users are able to add new tasks on next day if their number of tasks has reached the limit today. 
Example: There is one user who has the daily limit task is 5. On 20/05/2022, the number of tasks of this user is reached the limit. On 21/05/2022, this user can create 5 more tasks.

5. What else do you want us to know about however you do not have enough time to complete?
I would like to continue to add e2e testing cases more. Beside that, I would also like to add a new feature to check the number of tasks has been finished in a day to decide whether user is allowed to add more tasks on next day or not.