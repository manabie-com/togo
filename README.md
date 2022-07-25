### About the app
This app is built in Golang and PostgreSQL (for storing data). Gorilla Mux is the default router which is using in the app and for integration test too.

The app is inspired heavily by Laravel, so you would feel familiar if you had developed with Laravel. I also update JWT Authentication and full CRUD for this app.

### How to run the app
```shell
docker-compose up
```

Default app port is `8080`, if you your `8080` port is not available, do not hesitate to change to another one in `docker-compose.yml`

### How to use the app
1. Register new account
Method: `POST`

URL: `/api/register`

cURL: `curl --location --request POST 'http://localhost:8080/api/register' \
--header 'Content-Type: application/json' \
--data-raw '{
  "email": "huuthuan.nguyen@hotmail.com",
  "password": "123@pass!word"
}'`

Payload:
```json
{
  "email": "huuthuan.nguyen@hotmail.com",
  "password": "123@pass!word"
}
```
Response:
```json
{
    "status": 1,
    "messages": [
        "Successful."
    ],
    "data": {
        "id": 1,
        "email": "huuthuan.nguyen@hotmail.com",
        "is_active": true,
        "daily_limit": 3,
        "create_at": "2022-07-25T18:02:49.291850383Z",
        "updated_at": "2022-07-25T18:02:49.291852258Z"
    }
}
```

Default quota for your tasks is 3 per day, you can not change this.

2. Login with your email and password
Method: `POST`

URL: `/api/auth/login`

cURL: `curl --location --request POST 'http://localhost:8080/api/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
  "email": "huuthuan.nguyen@hotmail.com",
  "password": "123@pass!word"
}'`

Payload:
```json
{
  "email": "huuthuan.nguyen@hotmail.com",
  "password": "123@pass!word"
}
```
Response:
```json
{
    "status": 1,
    "messages": [
        "Successful."
    ],
    "data": {
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Imh1dXRodWFuLm5ndXllbkBob3RtYWlsLmNvbSIsImV4cCI6MTY1ODg1ODY0NH0.vfglN4fOZ7NDat1hxznivNk9T4znCgw4l8K6XFemAYQ",
        "expired_at": "2022-07-26T18:04:04.060592793Z"
    }
}
```

3. Add new task
Method: `POST`

URL: `/api/tasks`

cURL: `curl --location --request POST 'http://localhost:8080/api/tasks' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Imh1dXRodWFuLm5ndXllbkBob3RtYWlsLmNvbSIsImV4cCI6MTY1ODg1ODY0NH0.vfglN4fOZ7NDat1hxznivNk9T4znCgw4l8K6XFemAYQ' \
--header 'Content-Type: application/json' \
--data-raw '{
  "content": "Do homework"
}'`

Payload:
```json
{
  "content": "Do homework"
}
```
Response
```json
{
    "status": 1,
    "messages": [
        "Successful."
    ],
    "data": {
        "id": 1,
        "content": "Do homework",
        "published_date": "2022-07-25",
        "status": 0,
        "created_by": 1,
        "create_at": "2022-07-25T18:04:30.128784138Z",
        "updated_at": "2022-07-25T18:04:30.128785013Z"
    }
}
```

4. Update your task by ID
Method: `PUT`

URL: `/api/tasks/:id`

cURL: `curl --location --request PUT 'http://localhost:8080/api/tasks/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Imh1dXRodWFuLm5ndXllbkBob3RtYWlsLmNvbSIsImV4cCI6MTY1ODg1ODY0NH0.vfglN4fOZ7NDat1hxznivNk9T4znCgw4l8K6XFemAYQ' \
--header 'Content-Type: application/json' \
--data-raw '{
  "content": "Do homework later"
}'`

Payload:
```json
{
  "content": "Do homework later"
}
```
Response:
```json
{
    "status": 1,
    "messages": [
        "Successful."
    ],
    "data": {
        "id": 1,
        "content": "Do homework later",
        "published_date": "2022-07-25",
        "status": 0,
        "created_by": 1,
        "create_at": "2022-07-25T18:04:30.128784138Z",
        "updated_at": "2022-07-25T18:07:45.866218Z"
    }
}
```

5. Remove your task by ID
Method: `DELETE`

URL: `/api/tasks/:id`

cURL: `curl --location --request DELETE 'http://localhost:8080/api/tasks/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Imh1dXRodWFuLm5ndXllbkBob3RtYWlsLmNvbSIsImV4cCI6MTY1ODg1ODY0NH0.vfglN4fOZ7NDat1hxznivNk9T4znCgw4l8K6XFemAYQ'`

Payload:
```json

```
Response
```json
```

6. List your tasks
Method: `GET`

URL: `/api/tasks`

cURL: `curl --location --request GET 'http://localhost:8080/api/tasks' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Imh1dXRodWFuLm5ndXllbkBob3RtYWlsLmNvbSIsImV4cCI6MTY1ODg1ODY0NH0.vfglN4fOZ7NDat1hxznivNk9T4znCgw4l8K6XFemAYQ'`

Payload:
```json

```
Response:
```json
{
    "status": 1,
    "messages": [
        "Successful."
    ],
    "data": {
        "items": [
            {
                "id": 1,
                "content": "Do homework later",
                "published_date": "2022-07-25",
                "status": 0,
                "created_by": 1,
                "create_at": "2022-07-25T18:04:30.128784138Z",
                "updated_at": "2022-07-25T18:07:45.866218Z"
            },
            {
                "id": 1,
                "content": "Do laundry",
                "published_date": "2022-07-25",
                "status": 0,
                "created_by": 1,
                "create_at": "2022-07-25T18:04:30.128784138Z",
                "updated_at": "2022-07-25T18:07:45.866218Z"
            }
        ]
    }
}
```