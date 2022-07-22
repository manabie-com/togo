### Requirements

- Implement one single API which accepts a todo task and records it
  - There is a maximum **limit of N tasks per user** that can be added **per day**.
  - Different users can have **different** maximum daily limit.
- Write integration (functional) tests
- Write unit tests
- Choose a suitable architecture to make your code simple, organizable, and maintainable
- Using Docker to run locally
  - Using Docker for database (if used) is mandatory.
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

### How to run
## 1. Start database locally
```sh
docker-compose up --build -d
```
## 2. Run code locally
- Switch to `develop` branch

- Prepare
```sh
cp .env.example .env

npm i
```

- Migration
```sh
npm run typeorm migration:run
```

- Run
```sh
npm run start:dev
```

## 3. Curl
- Create user
```curl
curl --location --request POST 'localhost:3000/users' \
--header 'Content-Type: application/json' \
--header 'Accept: application/json' \
--data-raw '{
  "username": "username1",
  "firstName": "firstName1",
  "lastName": "lastName1"
}'
```

- Create todo (Get userId from above step)
```curl
curl --location --request POST 'localhost:3000/users/161a6144-0a21-4462-bbca-31b6b6b43831/todos' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "test",
    "date": "2022-07-21"
}'
```

- Update number of todos per day (Get userId from above step)
```curl
curl --location --request PUT 'localhost:3000/users/161a6144-0a21-4462-bbca-31b6b6b43831/settings' \
--header 'Content-Type: application/json' \
--header 'Accept: application/json' \
--data-raw '{
    "todoPerday": 1
}'
```

## 4. Testing
- Unit test
```sh
npm run test
```

- Intergration test
```sh
npm run test:e2e
```

## 5. What do you love about your solution?
- I follow RESTfull api guide
- I provide swagger for apis
- Code guarantee that number of todo never exceeded

## 6. What else do you want us to know about however you do not have enough time to complete?
- Add more validate for input: ex: title is limited by 255 characters

