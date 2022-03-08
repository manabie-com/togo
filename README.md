### How to run your code locally?
- cd service 
- cp -R .env.example .env
- docker-compose up -d --build
- docker exec -it togo_php_1 bash
- php artisan migrate
- php artisan db:seed

### Example Curl

```
curl --location --request POST 'http://localhost:8099/api/task?user_id=3&name=manabie_test_1'
```

### How to run your unit tests locally?

docker exec -it togo_php_1 bash
- ./vendor/bin/phpunit


### Design patterns Repository

### Code test

```
    tests/Feature/Controller/TaskControllerTest.php
    tests/Services/CheckCreateTaskServiceTest.php
```