# Task API

## Requirements

- Docker

## Project setup
```
docker run --rm \
    -u "$(id -u):$(id -g)" \
    -v $(pwd):/var/www/html \
    -w /var/www/html \
    laravelsail/php80-composer:latest \
    composer install --ignore-platform-reqs

cp .env.example .env
./vendor/bin/sail php artisan key:generate
./vendor/bin/sail php artisan migrate
```

## Usage

### Register

```
curl -XPOST -H 'Accept: application/json' -H "Content-type: application/json" -d '{"name": "User", "email": "email@test.com", "password": "password", "password_confirmation": "password"}' 'http://localhost/api/register'
```

### Login

```
curl -XPOST -H 'Accept: application/json' -H "Content-type: application/json" -d '{"email": "email@test.com", "password": "password"}' 'http://localhost/api/login'
```
This will return a token. You need to use at as a Authorization header for api's protected by authentication.

### Logout (Auth protected)

```
curl -XPOST -H 'Accept: application/json' -H 'Authorization: Bearer 2|q2702K7EybJuvmaFa3uOcEHymOTKS3sbbqt3RUvG' -H "Content-type: application/json" 'http://localhost/api/logout'
```

### Create Task (Auth protected)

```
curl -XPOST -H 'Accept: application/json' -H 'Authorization: Bearer 3|wMejWoyjQXgwCqUjlQdajkAFyiLlPwgQLy5Dn8Lr' -H "Content-type: application/json" -d '{"name": "Sample task"}' 'http://localhost/api/tasks'
```

## Run Test

```
./vendor/bin/sail php artisan test
```

## Misc

This solution uses the repository pattern to abstract data access. I love my solution because it is easy to read and adhears to SOLID principle.
