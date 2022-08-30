### Submission

#### How to run code locally?

  1. Dependency: Need to have ruby 3.0.1, and Postgres 12 installed
  2. Run bundle install: `bundle install`
  3. Prepare config file: `cp config/database.yml.example config/database.yml`
  4. Setup database: `rails db:setup`
  5. Start the server: `rails server`

#### A sample "curl" command

  I have seeded 2 users in the database, defined in the `seeds.rb` file, and was inserted into the database when running `rails db:setup` command. 
  Note that one user has attributes: `id: 1, daily_task_limit: 3` and the other user has attributes `id: 2, daily_task_limit: 5`

  A sample `curl` command to call the API: 

    curl 'http://localhost:3000/tasks/create' \
      -H "Content-Type: application/json" \
      -d '{ "user_id": "1", "name": "Task 1" }'

  If task is created succesfully, we would get a response as follow:
  
  `{"id":1,"name":"Task 1","user_id":1,"created_at":"2022-04-29T14:40:22.669Z","updated_at":"2022-04-29T14:40:22.669Z"}`

  If task failed to be created due to user reaching maximum daily task, we would get a following response:

  `{"error":"This user has reached its maximum daily limit of adding tasks! (n tasks/day)"}` (n can be 3 or 5)

#### To run test locally:

  `rails test`

#### What do I love about my solution:

   - My solution uses the MVC framework (without the view since this is just an API), it helps to separate handling of data with the processing of request parameters and sending a response back to the client
   - Seed the user in the database instead of having to created them with an api or manually insertion to database
   - It is easier to write unit test and integration test as controllers are separated from models



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
