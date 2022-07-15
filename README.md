## Prerequisite

- Golang 1.16
- Make
- Docker/docker-compose

## How to run

- Docker startup, run :
  `make docker/up`

- Database migration, run:
  `make db/up`

- Start server, run:
  `make run`

## Tools I used

- sqlboiler: this library generate repositories that handle database operations, with much faster speed than casual ORM
- migrate: this is used to run migration scripts, written in sql queries, to manage database schema easily

## My Design

My code base was designed following clean architecture with mainly 3 layers:

- UseCase: this layer is used to handle business logic
- Infrastructure: this layer contains packages related to infrastructure and does not involve into business logic, including database, midlewares, etc.
- Handler: this is the handler layer that expose API endpoint

### Directory Structure

- cmd: implemented necessary cmd, including main server application
- internal: wrapped and encapsulated source of code that will not be exposed
- db: includes sqlboiler configuration and migrations scripts
- docker: includes dockerfile and docker-compose file as well as neccessary configuration files
- bin: mostly executable plugins and generated/built applications
- scripts: bash scripts to handle operations quickly

## What I have done and have not done

- [ ] Covering the core functionality with unit test
      (As not enough time, I have implemented an unit test to demonstrate how I do test)

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
