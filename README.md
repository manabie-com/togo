### Description
This API is to implement rate limit handling for task creation, the API was written in Nodejs.
This API follow clean architecture with `Controller` -> `Service` -> `Model` layer and trying to follow as much as possible SOLID principle.

### Tech stack
- Nodejs
- MongoDB
- Docker

### Idea, Model
There are 2 models in this case
- User
  - every user has his own quota of creating task, the `max_post_by_day` (default `MAX_NUMBER_TASK_CREATED` = 3)indicate maximum number of task user can create while the `remaining_post` is number of remaining tasks user is allowed to create. the `last_task_created_at` is date of the last task was created. (its null if user has never created task before)
  - every time a task is created, app check the `remaining_post` , if user has no more quota on that day then return `IN_SUFFICIENT_QUOTA` error, otherwise duduct 1 task number.
  
   ```
   {
        email: { type: String, unique: true },
        quota: {
            max_post_by_day: { type: Number },
            last_task_created_at: { type: Date },
            remaining_post: { type: Number }
        }
   }


- Task
  - Simply a task object with author information
   ```
   {
      title: { type: String },
      content: { type: String },
      author: { type: ObjectID, ref: 'user' }
   }

### How to run
- start mongo via docker : `docker-compose up`
- run app `npm run start`

### Give a try
- try api via swagger API at : `http://localhost:3001/api-docs`
- or curl
  - Create user
  ```
  curl -X POST \
    'localhost:3001/user/register' \
    --header 'Accept: */*' \
    --header 'User-Agent: Thunder Client (https://www.thunderclient.com)' \
    --header 'Content-Type: application/json' \
    --data-raw '{
    "email": "user999@gmail.com",
    "password" : "123"
  }'
  ```

  - Login
  ```
    curl -X POST \
    'localhost:3001/user/login' \
    --header 'Accept: */*' \
    --header 'User-Agent: Thunder Client (https://www.thunderclient.com)' \
    --header 'Content-Type: application/json' \
    --data-raw '{
    "email": "user999@gmail.com",
    "password" : "123"
  }'
  ```
  - Get User
  ```
  curl -X GET \
    'localhost:3001/user/64342477f3de8adde575c27f' \
    --header 'Accept: */*' \
    --header 'User-Agent: Thunder Client (https://www.thunderclient.com)' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjY0MzQyNDc3ZjNkZThhZGRlNTc1YzI3ZiIsImlhdCI6MTY4MTEzODk1Nn0.qh5L0PQUaYjzwXXN-CnQsR0Fhg_4NrIMxnK1u2QAVgY' \
    --header 'Content-Type: application/json' \
    --data-raw '{
    "email": "hai@gmail.com"
  }'
  ```

  - Create task (send more than `MAX_NUMBER_TASK_CREATED` to hit rate limit)
  ```
    curl -X POST \
    'localhost:3001/task' \
    --header 'Accept: */*' \
    --header 'User-Agent: Thunder Client (https://www.thunderclient.com)' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjY0MzQyNDc3ZjNkZThhZGRlNTc1YzI3ZiIsImlhdCI6MTY4MTEzODk1Nn0.qh5L0PQUaYjzwXXN-CnQsR0Fhg_4NrIMxnK1u2QAVgY' \
    --header 'Content-Type: application/json' \
    --data-raw '{
    "title": "this is title",
    "content" : "123"
  }'
  ```

### Run Test
- unit test: `npm run test:unit`
- integration test: `npm run test:integration`
- all test: `npm run test`

### Notes
- What if I have more time?
  - I will try to find better solution for scale 
  

