## Running Test

```phpt
./vendor/bin/phpunit tests/Feature/ListApiTest.php
```

## Api post test
```shell
php artisan serve

curl -d '{"task":"value1", "description":"value2", "is_complete":"0"}' -H "Content-Type: application/json" -X POST http://127.0.0.1:8000/api/lists/create
```
