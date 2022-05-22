# Manabie

Project using node.js to do the test

---

## Requirements
For development, you will need Node.js and a node global package, Npm, and MongoDB installed in your environment.

### 1. Mongo DB
  If MongoDB is not installed, just run docker-compose to set up your mongo DB in the local environment for your own.

    $ docker-compose up -d
### 2. Nodejs + npm

  Just go on [official Node.js website](https://nodejs.org/) and download the installer.
Also, be sure to have `git` available in your PATH, `npm` might need it (You can find git [here](https://git-scm.com/).

## Running the project
### 1. Open terminal/command prompt in folder togo
### 2. Install the necessary packages

    $ npm install

* ***Note: Just run this command when running the program for the first time*** 

### 3. Assign environment variables
  If using your own MongoDB, you should change **DB_CONNECTION** variable to your connection URI in **.env** file
### 4. Start project

    $ npm start

## A sample “curl” command to call my API

```bash
curl --location --request POST 'http://localhost:3000/api/v1/auth/register' --header 'Content-Type: application/x-www-form-urlencoded' --data-urlencode 'name=Manabie'\''s user' --data-urlencode 'email=manabie1@manabie.com' --data-urlencode 'password=manabie1' --data-urlencode 'maxTask=10'
```
  For easy tracking of all project's API, you should visit [here](http://localhost:3000/swagger) to use swagger to test with account and password below:
  ```
  Url: http://localhost:3000/swagger
  Account: admin
  Password: admin
  ```

## Running the test
### 1. Open terminal/command prompt in folder togo
### 2. Install the necessary packages

    $ npm install

* ***Note: Just run this command when running the program for the first time*** 

### 3. Assign environment variables
  If using your own MongoDB, you should change **DB_CONNECTION** variable to your connection URI in **.env.test** file
### 4. Test project

    $ npm test

### 5. After run the test
  You can see test coverage by go to **togo/coverage/lcov-report** folder and open **index.html** file to see it.
## Future
 * Apply TDD in future function
 * Using Go to do this test instead of node.js