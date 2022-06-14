
1. setup env: run cmd bellow

mpm install
node migrations/database.js

2. run locally

npm start

- api get all user: GET: http://localhost:3000/user/list
   response:
    [
        {
            "_id": "62a8d6faecf2e20029fd3435",      => user_id used for todo api
            "user_name": "user01",
            "email": "user01@gmail.com",
            "limit_task_per_day": 1                 => limited user todo tasks per day
        },
        ...
    ]

- api create a todo task

    params: {
        title: 
        description:
        user_id:   // user_id is user unique ID get from api get all user.
    }

    curl -d "title=title 1&description=description 1&user_id=62a8d6faecf2e20029fd3435" -X POST http://localhost:3000/todo/add