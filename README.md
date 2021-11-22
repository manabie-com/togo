### Requirements.
 - This project only supports to run on Ubuntu currently
 - go version >= 1.16
 - docker
 - docker-compose
### Install Protobuffer
 - https://github.com/protocolbuffers/protobuf/blob/master/src/README.md

### Install Docker & docker-compose.
 * for Ubuntu : 
    - [Docker](https://docs.docker.com/engine/install/ubuntu/)
    - [docker-compose](https://docs.docker.com/compose/install/)

### Start docker-compose to run MySQL database & KeyStone.
    $cd docker
    $docker-compose up -d

### Build & Run mini_project.
    $make
    $./mini_project

### How To test

1. login: run login by terminal windows:
----------------------------------------
    Note: For generating a base64 string for user: admin, pass: abc123
    $ echo -ne "admin:abc123" | base64
    YWRtaW46YWJjMTIz
    - Endpoint: /api/v2/auth/login
    - Example:
        curl -i -X POST -H "Authorization: basic YWRtaW46YWJjMTIz" http://localhost:8080/api/v2/auth/login

2. assign OS_TOKEN=Key_got_from_above:
-------------------------------------
    - Example: OS_TOKEN=eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiJ9.eyJzdWIiOiIwMTgwNjVmNjlkYTY0MTgzYjNmNDc5MDAwMWRkOWE5ZSIsImlhdCI6MTYzNzUwNjY4NSwiZXhwIjoxNjM3NTEwMjg1LCJvcGVuc3RhY2tfbWV0aG9kcyI6WyJwYXNzd29yZCJdLCJvcGVuc3RhY2tfYXVkaXRfaWRzIjpbIk5GbTRKWFBSVHBXRVowdlpFWHgzX0EiXSwib3BlbnN0YWNrX3Byb2plY3RfaWQiOiIzYThlZGMyZDcwMDE0NmViOWVjMGQyZjM1M2YxMjQ3ZiJ9.DKJzr82g6My4z3FoshegAvlz1zF1yiZSCMJ-RQL9LP7yLxjYE4Oy107zBsqDBZ9cQHEn0cT66V2pzMhTwyXrjQ

3. Create new User:
-------------------
    - Endpoint: /api/v2/auth/users
    - Example:
        curl -X POST -H "Authorization: bearer $OS_TOKEN" -H "Content-Type: application/json" -d '{"name": "newuser1", "password": "changeme", "enaled": true, "email" : "test123@123", "phone": "123456","number_task":"10"}' "http://localhost:8080/api/v2/auth/users" | jq '.'

4. Get the List of User ID:
---------------------------
    - Endpoint: /api/v2/auth/userids
    - Example:
        curl -X GET -H "Authorization: bearer $OS_TOKEN" -H "Content-Type: application/json" "http://localhost:8080/api/v2/auth/userids" | jq '.'

5. Update user:
---------------
    - Endpoint: /api/v2/auth/users/{user_id}
    - Example:
    curl -X PATCH -H "Authorization: bearer $OS_TOKEN" -H "Content-Type: application/json" -d '{"user_name" : "test", "email": "test@test", "phone": "0904123456","number_task":12}' http://localhost:8080/api/v2/auth/users/aae669b279cd4cf999b85a57d93ffc83 | jq '.'

6. Create new Task for user:
----------------------------
    - Endpoint: /api/v1/newtask
    - Ex: $ curl -s -H "Authorization: Bearer $OS_TOKEN" -H "Content-Type: application/json" -d '{"user_id":"aae669b279cd4cf999b85a57d93ffc83","task_name": "task3"}' http://localhost:8080/api/v1/newtask | jq '.'