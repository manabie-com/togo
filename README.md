# Node.js project with Express, Sequelize, Jest and Postgres

by [Buudld]

### Directory structure

```shell
src
  ├── app.js         app entry point
  ├── /routes        controller layer: api routes
  ├── /config        config settings
  ├── /services      service layer: business logic
  ├── /models        data access layer: database models	
test       
  ├── /unit          unit test suites
  ├── /integration   test api routes
 ```

### Installation and execution

1. Run yarn to install dependencies: `yarn`;
1. Config database credentials inside `/src/config/database.js`;
1. Create the table manually in order to start the server from the file `./database.sql`;
1. Run `yarn dev` to start the server.
1. Run `yarn test` to start the server.


### API 

- Sign-up: `http://localhost:3000/api/auth/signup`
- Login: `http://localhost:3000/api/auth/signin`
- Users: `http://localhost:3000/api/users`
- Task: `http://localhost:3000/api/task`

### Curl Sign Up
curl --location --request POST 'http://localhost:3000/api/auth/signup' \
--header 'Content-Type: application/json' \
--header 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MzUsImlhdCI6MTY1NzM3Mzk3NywiZXhwIjoxNjU3Mzc3NTc3fQ.QzE-m84YuUiMdQ_FH9zarsiYqUSX-zH3-__1SMJ_i08' \
--data-raw '{
    "name": "Đình Bửu",
    "email": "dinhbuu1208@gmail.com",
    "password": "admintodo",
    "role": "Admin"
}'

### Curl Sign In 
curl --location --request POST 'http://localhost:3000/api/auth/signin' \
--header 'Content-Type: application/json' \
--header 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MzcsImlhdCI6MTY1NzM3NDAwNywiZXhwIjoxNjU3Mzc3NjA3fQ.Wed21w5ETRt9RHW4UNkyfDk-nVNJ-ATIBi2u2pHT4IU' \
--data-raw '{
    "email": "dinhbuu1208@gmail.com",
    "password": "admintodo"
}'

### Curl Create Task
curl --location --request POST 'http://localhost:3000/api/task' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Test Todo Task",
    "description": "Test Todo Task",
    "text": "Test Todo Task"
}'

### What do you love about your solution?
1. This backend is written in Nodejs:  I can write Javascript code outside the browser to create server-side web applications that are non-blocking, lightweight, fast, robust and scalable.
2. Separation of concern principle is applied: Each component has been given a particular role. The role of the components is mutually exclusive. This makes the project easy to be unit tested.
3. Feature encapsulation is adopted: The files or components that are related to a particular feature have been grouped unless those components are required in multiple features. This enhances the ability to share code across projects.
4. Centralised Error handling is done: We have created a framework where all the errors are handled centrally. This reduces the ambiguity in the development when the project grows larger.
5. Async execution is adopted: We have used async/await for the promises and made sure to use the non-blocking version of all the functions with few exceptions.
6. Unit test is favored: The tests have been written to test the functions and routes without the need of the database server. Integration tests has also been done but the unit test is favored.

### What else do you want us to know about however you do not have enough time to complete?
1. Writting skill unit test and intergation test not perfect.
