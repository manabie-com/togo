# Todo app
Expose one API to record todo tasks, with different limit of maximum tasks can be created for each user.
## How to run
```shell
go install
docker-compose up -d
cp config.example.yaml config.yaml
go run .
```

## Example
When the app starts, a sample user and a rule is associated to the user are created for testing.

Start adding some tasks
```shell
curl --location --request POST 'localhost:8080/users/9725cc63-4e92-4893-a6b2-216617f3a5dd/tasks' \
--header 'Content-Type: application/json' \
--data-raw '{
    "content": "fetch milk"
}'
```
Run the curl command above a few times to see the error.

Update limiting rule
```shell
curl --location --request PUT 'localhost:8080/users/9725cc63-4e92-4893-a6b2-216617f3a5dd/rules' \
--header 'Content-Type: application/json' \
--data-raw '{
    "action": "tasks/create",
    "unit": "day",
    "requests_per_unit": 10
}'
```

You can also update rule to limit the number of requests per minute instead of day
```shell
curl --location --request PUT 'localhost:8080/users/9725cc63-4e92-4893-a6b2-216617f3a5dd/rules' \
--header 'Content-Type: application/json' \
--data-raw '{
    "action": "tasks/create",
    "unit": "minute",
    "requests_per_unit": 5
}'
```
The supported units are `day`, `hour`, `minute` and `second`.

## Test
```shell
go test -v ./...
```

# About the solution
## Data model
![diagram](https://user-images.githubusercontent.com/40640560/160240700-2c0d6a41-df76-4a4e-9972-373563a2be68.png)


Rule is flexible and new rule can be created to throttle requests for any user action with small amount of code change.
For example, to throttle number of update requests per second, we can create a rule like this:
```json
{
  "user_id": "9725cc63-4e92-4893-a6b2-216617f3a5dd",
  "action": "tasks/update",
  "unit": "minute",
  "requests_per_unit": 5
}
```
And then place `RateLimiter("tasks/update")` next to the route handler. However, to apply a global rule for all users, we have to tweak the middleware a bit to make it work.

## Overview
From high level, a rate limiter middleware is used to handle user requests before sending it to lower layers.
The number of requests an user can perform is limited by a rule.
![image](https://user-images.githubusercontent.com/40640560/160240960-92d6b1fb-33b5-4a02-aace-bcc38f53703c.png)


## How rate limiter middleware works
Rule is stored persistently in MySQL database. Because this data is not expected to be changed often, it is cached in Redis using cache-aside model (server fetch from Redis first, if not found, fetch from MySQL and cache rule data in Redis for subsequent requests).

Each user has a bucket to store the number of requests they have performed in the past time window.
Bucket key is a combination of user id, action and time window, followed the format `counter:{user_id}:{action}:{unit}:{time_window}`. That means each time window for each user and action will have its own counter.

Example `counter:9725cc63-4e92-4893-a6b2-216617f3a5dd:tasks/create:day:26` means the number of task creation requests performed by user 9725cc63-4e92-4893-a6b2-216617f3a5dd on 26th of current month

When a request arrives, the rate limiter middleware do the following:
- Fetch limit rule associated with the user and action
  - If rule is not found, the request is passed
  - If rule is found, proceed to next step
- [INCR](https://redis.io/commands/incr/): increase the counter by 1
- [EXPIRE](https://redis.io/commands/expire/): set timeout for the counter, if this is the first request
- Check if the limit is reached or not by comparing current counter value with the limit value
   - If the limit is reached, request is rejected with 429 status code
   - If the limit is not reached, request is passed to the next middleware to be further processed
  
Advantages:
- Simple to implement and understand
- Memory efficient
- Resetting at the end of time window fits this requirement

Disadvantages:
- If requests come at the edges of time window, the system may allow more requests than the set limit. For example, one user create N tasks at last minute of one day, and create N tasks at first minute of next day. This is not a problem for current requirement, but can be a problem in a more extreme case like only N requests are allowed per one second.


Other approaches:
- Sliding window log: keep tracks of request timestamps (can be used with Redis's sorted set), outdated timestamps are removed when new requests come in
- Sliding window counter: take into account number of requests in the previous time window
- Token bucket: use a bucket of pre-defined token (task ID in this scenario), each request consumes a token, and tokens are put into the bucket at a specific rate
- Leaky bucket: requests are put into a queue (bucket), the bucket is leaked at a preset rate to allow requests to pass through

# Improvements
## Authentication/Authorization
For the sake of simplicity, this part is not implemented. In practice, authentication/authorization is required for some routes to be accessible.

## Avoid redis counter key being leaked
In the rate limiter middleware, the counter key is incremented first and then set an expiration time if its value is 1 (the first time action is performed). If for some reason, the client performs INCR but does not perform EXPIRE, the key will be leaked until we get the same user action at the same unit of time again.

There are several ways to avoid this:

1. Use [EVAL](https://redis.io/commands/eval/) command with Lua script to make the process of incrementing and setting timeout into one atomic operation
2. Use [MULTI](https://redis.io/commands/multi/) and [EXEC](https://redis.io/commands/exec/) commands, also make it into one atomic operation, but this way we always have to set expiration for the key for every request
3. Use other limiter approaches mentioned in previous section

## Error handling
For the time being, 400 error is returned for all errors happen in the server. An error handler middleware can be implemented to make error logging and response message more meaningful.
