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


### Installation
- Run code locally
  - docker-compose up -d
  - access link: http://localhost:5000/swagger/index.html
- Run unit test 
  - Download Dotnet SDK and install: https://dotnet.microsoft.com/en-us/download
  - dotnet test Todo.Application.Test
  - dotnet test Todo.Api.Test
- My solution is:
    - Using clean code structure: https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
      - Todo.Api : Main project which define API.
      - Todo.Application : Define some services for handling logic
      - Todo.Domain : Define Entity Class
      - Todo.Infracture: Create DbContext and communicate with Todo.Application via IApplicationDbContext(interface)
      - Todo.Application.Test : Testing services(unit test)
      - Todo.Api.Test : Testing controllers(intergation test)
    - Entity Framework code first to build and manage versions.
    - Solid principle
    - Using docker for setup local environment
    - Database design suit requirement
- Other ideals:
   - Get UserId from JWT Token to assign CreatedBy field automatically
   - Handle Exception
   - Define Error Message/ Status response
   - Custom validation in Dto class
   - Handling timezone which is depending on each user.