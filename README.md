# Introduction
A challenge from Manabie. Build togo API assignment.
This project is automatically tested [ integation-test, unit-test ] on github actions. Please check the status [here](https://github.com/tienvnz98/togo/actions).

## Usage
**Requirements**
* Node-v16.x
* Docker-20.x.x

### How to run test localy
**Run integration test and unit test**
```
npm i
npm test
```

### How to run this app
Install node packages.
```
npm i
```

**First options: Run this app on the host localy (development mode).**

Start dependencie services in depend-services.yml file.
```
docker-compose -f depend-services.yml up -d
```

Start node http server on the host.
```
npm start
```

**Second options: Run this app inside container (production mode)**

Start docker-compose file.
```
docker-compose up -d
```

### How to request to this app by curl
All request detail inside [postman collection](./docs/TOGO-TEST.json)

```
curl --location --request GET 'http://localhost:9200/api/public/home'

It will response
{
    "success": true,
    "code":200,
    "data": {
        "message": "REST API VERSION 0.0.1.",
        "date":"2022-05-23T16:46:05.029Z"
    }
}
```

## Source code structure
```
    |-docs/
    |-src/
        |-controllers/
        |-middlewares/
        |-models/
        |-services/
        |-utils/
        |-app.js
        |-routes.js
        |-server.js
    |-test/
        |-integration-test/
        |-unit-test/
            |-auth/
            |-task/
    |-package.json
```

The application is divided into distribution packages:
* **controllers**: it contains pakages, this pakages provides functions to receive requests and return repsonses.
* **middlewares**: it contains is a next function, used to append some function in context or verify token in headers request.
* **models**: it contains the schemas to working with the database.
* **services**: provides function to handle data with database. 
* **test**: it contains integration test and unit test.
* **docs**: it contains Postman requests collection.

