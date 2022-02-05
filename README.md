## Introduction
- I am using python (flask framework & other libs) to create an API Server
- From my point of view, I create 4 tables (user, task, pricing, invoice) indicate how our "togo" app works
---
## Environment
- Python 3.8+
- make 3.8+
- virtualenv 20.7.2+
- A Unix cmd of your choices (bash/sh/zsh) (I am using zsh)
---
## Installation
- I assume that every command in the 'Environment' installed
1. Moving to root directory after clone it. 
```shell
# You can check it by simple using this command
$ pwd 
# the result should look like
# /some/where/in/your/computer/togo
```
2. Create a virtual environment for python
```shell
$ virtualenv venv -p=python3.8
$ mkdir logs
```

3. Install package
```shell
$ ./venv/bin/pip install -r requirements.txt 
```

4. Run the application
```shell
$ make start
```

5. Run the test
```shell
$ make test 
```

6. Note: I also prepare a script to automate install step.
```shell
$ zsh scripts/setup.sh
```
---
## Code structure
- I prepared 3 main endpoint
1. auth - For user authentication
   1. `auth/sign-up` - POST: for register a user
    ```shell
    $ curl --location --request POST 'http://127.0.0.1:6000/api/auth/sign-up' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "email": "demo-sonlh-3@gmail.com", 
        "password":"sample-password",
        "fullname":"Luu Hoang Son"
    }'
    ```
    Result - successful
    ```shell
    {
        "message": "User created"
    }
    ```
    Result - error 
    
    ```shell
    {
        "description": "reason",
        "error": "error name"
    }
    ```
   2. `auth/sign-in` POST: for sign a user in to get jwt token (for further operations)
    ```shell
    $ curl --location --request POST 'http://127.0.0.1:6000/api/auth/sign-in' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "email": "demo-sonlh-3@gmail.com", 
        "password":"sample-password"
    }'
    ```
    Result - successful
    ```shell
    {
        "token": "eyJ0eXAiOiJKV1QiLCJ..."
    }
    ```
    ```shell
    {
      "description": "reason",
      "error": "error name"
    }
    ```
