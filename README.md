# Project Name

AKARU Todo API

## Technologies
* Nodejs
* Express
* Mongodb
* Mongodb Atlas

## Installation

`cd backend` to go into the project root

### npm to install the backend's npm dependencies
`npm install`

## Running Locally

`nodemon server` to start the development server

## Curl Commands

### Register 
`curl -d "username=admin&password=password12345&fullname=Admin&confirmpassword=password12345" http://localhost:5000/api/user/register` supply username, password, confirmpassword and fullname.
### Sample JSON 
```javascript
{
    "data": {
        "fullname": "Akaru Admin",
        "username": "akaruadmin",
        "password": "$2a$10$I.4.2LcYM9N3bDRc1nDhDuvOz3QoHnW5ozOLp5lo75GTfjx59guQi",
        "limit": 0,
        "_id": "62a2b1a2a0295faf6ca23b1a",
        "created_at": "2022-06-10T02:51:14.626Z",
        "updated_at": "2022-06-10T02:51:14.626Z",
        "__v": 0
    },
    "status": "successfuly registered"
}
```

### Login
`curl -d "username=admin&password=password12345" http://localhost:5000/api/user/login` token key will use for adding a task as HEADER
### Sample JSON
```javascript
{
    "status": "success",
    "data": {
        "_id": "62a2b1a2a0295faf6ca23b1a",
        "fullname": "Akaru Admin",
        "username": "akaruadmin",
        "password": "$2a$10$I.4.2LcYM9N3bDRc1nDhDuvOz3QoHnW5ozOLp5lo75GTfjx59guQi",
        "limit": 0,
        "created_at": "2022-06-10T02:51:14.626Z",
        "updated_at": "2022-06-10T02:51:14.626Z",
        "__v": 0,
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFrYXJ1YWRtaW4iLCJwYXNzd29yZCI6IlBhc3N3b3JkMTIzNDUiLCJpYXQiOjE2NTQ4Mjk1NTUsImV4cCI6MTY1NDgzMzE1NX0.jV0id-DgeINGv18M0in601tn-SI7dnpEoE7Faphjldg"
    }
}
```

### Adding a task with token
`curl -d "task=coding and sleeping" http://localhost:5000/api/task/addTask -H "x-auth-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiJwYXNzd29yZDEyMzQ1IiwiaWF0IjoxNjU0OTIxMDQxLCJleHAiOjE2NTQ5MjQ2NDF9.lhI3ePCzm8Ixhln8LIXoB7Qjp_j2Nd2y4cP38Oiv63E"`
### Sample JSON
```javascript
{   
    "status":"success",
    "data": {
        "task":"play",
        "userName":"akaruadmin",
        "_id":"62a2b2c0a0295faf6ca23b1e",
        "created_at":"2022-06-10T02:56:00.028Z",
        "updated_at":"2022-06-10T02:56:00.028Z","__v":0       
    }
}
```

### Adding a task without token
`curl -d "task=play" http://localhost:5000/api/task/addTask`
### Sample JSON
```javascript
{   
    "status":401,
    "message":"You are not logged in"
}
```

### Getting all the task
`curl http://localhost:5000/api/task/getTask`
### Sample JSON
```javascript
{
    "status": "success",
    "data": [
        {
            "_id": "62a426713910f3ed137b5ed4",
            "task": "coding and sleeping",
            "userName": "admin",
            "created_at": "2022-06-11T05:21:53.054Z",
            "updated_at": "2022-06-11T05:21:53.054Z",
            "__v": 0
        },
        {
            "_id": "62a426953910f3ed137b5ed8",
            "task": "eating",
            "userName": "admin",
            "created_at": "2022-06-11T05:22:29.229Z",
            "updated_at": "2022-06-11T05:22:29.229Z",
            "__v": 0
        }
    ]
}
```

## How to run your unit tests locally?

`cd backend`

`npm run test`

## What do you love about your solution?

<p> What I love about my journey developing the backend challenge is I have learned a lot of things. This is the first time I used the curl commands and also the cron job for scheduling the reset limit per day. <p>

## What else do you want us to know about however you do not have enough time to complete?

* A task can be completed.
* A task can be updated.
* A task can be deleted.
* A task can have subtasks.
* Individual user can get task

## Note

### Add ENV file
<details>
<p> ATLAS_URI = mongodb+srv://{username}:{password}@cluster0.3nk8zqt.mongodb.net/akaru?retryWrites=true&w=majority. </p>

<p> I will attach the username and password of my database in the email. <p>


