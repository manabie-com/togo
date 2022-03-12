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

### Run code locally

```sh
~ git clone https://github.com/Legacy107/togo.git
~ cd togo
~ npm i
~ npm run start
```

### Sample usage

#### create a new user
```sh
~ curl --location --request POST 'localhost:5050/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "firstUser",
    "password": "example",
    "limitPerDay": 5
}'
```

#### login
```sh
~ curl --location --request POST 'localhost:5050/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "firstUser",
    "password": "example"
}'
# response:
{"access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImZpcnN0VXNlciIsImlhdCI6MTY0NzA5MzEzNSwiZXhwIjoxNjQ3MTA3NTM1fQ.5Xn8breH_VuidyagX9iLbgXpYNrDqHZjtmMwTH-s1SM"}
```

#### create todo
```sh
curl --location --request POST 'localhost:5050/tasks' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImZpcnN0VXNlciIsImlhdCI6MTY0NzA5MzEzNSwiZXhwIjoxNjQ3MTA3NTM1fQ.5Xn8breH_VuidyagX9iLbgXpYNrDqHZjtmMwTH-s1SM' \
--header 'Content-Type: application/json' \
--data-raw '{
    "content": "first todo"
}'
# response
{"data":{"content":"first todo","id":"303fade2-ea8c-4259-87a5-d7dff27d4844","status":"ACTIVE","createAt":"2022-03-12T13:54:56.036Z","updateAt":"2022-03-12T13:54:56.036Z"}}
```

#### get user's todos
```sh
curl --location --request GET 'localhost:5050/tasks' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImZpcnN0VXNlciIsImlhdCI6MTY0NzA5MzEzNSwiZXhwIjoxNjQ3MTA3NTM1fQ.5Xn8breH_VuidyagX9iLbgXpYNrDqHZjtmMwTH-s1SM'
# response
{"data":[{"id":"303fade2-ea8c-4259-87a5-d7dff27d4844","content":"first todo","status":"ACTIVE","createAt":"2022-03-12T13:54:56.036Z","updateAt":"2022-03-12T13:54:56.036Z"}]}
```

#### get user's todos with status
```sh
curl --location --request GET 'localhost:5050/tasks?status=completed' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImZpcnN0VXNlciIsImlhdCI6MTY0NzA5MzEzNSwiZXhwIjoxNjQ3MTA3NTM1fQ.5Xn8breH_VuidyagX9iLbgXpYNrDqHZjtmMwTH-s1SM'
# response
{"data":[]}
```

#### update todo's status
```sh
curl --location --request PUT 'localhost:5050/tasks' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImZpcnN0VXNlciIsImlhdCI6MTY0NzA5MzEzNSwiZXhwIjoxNjQ3MTA3NTM1fQ.5Xn8breH_VuidyagX9iLbgXpYNrDqHZjtmMwTH-s1SM' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "303fade2-ea8c-4259-87a5-d7dff27d4844",
    "status": "COMPLETED"
}'
```

#### update many todos' status
```sh
curl --location --request PUT 'localhost:5050/many-tasks' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImZpcnN0VXNlciIsImlhdCI6MTY0NzA5MzEzNSwiZXhwIjoxNjQ3MTA3NTM1fQ.5Xn8breH_VuidyagX9iLbgXpYNrDqHZjtmMwTH-s1SM' \
--header 'Content-Type: application/json' \
--data-raw '{
    "ids": [
        "303fade2-ea8c-4259-87a5-d7dff27d4844"
    ],
    "status": "COMPLETED"
}'
```

#### delete todo by Id
```sh
curl --location --request DELETE 'localhost:5050/tasks' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImZpcnN0VXNlciIsImlhdCI6MTY0NzA5MzEzNSwiZXhwIjoxNjQ3MTA3NTM1fQ.5Xn8breH_VuidyagX9iLbgXpYNrDqHZjtmMwTH-s1SM' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "303fade2-ea8c-4259-87a5-d7dff27d4844"
}'
```

#### delete all user's todos
```sh
curl --location --request DELETE 'localhost:5050/tasks' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImZpcnN0VXNlciIsImlhdCI6MTY0NzA5MzEzNSwiZXhwIjoxNjQ3MTA3NTM1fQ.5Xn8breH_VuidyagX9iLbgXpYNrDqHZjtmMwTH-s1SM'
```

### Run test locally

#### unit test
```sh
~ npm run test
```

#### e2e test
```sh
~ npm run test:e2e
```
