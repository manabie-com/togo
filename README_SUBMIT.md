## How to run your code locally?

1. Create a MySQL Database.
2. Edit database connection values (MySQl - Host, Port, database name) in **.env** file.
3. Open Terminal, cd to the project directory and run the first time

```batch
make run-dev
```

4. Insert into database after first running:

```sql
insert into users (id, limit_task)
values (1, 5) #5 is Limit task per day of user
```

5. Restart server

<br>

## A sample “curl” command to call your API

```curl
curl --header "Content-Type: application/json" --request POST --data '{"id": 1, "userId": 1, "title": "Task CURL", "description": "Task CURL", "isCompleted": false}' http://localhost:8000/tasks
```

<br>

## How to run your unit tests locally?

-   Open Terminal, cd to the project directory and run

```batch
make run-test
```

<br>

## What do you love about your solution?

This is my first time trying Golang, so my love thing is probably it's written in Go.

<br>

## Architecture:

-   MVC

<br>

## What else do you want us to know about however you do not have enough time to complete?

-   Integration (functional) tests
