### Requirements

- [x] Implement one single API which accepts a todo task and records it
  - [x] There is a maximum **limit of N tasks per user** that can be added **per day**.
  - [x] Different users can have **different** maximum daily limit.
- [x] Write integration (functional) tests
- [x] Write unit tests
- [x] Choose a suitable architecture to make your code simple, organizable, and maintainable
- [x] Write a concise README
  - [x] How to run your code locally?
  - [x] A sample “curl” command to call your API
  - [x] How to run your unit tests locally?
  - [x] What do you love about your solution?
  - [x] What else do you want us to know about however you do not have enough time to complete?

### How to run my code locally?
- Require: Docker
- Create .env file.
```
cp .env.example .env
```
- Build and run source with docker
```
docker-compose up -d --build
```
- Install composer
```
docker exec -it manabie_togo composer install
```
- Migrate table and seed data (users)
```
docker exec -it manabie_togo php artisan migrate:fresh --seed
```
- Run test (unit test and integration test)
    + Unit test: `/tests/Unit'
    + Function test: `/tests/Feature'
```
docker exec -it manabie_togo php artisan test
```

### Sample “curl” command to call your API
1. Get all users
```
curl http://localhost:8088/api/users
```
2. Login
```
curl -d '{"user_name":"{user_name}","password":"password"}' -H "Content-Type: application/json" -X POST http://localhost:8088/api/auth/login
```
- `user_name`: get one `user_name` from api get users.
- `password`: default `password`
- ex: curl -d '{"user_name":"altenwerth.cory","password":"password"}' -H "Content-Type: application/json" -X POST http://localhost:8088/api/auth/login
3. Get user's task
```
curl http://localhost:8088/api/tasks -H "Accept: application/json" -H "Authorization: Bearer {token}"
```
+ `token`: get token from api login.
+ ex: curl http://localhost:8088/api/tasks -H "Accept: application/json" -H "Authorization: Bearer 3|Ka5ftoBfueY854hobEFflPO2BK6so6EhJEwlpP01"
4. Create new task
```
curl -d '{"name":"Test add new task"}' -H "Content-Type: application/json" -H "Authorization: Bearer {token}" -X POST http://localhost:8088/api/tasks
```
+ `token`: get token from api login.
+ ex: curl -d '{"name":"Test add new task"}' -H "Content-Type: application/json" -H "Authorization: Bearer 3|Ka5ftoBfueY854hobEFflPO2BK6so6EhJEwlpP01" -X POST http://localhost:8088/api/tasks
### What do you love about your solution?
I applied repository pattern, beyond crud struct to my code simple, organizable, and maintainable.

### What else do you want us to know about however you do not have enough time to complete?
If I have more time, I want to apply CI/CD to my project.
