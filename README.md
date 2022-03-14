# NodeJS API
This is a nodejs API (MVC)

## Installation

Install the packages

```bash
npm install
```

Run migration and seeder
```bash
cd db
```
```bash
npx sequelize-cli db:migrate
```
```bash
npx sequelize-cli db:seed:all
```


Run the program - dev env

```bash
npm run dev
```

## Run API

You can import link https://www.getpostman.com/collections/8c5a8007b02e6d9f147e into Postman
* Login API http://localhost:8000/login
    * Method: POST
    * Header: 
        ```json
        {
          "Content-Type": "application/json"
        }
        ```
    * Body content
        ```json
        {
          "email": "admin@gmail.com",
          "password": "admin123"
        }
        ```
* Tasks API http://localhost:8000/tasks
  * Method: POST
  * Header: 
    ```json
    {
      "Content-Type": "application/json",
      "authorization": "Bearer {TOKEN}"
    }
    ```
  * Body content
      ```json
      {
        "error": false,
        "message": "",
        "data": [
            {
                "id": 2,
                "name": "Task 1",
                "description": "Write API",
                "status": "in-process",
                "estimated_time": 1647047902,
                "due_date": 1647047902,
                "user_id": 2,
                "created_at": 1646961402,
                "updated_at": 1646961402
            }
        ]
      }
      ```
* Users API http://localhost:8000/users
  * Method: POST
  * Header: 
    ```json
    {
      "Content-Type": "application/json",
      "authorization": "Bearer {TOKEN}"
    }
    ```
  * Body content
      ```json
      {
        "error": false,
        "message": "",
        "data": [
            {
                "id": 1,
                "name": "Admin",
                "email": "admin@gmail.com",
                "password": "$2b$10$N.XayDSVMF5FKYedKRofgOI7vxVwjqLz90A9V8ibyzsIpZV0hGePm",
                "status": "active",
                "role": "admin",
                "created_at": 1646873640,
                "updated_at": 1646873640
            }
        ]
      }
      ```
## Unit Test
Run the program

```bash
npm run test
```