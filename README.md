# TOGO PROJECT

## Install requirement

- **Nodejs**: 12.16.0 (every version >= 12.\* is OK)
- **Docker**: 20.10.12 (required if setup database)
- **Docker-compose**: 1.25.0 (required if setup database)

## Development environments

- **LCL**: Local environment

- The configuration file with the corresponding environments is located in the directory `env`.

## Run docker and setup database

- **Step 1:** Go to folder docker:
  `cd docker`
- **Step 2:** Run docker compose:
  `sudo docker-compose up -d`
- **Step 3:** Check database is running by command:
  - **1** `sudo docker-compose exec mongodb /bin/sh`
  - **2** `mongo`
- **Step 4:** Get uri in connecting to: mongodb://127.0.0.1:27017/  and add to LCL.env
- **Note**: uri: take the part before the question mark

## Install project and run server & run unit tests

- **Step 1:** Clone git repository:
  `git clone https://github.com/TTKirito/togo.git`
- **Step 2:** Install packages:
  `npm install`
- **Step 3:** In LCL.env.example:

  - Copy file `LCL.env.example`
  - Rename to `LCL.env`

    Change configuration to connect database , mongo at local, ...

- **Step 4:** Running server
  - Normal run: `npm start`

- **Step 5:** Running unit tests
  - Normal run: `npm run test`


### Explore Rest APIs
- The app defines following CRUD APIs.

	`POST /users/signup`

  ```
  curl --location --request POST 'http://localhost:3000/api/users/signup' \
    --header 'Content-Type: application/json' \
    --data-raw '{
      "email": "thuanton983@gmail.com",
      "password": "password"
    }'

  ```

	`POST /users/signin`

  ```
  curl --location --request POST 'http://localhost:3000/api/users/signin' \
    --header 'Content-Type: application/json' \
    --data-raw '{
      "email": "thuanton985@gmail.com",
      "password": "password"
    }'

  ```


	`GET /users/currentuser`

  ```
  curl --location --request GET 'http://localhost:3000/api/users/currentuser' \
  --header 'Content-Type: application/json' 

  ```

	`POST /users/signout`

  ```
  curl --location --request POST 'http://localhost:3000/api/users/signout'
  ```

  `POST /api/tasks`

  ```
  curl --location --request POST 'http://localhost:3000/api/tasks' \
  --header 'Content-Type: application/json' \
  --header 'Cookie: express:sess=eyJqd3QiOiJleUpoYkdjaU9pSklVekkxTmlJc0luUjVjQ0k2SWtwWFZDSjkuZXlKcFpDSTZJall5WVRoak1qVmhZbVZqTUdZMFpUY3lPRGM1T0RNMk1pSXNJbVZ0WVdsc0lqb2lkR2gxWVc1MGIyNDVPRE5BWjIxaGFXd3VZMjl0SWl3aWFXRjBJam94TmpVMU1qTTBORGc0ZlEua2tlRnBoMTlTbUl4ampTR0hVZkdmQUJnS2JBTENTcktqX19sd2xkd3BGNCJ9' \
  --data-raw '[
    {
        "description": "test1",
        "title": "title1"
    }
  ]'

  ```
  `GET /api/tasks`

  ```
  curl --location --request GET 'http://localhost:3000/api/tasks' \
  --header 'Content-Type: application/json' \
  --header 'Cookie: express:sess=eyJqd3QiOiJleUpoYkdjaU9pSklVekkxTmlJc0luUjVjQ0k2SWtwWFZDSjkuZXlKcFpDSTZJall5WVRoa1pETmlaV0k1TURRMllqTmpZV1l4T0dWaFlpSXNJbVZ0WVdsc0lqb2lkR2gxWVc1MGIyNDVPRFZBWjIxaGFXd3VZMjl0SWl3aWFXRjBJam94TmpVMU1qTTBOVFl4ZlEuX00yREFJSmpxTWhBSWdhcWd2OG1pSFpkWThvVnhtQmt2UHhQZENzNWpidyJ9' 
  ```


You can test them using postman or any other rest client.


# What else do you want us to know about however you do not have enough time to complete?
- handler and define error for return 
- add document (ex: swagger)
- deploy by k8s 