### Requirements

- Implement one single API which accepts a todo task and records it
  - There is a maximum **limit of N tasks per user** that can be added **per day**.
- Write integration (functional) tests
- Write unit tests
- Choose a suitable architecture to make your code simple, organizable, and maintainable
- Write a concise README
  - How to run your code locally?
  - A sample “curl” command to call your API
  - How to run your unit tests locally?
  - What do you love about your solution?
  - What else do you want us to know about however you do not have enough time to complete?

### Setup and running

- At root path run command ``` go run main.go```
- List api: 
```
POST   /local/api/auth/login    
Sample payload: 
  {
    "email": "admin@gmail.com",
    "password": "123456"
  }
  
GET    /local/api/task
POST   /local/api/task
Sample payload: 
  {
    "content": "test"
  }
  
PUT    /local/api/task/:id
DELETE /local/api/task/:id
```
- Setup postman:
```
Headers
Authorization: jwt key
```

