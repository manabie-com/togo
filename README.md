### Instructions
- This app is can be used using Java 8/11
- This is running using H2 Database, an In-Memory database which means that restarting the app will restart the data
- to run, first execute "mvn clean install" then "mvn spring-boot:run"
- I have added a postman collection for authentication and task endpoint: https://github.com/plurantee/togo/blob/master/postman%20collection/Togo%20Collection.postman_collection.json
- default user is "florante" and password is "password"
- Sample curl command for user registration:
  ``
  curl --location --request POST 'http://localhost:8080/register' \
  --header 'Content-Type: application/json' \
  --header 'Cookie: JSESSIONID=07446863A868BC86414DAFC270DAC644' \
  --data-raw '{
  "username": "user123",
  "password": "password",
  "limit": 2
  }'
  ``
- Sample login - this will return a token response, and it must be used for the task endpoints:
  ``
  curl --location --request POST 'http://localhost:8080/auth' \
  --header 'Content-Type: application/json' \
  --header 'Cookie: JSESSIONID=07446863A868BC86414DAFC270DAC644' \
  --data-raw '{
  "username": "florante",
  "password":  "password"
  }'
  ``

- Get tasks for a user (task is filtered using the JWT token)
  ``
  curl --location --request GET 'localhost:8080/api/tasks' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJmbG9yYW50ZSIsImV4cCI6MTY1NDA3ODgxMywiaWF0IjoxNjU0MDYwODEzfQ.NeOfAPH-vo44HAfghOPBqsd8SucPNRMRjY0Qy-afIYP1uepaVHPR7GvqebUDwdgGhSM1oqXBh02rZonFV7xG8g' \
  --header 'Cookie: JSESSIONID=07446863A868BC86414DAFC270DAC644'
  ``
- To Run Unit tests and Integration Tests, execute "mvn clean install" and the testing result will be in the logs

### About the project
- What I like about my journey developing this is I have learned a lot of stuff. I was enjoying building this app while also learning some stuff I didn't know until I did this exam.
- Total Working Hours: 5 Hours
- Email me at: rapioflorante1@gmail.com if you have concerns running this app

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

