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

-------------------------------------------------------------------------------------------------------------------
1. Download and install Git bash

2. Clone repository https://github.com/bernardcvu/togo.git

3. Build application using below command:
./gradlew bootRun

4. Hit http://localhost:7000/. You'll be greeted with a rather dull error message, but you're up and running alright!

Test the application:

5. Add items with below command:

curl -X POST -H "Content-Type:application/json" -d "{\"userId\":\"bernard\", \"taskName\":\"Create\", \"taskDescription\":\"Create Record\"}" http://localhost:7000/api/tasks/todo -i

curl -X POST -H "Content-Type:application/json" -d "{\"userId\":\"bernard\", \"taskName\":\"Retrieve\", \"taskDescription\":\"Retrieve Record\"}" http://localhost:7000/api/tasks/todo -i

curl -X POST -H "Content-Type:application/json" -d "{\"userId\":\"bernard\", \"taskName\":\"Update\", \"taskDescription\":\"Update Record\"}" http://localhost:7000/api/tasks/todo -i

curl -X POST -H "Content-Type:application/json" -d "{\"userId\":\"bernard\", \"taskName\":\"Delete\", \"taskDescription\":\"Delete Record\"}" http://localhost:7000/api/tasks/todo -i

#7. Get all items:
#curl http://localhost:7000/api/tasks/todo -i

8. Get item/s:
curl http://localhost:7000/api/tasks/todo/?userId=bernard&date=24-07-2022 -i

9. Delete an item:
curl -X DELETE http://localhost:7000/api/tasks/todo/{userId} -i

#10. Update an item:
#curl -X PUT -H "Content-Type:application/json" -d "{\"userId\":\"bernard\", \"taskName\":\"Test Update\", \"taskDescription\":\"Test Update\"}" http://localhost:7000/api/tasks/todo/{userId} -i
