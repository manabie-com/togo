# Account:
## POST /login
curl http://localhost:8000/login
### Parameters
    username (required)
    password (required)
    email (optional)
    name (optional)
### reponse 
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjcyYzg5OWE2LWI1MGEtNDJhMy04OWU3LTQ4YzNlMzU2YzYxOCIsImFjY291bnRfaWQiOjEsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTY0MzE4MzAwNH0.SQnLADUGCK8sfTaJGZq4vn5ol5YLfDsH_1M5HgeaTX4"
}

## POST /account/create
curl http://localhost:8000/account/create
### Parameters
    username (required)
    password (required)
### response
{
    "isCreated": true
}

## POST /logout
curl http://localhost:8000/logout
### Parameters
    token (required)
### response
{
    "isLoggedOut": true
}

## GET /account/show (work in progress)
## PUT /account/update (work in progress)

# Todo:
## POST todo/create
curl http://localhost:8000/todo/create
### Parameters
    title (required)
    description (optional)
    status (optional)
    token (required)
### response
{
    "success": true,
    "todoId": 1
}

## GET todo/get/:id
curl http://localhost:8000/todo/get/1
### response
{
    "Status": 0,
    "description": "my first todo",
    "title": "todo1"
}

## PUT todo/update/:id (work in progress)
## DELETE todo/delete/:id (work in progress)