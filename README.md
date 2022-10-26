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

## Setup & Run Code

### Without Docker

> npm install

or

> yarn install

then

Running with DB online:

> yarn start

or

> npm start

Running locally:

>yarn dev

or

> npm run dev

Testing:

> yarn test

### With Docker

run docker container in development mode
> yarn docker:dev

run docker container in production mode
>yarn docker:prod

run all tests in a docker container
>yarn docker:test

## API Documentation

To view the list of available APIs and their specifications, run the server and go to http://localhost:6363/v1/docs in your browser.

### API Endpoints
List of available routes:

**Auth routes**:\
`POST /v1/auth/register` - register\
`POST /v1/auth/login` - login\
`POST /v1/auth/refresh-tokens` - refresh auth tokens\
`GET /v1/auth/logout`

**User routes**:\
`GET /v1/user` - get user\
`PATCH /v1/user` - update user\

**Task routes**:\
`POST /v1/todo` - create a task\
`GET /v1/todo` - get all tasks of user\
`GET /v1/task/:taskId` - get a task\
`PATCH /v1/task/:taskId` - update task\
`DELETE /v1/task/:taskId` - delete task
