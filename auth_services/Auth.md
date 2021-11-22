Auth API
========

1. login:
---------

    Note: For generating a base64 string for user: admin, pass: abc123
    $ echo -ne "admin:abc123" | base64
    YWRtaW46YWJjMTIz

- Endpoint: /api/v2/auth/login
- Example:
    curl -i -X POST -H "Authorization: basic YWRtaW46YWJjMTIz" http://localhost:8080/api/v2/auth/login

       
2. create user:
---------------

- Endpoint: /api/v2/auth/users
- Example:
    curl -X POST -H "Authorization: bearer $OS_TOKEN" -H "Content-Type: application/json" -d '{"name": "newuser1", "password": "changeme", "enaled": true, "email" : "test123@123", "phone": "123456","number_task":"10"}' "http://localhost:8080/api/v2/auth/users"

3. delete user:
---------------

- Endpoint: /api/v2/auth/users/{user_id}
- Example:
    curl -X DELETE -H "Authorization: bearer $OS_TOKEN" -H "Content-Type: application/json" http://localhost:8080/api/v2/auth/users/5aeeede4c5044768979267f68ff63545

If you want to delete user permanently use api: /api/v2/auth/users/{user_id}/permanently

4. Update user:
----------------------------------

- Endpoint: /api/v2/auth/users/{user_id}
- Example:
    curl -X PATCH -H "Authorization: bearer $OS_TOKEN" -H "Content-Type: application/json" -d '{"email": "test@test", "phone": "123456", "status": {"is_enabled": true}}' http://localhost:8080/api/v2/auth/users/b550778e2de2420c9dfd373b4837229e

5. reset password (require admin):
----------------------------------

- Endpoint: /api/v2/auth/users/{user_id}/reset
- Example: 
    curl -X PATCH -H "Authorization: bearer $OS_TOKEN" -H "Content-Type: application/json" -d '{"password": "changeme"}' http://localhost:8080/api/v2/auth/users/cd160623d03c4d4e855d5e1bb894f9ad/reset
    
6. change password:
-------------------

- Endpoint: /api/v2/auth/users/{user_id}/password
- Example: 
    curl -X PATCH -H "Authorization: bearer $OS_TOKEN" -H "Content-Type: application/json" -d '{"original_password": "changeme", "password": "abc123"}' http://localhost:8080/api/v2/auth/users/1f15668d186a4ffdb50b3120ba0883d7/password


