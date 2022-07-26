### Requirements

- Implement one single API which accepts a togo task and records it
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

________________________________________________________________________________________________________________________

### NGO HOAI TRONG

### user guide and installation guide 

Get source after that try run "go install" and "go build" then open togo.exe to start hosting

can use Postman or just patse this curl to send request to API

change userId to another number to switch to another user

curl --location --request POST 'http://localhost:8080/togo/add' \
--header 'Content-Type: application/json' \
--data-raw '{
    "task":"example",
    "userid":1
}'

to run api test  execute this command 
"go test -timeout 30s -run ^TestAddtogoTask$ github.com/manabie-com/togo/test"
"go test -timeout 30s -run ^TestAddtogoTaskWithIncorrectParameter$ github.com/manabie-com/togo/test"

to run integration test execute this command 
- go test -timeout 30s -run ^TestAddtogo$ github.com/manabie-com/togo/handlers
- go test -timeout 30s -run ^TestAddtogoLimitTask$ github.com/manabie-com/togo/handlers
- go test -timeout 30s -run ^TestGetUserById$ github.com/manabie-com/togo/handlers
- go test -timeout 30s -run ^TestCreateUser$ github.com/manabie-com/togo/handlers

### technologies

- SQLite3
- Gin
- Gorm
- httpTest
- testing


### What do you love about your solution?
- I don't have many experience with go lang so I just use factory pattern to handle this task  , I think its good to maintenance and extension 
- I use some popular technologies so Its good to follow up and understand code.

### What else do you want us to know about however you do not have enough time to complete?
- I just want to say , I'm very interested with Go lang.