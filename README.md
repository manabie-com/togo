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

### To run code locally
#### Database
- For the development phase, please define your database name at: conf/app.ini
- For testing, define database name at: test/conf/app.ini
#### Run from the command line
```shell
go mod vendor
go run main.go
```

### A sample “curl” command to call API
#### Login curl
curl --location --request POST 'http://localhost:8003/api/sign-in' \
--header 'Content-Type: application/json' \
--data-raw '{
"username": "test1",
"password": "123456"
}'
#### Add Task curl
curl --location --request POST 'http://localhost:8003/api/task' \
--header 'Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRHVnpkREU9Iiwicm9sZSI6ImJXVnRZbVZ5IiwidXNlcl9pZCI6MSwiZXhwIjoxNzA5ODE5NjQ4LCJpc3MiOiJ2c2Nob29sLWFwaSJ9.7mIeIsI-7TM6Z_2P_KadUD0ifY8c7LOQrjqnrT7BkzWuevEcMBUqwQZMQA0ETo2SPuMIyaaIHsbbcD_cJ951UA' \
--header 'Content-Type: application/json' \
--data-raw '{
"name": "task name"
}'

### How to run unit tests locally
- We have a unit test at: modules/member/handler 
- Run from the command line
```
cd modules/member/handler
go test
```
```
cd modules/tast/service
go test
```

### Things else
- I have try to make 
- The test case that exceeded the limit of added tasks per day was just suitable for my prepared data. I think the best should cover for any setting of limitation.
