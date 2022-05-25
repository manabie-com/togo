## Contents
- [Usage](#usage)
   - [Run normally](#run-normally)
   - [Run through docker](#run-through-docker)
- [Call API](#call-api)
   - [With curl command](#curl-command)
   - [With swagger](#swagger)
- [Tests locally](#tests-locally)
   - [Check code format](#check-code-format)
   - [Run unit test only](#run-unit-test-only)
   - [Run integration test only](#run-integration-test-only)
   - [Run all test](#run-all-test)
   - [Run coverage test](#run-coverage-test)
- [What do you love about your solution?](#what-do-you-love-about-your-solution)
- [What else do you want us to know about however you do not have enough time to complete?](#what-else-do-you-want-us-to-know-about-however-you-do-not-have-enough-time-to-complete)

## Usage

- You must create file `.env` for the `first time` (in this case, I will put file.env into git).
- Description variable of env file:

| Environment    | Is_Required | Description                                                                                           |
| -------------- | ----------- | ----------------------------------------------------------------------------------------------------- |
| DATABASE_URL   | `true`      | MongoDB connection string with the standard URI connection scheme has the form: mongodb://[username:password@]host1[:port1][,...hostN[:portN]][/[defaultauthdb] |
| PORT           | `true`      | Project's port                                                                                        |


- There is several way to run the project:

   ### Run normally:
   - You must use `file .env` and `mongoose` are running local with port `27017` and `Nodejs must be > 16.x`:
      ```env
      DATABASE_URL=mongodb://localhost:27017
      PORT=3000
      ```
   - Then run `cmd`:

      ```bash
      yarn
      yarn start
      ```
   - If you don't have `yarn` or you prefer to run with `npm`:
      ```bash
      npm i
      npm start
      ```

   ### Run through docker:
   - You must use `file .env` 
      ```env
      DATABASE_URL=mongodb://mongo:27017
      PORT=3000
      ```
   - Then run `cmd`

      ```bash
      sudo docker-compose up
      ```


## Call API
- I have several choice for you:

   ### `curl` command: 
   - Check health of server: 
      ```bash
      curl http://localhost:3000/api/ping  
      ```
   - Create new user : 
      ```bash
      curl -d '{"username":"first_user"}' -H 'Content-Type: application/json' http://localhost:3000/api/user/
      ```
   - Get user's profile : 
      ```bash
      curl http://localhost:3000/api/user/first_user
      ```
   
   - After we get user's profile we get userID and create new task: 
      ```bash
      curl -d '{"name":"Task name"}' -H 'Content-Type: application/json' http://localhost:3000/api/task/{userID}
      ```

   - After we get user's profile we get userID and get list task: 
      ```bash
      curl http://localhost:3000/api/task/{userID}
      ```

   ### Swagger:
   - Go to http://localhost:3000/api/docs

## Tests locally
- Before doing any test please install the package first with command:
   ```bash
   yarn
   ```
- Or with `npm`:
   ```bash
   npm i
   ```
   
- Then I will give you some choice again:

   ### Check code format:
   ```bash
   yarn lint
   ```
   - Or with `npm`:
   ```bash
   npm run lint
   ```
   
   ### Run unit test only:
    ```bash
   yarn test:unit
   ```
   - Or with `npm`:
   ```bash
   npm run test:unit
   ```

   ### Run integration test only:
   ```bash
   yarn test:integration
   ```
   - Or with `npm`:
   ```bash
   npm run test:integration
   ```

   ### Run all test:
   ```bash
   yarn test
   ```
   - Or with `npm`:
   ```bash
   npm run test
   ```

   ### Run coverage test:
   ```bash
   yarn test:coverage
   ```
   - Or with `npm`:
   ```bash
   npm run test:coverage
   ```

## What do you love about your solution?
- About design pattern:
   - I chose the MVC design pattern because it is suitable for this type of application and I found out that it is easy to maintain if there is a massive expanse of the size of the project.
      + I make the name file more specific, for example, the user in the controller is user.controller, because I discovered that if there are a lot of namesake, it is easy to find a file.
      + About test name file: Like above, it will be create the same folder name and the place they are to more maintainable and track down the test is belonging to.

- About the typescript:
   - The reason I prefer TypeScript more than pure JavaScript is that it is more strict and has tighter integration of the types of input and output. They can provide us with a predictable value that we will use and provide us with the ability to avoid the bug as much as possible. If there is a newcomer, they will make it easier to continue the work pending or maintain it. Besides, using typescript can help me decrease the number of documents I have to write while developing the system.

- I choose mongo it's provide:
   - User experience: It is an easy-to-use, high-performance, and versatile tool. Faster than SQL NoSQL, demoralized allows us to get all the information we need about a particular item with conditions without the need for JOINs involved or complex SQL queries.
   - The ability to scale

## What else do you want us to know about however you do not have enough time to complete?
- There is a lot of things a want to mess around:
   - About the user :
      + We can have authentication, like using bcrypt to hash the password then save to the database, and using jwt with access token with refresh token to send token when they login
      + With the refresh token, it will send us the access token when we log in, and while using it, if the access token is out of time, we can use the refresh token to create a new access token and we can reset the refresh token as well. Normally, I'd use Redis to keep track of this refresh token and get the refresh token immediately and no need to go to the DB to query.
      + With the access token, usually for more secure I will store the access Token in the DB with the createDate , it can use until that date.
      + Beside the user DB will have a active field so when using the API if the active field false, that account can not use any thing in our system

   - System: 
      + Redis to store the refresh token, all information that all users can get.
      + The logger will be stored in order by date in the file.
      + Currently, I have it configured so that anyone can use our system to call API. In the real world, we only accept a few websites to call our API.
      + For more security, In the case that our server can be hacked, we can build our project to bundle the project, making it harder to read code and edit it. In some cases, we can also use Webpack to rename all files, minimize files, and translate to another type of code. It also helps us to speed up a little bit.