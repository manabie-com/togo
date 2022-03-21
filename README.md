## Requirements
- Node v.16 or higher

## How to run the code locally

- Clone the project

```bash
  git clone <link>
```

- Go to the project directory

```bash
  cd togo
```

- Install dependencies

```bash
  npm install
```

- Start the server

```bash
  npm run dev or npm run start
```

## Environment Variables

To run the application, you will need to add the following environment variables to your .env file. The env.txt in the repository includes the working environment variables

`PORT=24`

`NODE_ENV=`

`MONGO_URI=`

`JWT_SECRET=`

## Sample "curl" commands to call the API
- Register a user (Premium users has 5 max limit of tasks that can be added daily, non-premium users has 3). The first example is for premium users, second is for non-premium.
```curl
curl -X POST http://localhost:24/api/register \
   -H 'Content-Type: application/json' \
   -d '{"name":"example", "email":"e@m.com","password":"1234", "isPremium":true}'
```
```curl
curl -X POST http://localhost:24/api/register \
   -H 'Content-Type: application/json' \
   -d '{"name":"example", "email":"e@m.com","password":"1234"}'
```
- Login to get the bearer token
```curl
curl -X POST http://localhost:24/api/login \
   -H 'Content-Type: application/json' \
   -d '{"email":"e@m.com","password":"1234"}'
```
- Add a task request (Replace the {token} with the token without the "" marks)
```curl
curl -X POST http://localhost:24/api \
   -H 'Content-Type: application/json' \
   -H 'Authorization: Bearer {token}' \
   -d '{"description":"test sample"}'
```
- Get all the request of the user
```curl
curl -X GET http://localhost:24/api \
   -H 'Content-Type: application/json' \
   -H 'Authorization: Bearer {token}'
```
## DB Schema
```javascript
User = {
  id,
  name,
  email,
  password,
  isPremium,
  task[
    description, 
    date
  ],
  dateCreated
}
```
## Running Tests
To run tests, run the following command (Replace the assigned token to a new one from the logged in user in the test suite)

```bash
  npm test <name of the test file>
```
## What I like about my solution?
- I made the application structure and design closest to a real-world app
- I was able to maintain a good way of keeping the separation of concerns

## What else could have been improved?
- I wish I would have done the tests in a TDD manner. I hope through often exposure to the routine, I would make myself comfortable doing it.
 
