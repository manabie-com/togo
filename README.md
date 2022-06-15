### Run the project locally

- Install the dependencies: `npm i`
- Set up database with docker: `docker-compose -f "docker/docker-compose.local.yml" up -d --build'`
- Run the project: `npm start`
- Project starts on localhost:3000 by default

### Sample curl command

`curl --location --request POST 'http://localhost:3000/task' \ --header 'Content-Type: application/x-www-form-urlencoded' \ --data-urlencode 'userId=1' \ --data-urlencode 'title=lam bai tap' \ --data-urlencode 'desc=lam rat rat nhieu bai tap'`

### Run unit test

`npm test`

### Run integration test

`npm test:e2e`

### What do you love about your solution?

I think my solution make flexible and scalable. It can be easily expanded in the future.
