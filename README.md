### Installation
- Install the packages
  ```
    npm install
  ```

- Start project
  ```
    npm start
  ```

## Run test
  ```
    npm test
  ```

## Run API
- User:
  - GET all users: 
    ```
      curl --location --request GET 'http://localhost:8004/users/v1.1/get-all'
    ```
  - ADD new user to assign task
    ```
      curl --location --request POST 'http://localhost:8004/users/v1.1/sign-up' \
      --header 'Content-Type: application/json' \
      --data-raw '{
          "userName": "memberxx",
          "email": "member1@gmail.com",
          "password": "12345644",
          "dailyTaskLimit": 2
      }' 
    ```

- Task:
  - GET all tasks: 
    ```
      curl --location --request GET 'http://localhost:8004/operation/v1.1/tasks'
    ```
  
  - Create/update (upsert) task and assign this to user
    ```
      curl --location --request POST 'http://localhost:8004/operation/v1.1/tasks/assign' \
      --header 'Content-Type: application/json' \
      --data-raw '{
          "taskId": null,
          "userId": "624c30a8a8943abb36a608fc",
          "taskName": "Task 2",
          "taskCode": "task_2",
          "taskDescription": "Check out",
          "status": "in-process"
      }'
    ```

  - GET all users and tasks
    ```
      curl --location --request GET 'http://localhost:8004/operation/v1.1/tasks/by-user'
    ```
  
  - GET user's tasks: add filter userId to query param
    ```
      curl --location --request GET 'http://localhost:8004/operation/v1.1/tasks/by-user?userId=624c30a8a8943abb36a608fc'
    ```
