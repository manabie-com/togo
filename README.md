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

# TOGO
1. How to run your code locally?
2. A sample “curl” command to call your API
3. How to run your unit tests locally?
4. What do you love about your solution?
5. What else do you want us to know about however you do not have enough time to complete?

## How to run your code locally?
- Docker: docker build -t todoapp . && docker run -d -p 8080:8080 todoapp
- Local: setup postgres and set env into .env && go run main.go server

## Curl example
Get List Task

```
curl --location --request GET 'http://localhost:8080/api/task' \
--header 'userID: 123456' 
```

Create Task

```
curl --location --request POST 'http://localhost:8080/api/task' \
--header 'userID: 123456' \
--header 'Content-Type: application/json' \
--data-raw '{
    "content":"todo 1"
}'
```

Delete Task

```
curl --location --request DELETE 'http://localhost:8080/api/task/2' \
--header 'userID: 123456' 
```

Link Postman (https://www.getpostman.com/collections/6d379462a10d0825c40f)

3. Unit tests locally with sqlite

```
  go test ./...
```

4. Design DB
```
create table tasks (
    id serial primary key,
    content text,
    user_id varchar(36),
    status smallint default 1,
    created_at timestamp without time zone default current_timestamp,
    updated_at timestamp without time zone default current_timestamp,
    deleted_at timestamp without time zone default NULL
);

create index idx_tasks_user_id_status on tasks (user_id,status);
```

- The source structure is simple and clear

5. What else do you want us to know about however you do not have enough time to complete.
- Use cache to calculate the number of user tasks in 1 day
- Add monitor service like sentry, prometheus
- Deploy to cloud
