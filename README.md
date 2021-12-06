![CI](https://github.com/AnhPhan49/togo-1/actions/workflows/node.js.yml/badge.svg)
# Manabie

## Introduction
A challenge from manabie, build togo API assignment.

This project has been integrated for continuous testing with github CI. Check [status](https://github.com/AnhPhan49/togo-1/actions) here.

## Usage
**Requirements**
* Node-v14.x
* Docker-v3

**How to run app**
 * Run docker-compose to start the app
	```	
	$ docker-compose up
	```

* Run integration test and unit test:
	* Install pakages: 	
	
	```	
	$ npm test
	```
	* Run test:
	```	
	$ npm test
	```

## Source code structure

```
|-docs
|-src/
	|-config/
	|-controllers/
	|-middlewares/
	|-models/
	|-routes/
		|-validators/
	|-services/
|-test/
	|-integration_test/
	|-unit_test/
		|-auth/
		|-test/
|-index.js	
		
```

The application is divided into distribution packages:
* config: It contains `databse.config.js` connect to database.

* controllers: it contains `auth.controller.js`, `note.controller.js`  pakages, this pakages provides functions to handle requests and return repsonses.
* middlewares: it contains `checkLogin.middleware.js` is a next function, used to check logged.
* models: it contains `user.model.js`, `note.model.js` are the schemas to store when working with the database.
* services: provides function to handle data with database 
* test: it contains integration test and unit test
* docs: it contains Postman requests collection.
* `index.js`: it runs server 

## Future
* Use logger to write log requests.
* Add more apis to clear app.
* Use swagger to write doc apis.
* Use pattern design,  dependencies injection.
* Write more case for test.
