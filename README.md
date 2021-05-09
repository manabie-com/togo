### Overview
This is a simple backend for a good old todo service, right now this service can handle signup/login/list/create simple tasks.  

### Requirements

- golang 1.14 or above
- [Docker](https://docs.docker.com/engine/install/) and [docker-compose](https://docs.docker.com/compose/install/) has been installed. 

### Run

#### 1. Manual

```
go mod download && go run main.go
```

- The above script will create a backend server that uses sqlite as database

#### 2. Docker

Use docker-compose to build Docker image and run the server 

```
docker-compose
```

### API

`Host: http://localhost:5050`

#### 1. Signup

- Sign up new user
- After sign up successfully, a token will be returned.
- User can add this token to header at "Authorization" key to continue working with listing/adding new tasks

```
POST /signup
BODY
{
    "id": "firstUser"
    "password": "example"
}
Response: 200
{
    "data": "jwt token"
}
```

#### 2. Login

- After login, a token will be returned.
- User can add this token to header at "Authorization" key to continue working with listing/adding new tasks

```
POST /signup
BODY
{
    "id": "firstUser",
    "password": "example"
}
Response: 200 
{
    "data": "jwt token"
} 
```

#### 3. Add Task

- In order to use this API, user must login and has returned token
- Add this token to Header at "Authorization" key

```
POST /tasks
HEADER 'Authorization': 'JWT Token generated from login/signup'
BODY
{
    "user_id": "firstUser",
    "content": "testContent"
}
Response: 200
{
    "data": {
        "id": "123",
        "user_id": "a123",
        "content": "task 1"
    }
}
```
- Each user only has a limit `max_todo` per date. If number of tasks reaches the max_todo it will throw the following message:

```
Response: 400
{
    "data": "max limit tasks reached"
}
```

#### 4. List Tasks

- List all tasks belong to a user and their created date.
- In order to use this API, user must login and has returned token
- Add this token to Header at "Authorization" key

```
GET /tasks
HEADER 'Authorization': 'JWT Token generated from login/signup'
Response: 200
{
    "data": [
        {
            "id": "123",
            "user_id": "a123",
            "content": "task 1"
        },
        {
            "id": "124",
            "user_id": "a124",
            "content": "task 2"
        }
    ]
}
```

### DB Schema
```sql
-- users definition

CREATE TABLE users (
	id TEXT NOT NULL,
	password TEXT NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (id, password, max_todo) VALUES('firstUser', 'example', 5);

-- tasks definition

CREATE TABLE tasks (
	id TEXT NOT NULL,
	content TEXT NOT NULL,
	user_id TEXT NOT NULL,
    created_date TEXT NOT NULL,
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);
```

### Sequence diagram
![auth and create tasks request](https://github.com/manabie-com/togo/blob/master/docs/sequence.svg)
