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

1. login: run login by terminal windows
---------
    Note: For generating a base64 string for user: admin, pass: abc123
    $ echo -ne "admin:abc123" | base64
    YWRtaW46YWJjMTIz
- Endpoint: /api/v2/auth/login
- Example:
    curl -i -X POST -H "Authorization: basic YWRtaW46YWJjMTIz" http://localhost:8080/api/v2/auth/login

2. assign OS_TOKEN=Key_got_from_above
    - Example: OS_TOKEN=eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiJ9.eyJzdWIiOiIwMTgwNjVmNjlkYTY0MTgzYjNmNDc5MDAwMWRkOWE5ZSIsImlhdCI6MTYzNzUwNjY4NSwiZXhwIjoxNjM3NTEwMjg1LCJvcGVuc3RhY2tfbWV0aG9kcyI6WyJwYXNzd29yZCJdLCJvcGVuc3RhY2tfYXVkaXRfaWRzIjpbIk5GbTRKWFBSVHBXRVowdlpFWHgzX0EiXSwib3BlbnN0YWNrX3Byb2plY3RfaWQiOiIzYThlZGMyZDcwMDE0NmViOWVjMGQyZjM1M2YxMjQ3ZiJ9.DKJzr82g6My4z3FoshegAvlz1zF1yiZSCMJ-RQL9LP7yLxjYE4Oy107zBsqDBZ9cQHEn0cT66V2pzMhTwyXrjQ
