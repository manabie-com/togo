### How to run

1. Run command ```go mod vendor```
2. Run command ```cp .\.env.example .\.env``` (if running on window) or copy data from .env.example to new .env file 
3. Run command ```docker compose up -d --build```
4. Run command ```docker compose up -d```
5. Run command ```docker compose restart```

### Sample curl command to call API
1. API get todo list user:

    ```
    curl --location --request GET 'http://localhost:9999/api/v1/todo/user?user_id=6ac34862-4322-4437-9f98-e87fb6e8371b'
    ```
2. API create todo
    
   ```
   curl --location --request POST 'http://localhost:9999/api/v1/todo/user' --header 'Content-Type: application/json' --data-raw '{
    "user_id": "6ac34862-4322-4437-9f98-e87fb6e8371b",
    "description": "Complete API login user"
    }'
   ```

### Benefits of this solution
- An effective testing
- Frameworks are isolated in individual modules

### What is need upgrade
- Init database with docker compose is delay
- Valid data input API
- Relationship database
- More business rule
- More unit test & function test
