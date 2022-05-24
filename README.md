### Requirements

- Implement one single API which accepts a todo task and records it
  - There is a maximum **limit of N tasks per user** that can be added **per day**.
  - Different users can have **different** maximum daily limit.
- Write integration (functional) tests
- Write unit tests
- Choose a suitable architecture to make your code simple, organizable, and maintainable
- Write a concise README
  - How to run your code locally?
    In my computer nodejs version is 8.11.4
    > 1: Get source code from this branch
    > 2: intermial, cd to "togo" folder
    > 3: for first time, set variable environment: export todolist_jwtPrivateKey=<anywordsyouwant>
    > 4: install dependencies: Run "npm i"
    > 5: for Dev mode: run "nodemon"
  - A sample “curl” command to call your API
    > Register user:
        exam: curl -v -d '{"name":"testcurl2", "email":"testcurl2@gmail.com", "password": "12345"}' -H "Content-Type: application/json" -X POST http://localhost:3000/api/users
    > Add todo:
        curl -d '{"name":"valueofname",  "user_id": "getfromabovecommand"}' -H "Content-Type: application/json" -H "x-auth-token:getfromabovecommand"  -X POST http://localhost:3000/api/todolist

        exam: curl -d '{"name":"todotest1",  "user_id": "628c7fb3d5c04d0c8005856a"}' -H "Content-Type: application/json" -H "x-auth-token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2MjhjN2ZiM2Q1YzA0ZDBjODAwNTg1NmEiLCJpYXQiOjE2NTMzNzQ4OTl9.9NgJnxEbjRhnoeNExZY5zRBiosmUQRHUBJQi4kE0R54"  -X POST http://localhost:3000/api/todolist 
  - How to run your unit tests locally?
      > in "togo" folder: run "npm test"
  - What do you love about your solution?
      > In middleware "async.js" have some code support when coding routes no need try catch since used them many times, it noisy.
      > structure project in backend may be completed, so it only to need focus logic, bussines app.
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
