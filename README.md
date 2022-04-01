### FAQ

- How to run your code locally?

For running the server, use a simple command: ```$ make run```
- A sample “curl” command to call your API
```sh
curl --location \
  --request POST 'localhost:9000/to-do' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "user_id": "user-id",
  "entry": [
    {
      "content": "drink water"
    }
  ]
}'
```
- How to run your unit tests locally?

To run the unit tests locally, use the command: ```$ make test```
- What do you love about your solution? 

The solution for adding the to do list was designed to deal with high rate of update rate. 
In that, we try to commit the length of the data piece in order to control the limit value, before storing all the remaining data.

The project layout was inspired from the project [Standard Go Project Layout](https://github.com/golang-standards/project-layout). 
The project architecture was inspired from the Clean Architecture. 
I also added some innovations in order to simplify the model for smaller project (pretty common in today's context, microservices)

The project was built with gRPC, protobuf, MySQL and Redis, 
which are the main technologies that I have been working for over 3 years. 

- What else do you want us to know about however you do not have enough time to complete?

I have tried my best implementing the requirements that you offered. 
Some unit tests and functional tests have been added to keep the project work. 
I think it is enough to fulfil the requirements with this limitation of time.

### Requirements checklist
- [x] Implement one single API which accepts a todo task and records it
  - There is a maximum **limit of N tasks per user** that can be added **per day**.
  - Different users can have **different** maximum daily limit.
- [x] Write integration (functional) tests
- [x] Write unit tests
- [x] Choose a suitable architecture to make your code simple, organizable, and maintainable
- [x] Write a concise README
  - How to run your code locally?
  - A sample “curl” command to call your API
  - How to run your unit tests locally?
  - What do you love about your solution?
  - What else do you want us to know about however you do not have enough time to complete?
