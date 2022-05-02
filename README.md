### About this solution
- Run the code in locally
  - Please setup Redis cache
  - Add .net core 5.0 to the visual studio
- Sample “curl” command to call the API
  - Api for inserted
    - url : https://localhost:44388/userstask/insert
	- method : POST
	- body example : 
		{
			"isDeleted": false,
			"createdDate": "2022-05-01T12:42:18.100Z",
			"createdBy": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
			"modifiedDate": "2022-05-02T12:42:18.100Z",
			"modifiedBy": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
			"userId": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
			"taskName": "taskName-1",
			"description": "this is task 1",
			"taskDate": "2022-05-02T12:42:18.100Z"
		}
  - Able to using swagger for calling apis
- Run unit test : please go to test/ Test explorer (in visual studio) and chose "Run All Tests In View"
- About the solution
  - Using c# .net core 5.0 
  - Using json-textfile for repository
  - Using Redis for caching
  - Using CQRS pattern
- Somethings else
  - Able to using entity framework with Sql server instead json-textfile

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
