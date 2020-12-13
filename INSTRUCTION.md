# Instructions

Quick launch

```shell
docker-compose up
```

Shutdown application / test

```shell
docker-compose down
```

Run tests

```shell
docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
```

After testing, use `docker-compose -f docker-compose.test.yml down --volumes` to properly clean up test database.

## What I have (and have not) accomplished

- [x] Daily task limit functionality
- [x] Switch from SQLite to Postgres with `docker-compose`
- [x] Unit tests for `service` layer
- [x] Integration tests
- [ ] DRY code
- [x] (Optional) Unit tests for `storages` layer
- [ ] (Optional) Split `services` layer to `use_case` and `transport` layer

## Potential improvements

- Store password after hashing with salt for security reasons.
- Use environment variables (e.g. `.env` file) to store application parameters to prevent secret key leakages and conveniences (e.g. separating `PROD` and `DEV` environments).
- Give JWT refresh token in addition to access token so that user do not have to authenticate again.
- Instead of using `UUID` as PK for table `tasks`, an auto-increment integer should work better both in terms of inserting and querying data. We can also index `created_date` column for faster querying.
- Clean up the code base, for example:
    1. [Here](internal/storages/sqlite/db.go#L87) we could write `return err == nil`
    2. Rename folder and package away from `sqlite` and to a more generic form
    3. Refactor `tasks_test.go` and `main_test.go`, both of which use some similar code (e.g. `newLoginRequest()` vs `makeLoginRequest()`)

- Benchmarking

## API endpoints

### Login

```
GET /login
```

The response contains an access token (valid for 15 minutes) that can be used to make requests with authorization. For example:

<pre>
CURL localhost:5050/tasks?created_date=2020-11-22 -H "Authorization: <em>YOUR_ACCESS_TOKEN</em>"
</pre>

**Arguments:**
Name | Type | Mandatory | Description
--- | --- | --- | ---
user_id | string | YES |
password | string | YES |

**Response:**

```javascript
{
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDc1MjIyNjAsInVzZXJfaWQiOiJmaXJzdFVzZXIifQ.FZXUnwVIYbrOf6cxX-1dR4DxBaZu56-xytiKF2EAmlU"
}
```

### List user tasks

```
GET /tasks
```

**Login required:** YES

**Arguments:**
Name | Type | Mandatory | Description
--- | --- | --- | ---
created_date | string | YES | `yyyy-mm-dd` format

**Response:**

```javascript
{
    "data": [
        {
            "id": "e1da0b9b-7ecc-44f9-82ff-4623cc50446a",
            "content": "first content",
            "user_id": "firstUser",
            "created_date": "2020-06-29"
        },
        {
            "id": "055261ab-8ba8-49e1-a9e8-e9f725ba9104",
            "content": "second content",
            "user_id": "firstUser",
            "created_date": "2020-06-29"
        },
        ... // additional rows if available
    ]
}
```

### Add user task

```
POST /tasks
```

**Login required:** YES

**Arguments:**
Name | Type | Mandatory | Description
--- | --- | --- | ---
content | string | YES | Description of the task

**Response:**

```javascript
{
    "data": {
        "id": "94740640-0c15-458f-8f7a-6d7eaef356cb",
        "content": "Task description",
        "user_id": "firstUser",
        "created_date": "2020-12-09"
    }
}
```
