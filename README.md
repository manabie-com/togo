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

### Development Instructions

- Download and install [Visual Studio](https://visualstudio.microsoft.com/vs/community/) or [Visual Studio Code](https://code.visualstudio.com/download) with [C# extension](https://marketplace.visualstudio.com/items?itemName=ms-dotnettools.csharp) and [.NET 5 SDK](https://dotnet.microsoft.com/en-us/download/dotnet/5.0). Visual Studio is recommended as it is the most powerful IDE with powerful debugging features and better development experience when working with .NET technologies.
- Clone the code into a local folder.
- Open the solution file (*.SLN) with Visual Studio then click Run/Start Debugging (F5) or Run/Start without Debugging (Ctrl+F5) to run the app locally.
- If you are using a text editor only like Visual Studio Code, you need to run the app in the terminal or command prompt by navigating to the project folder and start the app by command `dotnet run --project ToDoBackend`.
- You can test the API with the following commands

```sh
# Get all tasks for user firstuser. There are currently two users initialized when app starts, firstuser and seconduser
curl -L -X GET 'http://localhost:5050/tasks?userId=firstuser'
```

```sh
# Create a new task

```

```sh
# Update a task status: 1 for ACTIVE, 2 for COMPLETED. Replace the <taskId> with the ID you get from the server.
curl -L -X PATCH 'http://localhost:5050/tasks/<taskId>' -H 'Content-Type: application/json' --data-raw '{ "status": 2 }'
```

```sh
# Delete a task. Replace <taskId> with the ID you get from the server.
curl -L -X DELETE 'http://localhost:5050/tasks/<taskId>'
```

```sh
# Create a new task
curl -L -X POST 'http://localhost:5050/tasks' -H 'Content-Type: application/json' --data-raw '{
    "content": "Something to do today",
    "status": 1,
    "userId": "firstuser"
}'
```
