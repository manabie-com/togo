![CI](https://cellphones.com.vn/sforum/wp-content/uploads/2021/10/mn.jpg)

## Introduction
A chanllege from manabie, implement one single API which accepts a todo task and records it.
There is a maximum limit of N tasks per user that can be added per day

## Usage
**Requirements**
* Node-v14.x
* Docker-v3

**How to run app**
 * Run docker-compose to start the app
  ```sh
  $ docker-compose up
  ```
* A sample “curl” command to call API

  ```sh
  `curl` --location --request POST 'http://localhost:3000/task' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2MWJlZjRlMzUzZTkyZDUxYjNmYzE0MWQiLCJpYXQiOjE2Mzk5MDQ0ODN9.1TnhkCd5HfGeOwOJqkUUNArfM1UpU5IYqfMoaRUX_jk' \
  --header 'Content-Type: application/json' \
  --data-raw '{
      "title": "Task1",
      "description": "Proin eget tortor risus. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque in ipsum id orci porta dapibus."
  }'
  ```

* Run integration test and unit test:
  ```sh
    $ npm test
  ```
## Source code structure

```
|-document
|-src/
	|-controllers/
  	|-db/
	|-middlewares/
	|-models/
	|-routers/
		|-validators/
	|-services/
|-test/
	|-integration_test/
	|-unit_test/
		|-user/
		|-task/
|-index.js	
		
```

### The application is divided into distribution packages:
* controllers: it contains `user.js`, `task.js`  pakages, this pakages provides functions to handle requests and return repsonses.
* db: It contains `db.js` connect to database.
* middlewares: it contains `auth.js` is a next function, used to check logged from Authorization Header.
* models: it contains `user.js`, `task.js`, `taskLimit.js` are the schemas to store when working with the database.
* routers: it contains `task.js`, `user.js` to route task and user methods, `index.js` to combine them to one router express
* services: provides function to handle data with database 
* test: it contains integration test and unit test
* document: it contains Postman requests collection.
* `index.js`: it runs server 

## Future
* Use logger to write log requests.
* Add more apis to clear app.
* Use swagger to write doc apis.
* Use pattern design,  dependencies injection.
* Write more case for test.