2. task - For task relative utilities   
  Require a jwt token of a signed-in user. I prove a full jwt here for visualization.
   1. `/task` - GET: get all task belong to signed-in user (I dont use paging & limit of task per request here.) 
    ```shell
    $ curl --location --request GET 'http://127.0.0.1:6000/api/task' \
    --header 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiIsImV4cCI6MTY3MjUwNjAwMC4wfQ.eyJ1c2VySWQiOiJmN2YwNjliYjAyMDA0NDUzYWVhY2I5MWVhZGE0NTZjMSJ9.e2gvc3KZVGwEWJYlBY9L7exeN9W6zezJEemFgo5iF08' \
    --header 'Content-Type: application/json'
    ```
    Result - successful
    ```shell
    {
        "data": [
            {
                "created": "2022-02-04 17:51:25",
                "description": "Its a random quote!",
                "finish": false,
                "id": "ef41eaf8fa114510b822032b5ebb1c42",
                "last_modified": "2022-02-05 00:03:21",
                "summary": "manabie"
            }
        ]
    }
    ```
   2. `/task` - POST: create a task which belong to signed-in user
    ```shell
    $ curl --location --request POST 'http://127.0.0.1:6000/api/task/' \
    --header 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiIsImV4cCI6MTY3MjUwNjAwMC4wfQ.eyJ1c2VySWQiOiJmN2YwNjliYjAyMDA0NDUzYWVhY2I5MWVhZGE0NTZjMSJ9.e2gvc3KZVGwEWJYlBY9L7exeN9W6zezJEemFgo5iF08' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "summary": "manabie", 
        "description":"Its a random quote!"
    }
    '
    ```
    Result - successful
    ```shell
    {
        "message": "Task created",
        "taskId": "0bc910fc2ce14a54aa9617dfe5f6e27d"
    }
    ```
    Result - error if user violate their daily limit
    ```shell
    {
        "description": "Daily limit for user 'f7f069bb02004453aeacb91eada456c1' exceed! Please upgrade to higher pricing options.",
        "error": "Too Many Requests"
    }
    ```
   3. `/task/<task_id>` - GET: get task data of a task by id
    ```shell
    $ curl --location --request GET 'http://127.0.0.1:6000/api/task/0bc910fc2ce14a54aa9617dfe5f6e27d' \
    --header 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiIsImV4cCI6MTY3MjUwNjAwMC4wfQ.eyJ1c2VySWQiOiJmN2YwNjliYjAyMDA0NDUzYWVhY2I5MWVhZGE0NTZjMSJ9.e2gvc3KZVGwEWJYlBY9L7exeN9W6zezJEemFgo5iF08' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "summary": "manabie", 
        "description":"Its a random quote!"
    }'
    ```
    Result - successful. If task id is not found or deleted. Empty object returned
    ```shell
    {
        "data": {
            "created": "2022-02-05 13:26:40",
            "description": "Its a random quote!",
            "finish": false,
            "id": "0bc910fc2ce14a54aa9617dfe5f6e27d",
            "last_modified": "2022-02-05 13:26:40",
            "summary": "manabie"
        }
    }
    ```
   4. `/task/<task_id>` - PATCH: update a task of signed-user by id.  
   Currently, only update 2 fields summary & description are updatable. if any field rather than 2 field above data passed, update operation doesnt run.
    ```shell
    $ curl --location --request PATCH 'http://127.0.0.1:6000/api/task/0bc910fc2ce14a54aa9617dfe5f6e27d' \
    --header 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiIsImV4cCI6MTY3MjUwNjAwMC4wfQ.eyJ1c2VySWQiOiJmN2YwNjliYjAyMDA0NDUzYWVhY2I5MWVhZGE0NTZjMSJ9.e2gvc3KZVGwEWJYlBY9L7exeN9W6zezJEemFgo5iF08' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "summary": "manabie-updated", 
        "description":"Its a random quote!-updated"
    }'
    ```
    Result - successful
    ```shell
    {
        "data": true
    }
    ```
   5. `/task/<task_id>` - DELETE: soft delete a task of signed-user by id
    ```shell
    curl --location --request DELETE 'http://127.0.0.1:6000/api/task/0bc910fc2ce14a54aa9617dfe5f6e27d' \
    --header 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiIsImV4cCI6MTY3MjUwNjAwMC4wfQ.eyJ1c2VySWQiOiJmN2YwNjliYjAyMDA0NDUzYWVhY2I5MWVhZGE0NTZjMSJ9.e2gvc3KZVGwEWJYlBY9L7exeN9W6zezJEemFgo5iF08' \
    --header 'Content-Type: application/json'
    ```
    Result - successful
    ```shell
    {
        "message": "Task '0bc910fc2ce14a54aa9617dfe5f6e27d' deleted!"
    }
    ```
   6. `subscription` - for current pricing options (level)
      1. `/subscription/` - GET : get all available pricing options
       ```shell
       curl --location --request GET 'http://127.0.0.1:6000/api/subscription' \
       --header 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiIsImV4cCI6MTY3MjUwNjAwMC4wfQ.eyJ1c2VySWQiOiJmN2YwNjliYjAyMDA0NDUzYWVhY2I5MWVhZGE0NTZjMSJ9.e2gvc3KZVGwEWJYlBY9L7exeN9W6zezJEemFgo5iF08' \
       --header 'Content-Type: application/json' 
       ```
       Result - successful
       ```shell
       {
           "data": [
               {
                   "daily_limit": 5,
                   "id": "ff06fc09202143dbb311ab19fde3096b",
                   "name": "Basic",
                   "unit_price": 0.0
               },
               {
                   "daily_limit": 10,
                   "id": "a6ebe3f94d4c41de904072c20bdde45a",
                   "name": "Standard",
                   "unit_price": 2.99
               },
               {
                   "daily_limit": 40,
                   "id": "f0bd47ebb20a494c82e341809bdf224a",
                   "name": "Premium",
                   "unit_price": 4.99
               },
               {
                   "daily_limit": 80,
                   "id": "56f1f35285ac4500afbb25cef9b02a6b",
                   "name": "Enterprise",
                   "unit_price": 9.99
               }
           ]
       }
       ```
      2. `/subscription/<price_level_id>` - POST: change current pricing options (change daily limit of a user by make them paid - Just pretending that they paid successfully ) 
       ```shell
       $ curl --location --request POST 'http://127.0.0.1:6000/api/subscription/f0bd47ebb20a494c82e341809bdf224a' \
       --header 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiIsImV4cCI6MTY3MjUwNjAwMC4wfQ.eyJ1c2VySWQiOiJmN2YwNjliYjAyMDA0NDUzYWVhY2I5MWVhZGE0NTZjMSJ9.e2gvc3KZVGwEWJYlBY9L7exeN9W6zezJEemFgo5iF08' \
       --header 'Content-Type: application/json'
       ```
       Result - successful
       ```shell
       {
           "message": "Invoices Created '859e5f37a3c34d6fb5a00b9ff905bfe6'"
       }
       ```

*NOTE*: if any error happens, it will be in format below 

```shell
{
  "description": "reason",
  "error": "error name"
}
```
---
# Original
### Requirements

- Implement one single API which accepts a todo task and records it
  - There is a maximum **limit of N tasks per user** that can be added **per day**.
  - Different users can have **different** maximum daily limit.
- Write integration (functional) tests
- Write unit tests
- Choose a suitable architecture to make your code simple, organizable, and maintainable
- Write a concise README
  - How to run your code locally?
  - A sample “curl” command to call your API
  - How to run your unit tests locally?
  - What do you love about your solution?
  - What else do you want us to know about however you do not have enough time to complete?

### Notes

- We're using Golang at Manabie. **However**, we encourage you to use the programming language that you are most comfortable with because we want you to **shine** with all your skills and knowledge.

### How to submit your solution?

- Fork this repo and show us your development progress via a PR

### Interesting facts about Manabie

- Monthly there are about 2 million lines of code changes (inserted/updated/deleted) committed into our GitHub repositories. To avoid **regression bugs**, we write different kinds of **automated tests** (unit/integration (functionality)/end2end) as parts of the definition of done of our assigned tasks.
- We nurture the cultural values: **knowledge sharing** and **good communication**, therefore good written documents and readable, organizable, and maintainable code are in our blood when we build any features to grow our products.
- We have **collaborative** culture at Manabie. Feel free to ask trieu@manabie.com any questions. We are very happy to answer all of them.

Thank you for spending time to read and attempt our take-home assessment. We are looking forward to your submission.
