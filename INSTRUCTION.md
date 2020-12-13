# Instructions

Quick launch

```shell
docker-compose up
```

Run tests

```go
go test ./...
```

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
created_date | string | NO | `yyyy-mm-dd` format

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

## Missions

- [x] Create 5 task limit checker per day
- [ ] Switch from SQLite to Postgres with Docker
- [ ] DRY code
- [x] Unit testing for `service` layer
- [x] Unit testing for `storages` layer (low prio)
- [ ] Integration testing
- [ ] Split `services` layer to `use_case` and `transport` layer

## Improvements

- [ ] Store pw NOT in plain text (need hash + salt)
- [ ] Remove hard-coded JWT key (suggestion: store as ENV variable)
- [ ] Remove hard-coded DB parameters
- [ ] Muxer? (low prio)

## Observations and Thoughts

- Login credentials can be put in query string (x-www-form-urlencoded), which may causes security problem on the front-end. Suggestion: enforce submitting credentials using form data (multipart/form-data) (thus switching to POST instead of GET).
- Per-day limit ambiguity: check within a date or a rolling 24hrs? The `created_date` field in database is `DATE` instead of `DATETIME` so maybe it is the former?
- What happens when you change the per-day limit (or can you change it at all)? Assuming it being unchanged right now is probably good enough.
- No registration as of now
- Refresh token for JWT?
- Database table `users` seems okay
- Database table `tasks` has PK is a `UUID`? Seems suboptimal with lots of task creations, and we search by `user_id` (? double-check this) anyway so it does not seems to be a good indexing. Suggestion: `int` PK internal, `UUID` external.
- `tasks:content` field seems to be cut off (13 chars)
- Benchmarking
