# **TOGO**

## Dscription
  **API using GO, MongoDB and jwt for authentication**

### **1. How to run your code locally?**
  -Git clone repository
  ```
  git clone https://github.com/glucozo192/todo
  ```
  
  -Dependencies Installation
  ```
  go mod download
  ```
  -Configuration Server
  
  creat ```.env```
  ```
  MONGOURI=mongodb+srv://glucozo:6677028a@cluster0.2ciga69.mongodb.net/?retryWrites=true&w=majority
  SECRET_JWT=glucozoisbestbestbestbestbest
  ```
  
  -Run sever
  ```
  go run TOGO
  ```
### **2. Sample “curl” Command:**

  - Token: "Bearer token"
  
  - Respon:
   ```json
   {
   "status": 200,
   "token": ,
   "message: "success",
   "data":
   }
   ```
#### - Authentication
- Signup:
  ```
  curl --location --request POST 'http://localhost:9099/user/signup' \
  --data-raw '{
    "username":"example1",
    "password":"123456",
    "name":"Nguyen van tuan"
  }'
  ```
    
- Login:
  ```
  curl --location --request POST 'http://localhost:9099/user/login' \
  --data-raw '{
    "username": "example1",
    "password": "123456"
  }'
  ```
#### - User
  
 - Get Me
    ```
    curl --location --request GET 'http://localhost:9099/me' \
    -header 'Authorization: Bearer <YOUR TOKEN>'
    ```
 - Update Me
    ```
    curl --location --request PUT 'http://localhost:9099/user' \
    --header 'Authorization Bearer <YOUTOKEN>;' \
    --data-raw '{
      "name":"<YOUR_NAME>",
      "password": "YOUR_PASSWORD"
    }'
    ```
  - Upgrade Premium
    ```
    curl --location --request PUT 'http://localhost:9099/limit' \
    --header 'Authorization Bearer <YOUR TOKEN> ;'
### - Admin-User
  ```
  "username: "admin"
  "password: "123456"
  ```
- Get All User
  ```
  curl --location --request GET 'http://localhost:9099/users?Authorization Bearer <YOUR TOKEN>'
  ```
- Get One User
  ```
  curl --location --request GET 'http://localhost:9099/user/<USER_ID>?Authorization Bearer <YOUR TOKEN>'
  ```
 - Delete User
  ```
  curl --location --request DELETE 'http://localhost:9099/user/<USER_ID>?Authorization Bearer <YOUR TOKEN>'
  
  ```
  
 ### - Task
- Create Task
    ```
    curl --location --request POST 'http://localhost:9099/task' \
  --header 'Authorization Bearer <YOU TOKEN;' \
  --data-raw '{
      "name":"<YOUR_NAME>",
      "content":"<YOUR_CONTENT>"
   }'
    ```  
- Get Tasks By User
  ```
  curl --location --request GET 'http://localhost:9099/user-tasks' \
  --header 'Authorization Bearer <YOUR TOKEN>;'
  ```
- Get Task By Id
  ```
  curl --location --request GET 'http://localhost:9099/task/<TASK_ID>?Authorization Bearer <YOUR TOKEN>'
  ```
- Get All Task Doing
  ```
  curl --location --request GET 'http://localhost:9099/task-status?Authorization Bearer <YOUR TOKEN>'
  ```
- Update Task
  ```
  curl --location --request PUT 'http://localhost:9099/task/<TASK_ID>?Authorization Bearer <YOUR TOKEN>' \
  --data-raw '{
    "name":"<TASK_NAME>",
    "content":"<CONTENT>"
  }'
  ```
- Update Task Status
  ```
  curl --location --request PUT 'http://localhost:9099/task/status/<TASK_ID>' \
  --header 'Authorization Bearer <YOUR TOKEN>;' \
  --data-raw '{
    "status": "completed"
  }'
- Delete Task
  ```
  curl --location --request DELETE 'http://localhost:9099/task/<TASK_ID>?Authorization Bearer <YOUR TOKEN>'
  ```
### 3.How to run your unit tests locally?
- Go to togo directory
- Command:``` go test ./... ```

### 4.What do you love about your solution?
I love everything about my code, so great when coding with golang! :heartbeat::heartbeat::heartbeat::heartbeat::heartbeat:
