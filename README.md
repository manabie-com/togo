# Project Name

AKARU Todo API

## Installation

`cd backend` to go into the project root

#npm to install the backend's npm dependencies
`npm install`

## Running Locally

`nodemon server` to start the development server

## Curl Commands

### Register 
`curl -d "username=sampleuser&password=samplepassword&fullname=samplefullname&confirmpassword=samplepassword" http://localhost:5000/api/user/register` supply username, password, confirmpassword and fullname.
### Sample JSON 
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

### Login
`curl -d "username=pergent100&password=password12345" http://localhost:5000/api/user/login` token key will use for adding a task as HEADER
### Sample JSON
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

### Adding a task with token
`curl -d "task=play" http://localhost:5000/api/task/addTask -H "x-auth-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFrYXJ1YWRtaW4iLCJwYXNzd29yZCI6IlBhc3N3b3JkMTIzNDUiLCJpYXQiOjE2NTQ4Mjk1NTUsImV4cCI6MTY1NDgzMzE1NX0.jV0id-DgeINGv18M0in601tn-SI7dnpEoE7Faphjldg"`
### Sample JSON
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

### Adding a task without token
`curl -d "task=play" http://localhost:5000/api/task/addTask`
### Sample JSON
{   
    "status":401,
    "message":"You are not logged in"
}

## How to run your unit tests locally?

TODO: Write here

## What do you love about your solution?

TODO: Write here

## What else do you want us to know about however you do not have enough time to complete?

TODO: Write here


