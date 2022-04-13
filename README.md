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

### How to run your code locally?

- Prerequisites: Docker
- Run postgresql:

```bash
docker-compose up -d
```

- Installation:

```bash
npm install
```

- Run app:

```bash
npm start
```

### Sample curl command to call API

- Register new user

```curl
curl -X POST "http://localhost:8080/v1/auth/register" -H "Content-type: application/json" -d '{"username": "togo", "password": "123456"}'
```

- Login user

```curl
curl -X POST "http://localhost:8080/v1/auth/login" -H "Content-type: application/json" -d '{"username": "togo", "password": "123456"}'
```

- Create task

```curl
curl -X POST "http://localhost:8080/v1/task" -H "Content-type: application/json" -H "Authorization: Bearer + token" -d '{"title": "togo", "content":"togo", "startDate":"2022-04-12"}'
```

### How to run your tests locally

- Run unit tests:

```bash
npm run test:unit
```

- Run integration tests:

```bash
npm run test:component
```

- Run both test cases:

```bash
npm run test
```

### What do you love about your solution

- I applied nestjs - good framework with DI and clean architecture to build this challenge.
- For integration testing, I decided to use sqlite for lightweight.

### Notes

- We're using Golang at Manabie. **However**, we encourage you to use the programming language that you are most comfortable with because we want you to **shine** with all your skills and knowledge.

### How to submit your solution?

- Fork this repo and show us your development progress via a PR

### Interesting facts about Manabie

- Monthly there are about 2 million lines of code changes (inserted/updated/deleted) committed into our GitHub repositories. To avoid **regression bugs**, we write different kinds of **automated tests** (unit/integration (functionality)/end2end) as parts of the definition of done of our assigned tasks.
- We nurture the cultural values: **knowledge sharing** and **good communication**, therefore good written documents and readable, organizable, and maintainable code are in our blood when we build any features to grow our products.
- We have **collaborative** culture at Manabie. Feel free to ask trieu@manabie.com any questions. We are very happy to answer all of them.

Thank you for spending time to read and attempt our take-home assessment. We are looking forward to your submission.
