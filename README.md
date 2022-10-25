### Notes
- This is a simple backend for a todo service, right now this service can handle login/list/create simple tasks.

### Requirements

- Implement one single API which accepts a todo task and records it
  - There is a maximum **limit of N tasks per user** that can be added **per day**.
  - Different users can have **different** maximum daily limit.
- Write integration (functional) tests
- Write unit tests
- Choose a suitable architecture to make your code simple, organizable, and maintainable
- Using Docker to run locally
  - Using Docker for database (if used) is mandatory.
- Write a concise README
  - How to run your code locally?
  - A sample “curl” command to call your API
  - How to run your unit tests locally?
  - What do you love about your solution?
  - What else do you want us to know about however you do not have enough time to complete?
---
#### DB Schema
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

---
## How to run project
### Step 1:
```shell
make dev.up
```
### Step 2:
```shell
make run
```
---
## How to run unit test
```shell
make test.unit
```
---
## How to run integration test
```shell
make test.integration
```
## A sample “curl” command to call API
- Need login to get token.
```
curl --location --request POST 'localhost:8001/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "firstUser",
    "password": "example"
}'
```
- Creat task by user.
```
curl --location --request POST 'localhost:8001/task' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZmlyc3RVc2VyIiwiYWRtaW4iOnRydWUsImV4cCI6MTY0NzQ1NzMyMH0.O4gPacbAAB5QMet0JSxk1zHbKIeyl3l2pPzVkxey-jI' \
--header 'Content-Type: application/json' \
--data-raw '{
    "content":"content task"
}'
```

- Get the user task.
```
curl --location --request GET 'localhost:8001/task?create_date=2022-03-14' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZmlyc3RVc2VyIiwiYWRtaW4iOnRydWUsImV4cCI6MTY0NzQ1NzMyMH0.O4gPacbAAB5QMet0JSxk1zHbKIeyl3l2pPzVkxey-jI'
```
![alt text](docs/flow.png)
