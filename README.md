##### Run the project locally

- Install the dependencies: `npm i`
- Run the project: `npm start`
- Project starts on localhost:5000 by default

##### Sample curl command

`curl --location --request POST 'http://localhost:5000/task' --header 'Content-Type: application/json' --data-raw '{\"user_name\": \"user1\",\"task_name\": \"task1\"}'`

##### Run unit test

`npm test`

#### Run integration test

- Ensure the project is running locally
- Add the endpoint URL in src\test\integration.test.js (url variable)
- Run: `npm run test:integ`

#### Other notes

- I love that I made the code testable and that I used the hexagonal architecture. For example, I used dependency injection in the repository. This makes the unit testing easier.
