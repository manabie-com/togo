### Desctiption
- This project provide some API which allow to:
  - Create user with infomation about maximum limit of tasks which user can added per day
  - Update user to change maximum limit of tasks when have demand
  - Assign 'todo' task for someone.
  - Create tasks with maximum limit per day per user (different user have different maximum daily limit)
  - Update/ Delete tasks owner
- Admin Dashboard.
- Technology used:
  - Python 3
  - Django REST Framework
  - Authentication method: Basic Authentication
  - Database: postpresql
### Usage
- Require: Python 3.8.10
- Install requirements modules: `pip install -r requirements.txt`
- Change directory to project folder: `cd todo`
- Run migrate: `python manage.py migrate`
- Start server: `python manage.py runserver` this command will start server at: `http://127.0.0.1:8000/`
- List API:
  - User:
    - Create User: `curl -d "username=username&password=password&number_todo_number=number_todo_number" -X POST http://127.0.0.1:8000/todo_api/user/register/`
    - Retrieve User info: `curl -u username:password http://127.0.0.1:8000/todo_api/user/`
    - Update User: `curl -u username:password -d "number_todo_number=number_todo_number" -X PUT http://127.0.0.1:8000/todo_api/user/`
  - Task:
    - Create Task: `curl -u username:password -d "title=title&status=status" -X POST http://127.0.0.1:8000/todo_api/todo/`
    - Retrieve User todo Tasks: `curl -u username:password http://127.0.0.1:8000/todo_api/todo/`
    - Retrieve Task Detailed: `curl -u username:password http://127.0.0.1:8000/todo_api/todo/<todo_id>/`
    - Update Task: `curl -u username:password -d "title=title&status=status" -X PUT http://127.0.0.1:8000/todo_api/todo/<todo_id>/`
    - Delete Task: `curl -u username:password -X DELETE http://127.0.0.1:8000/todo_api/todo/<todo_id>/`
- Document API: `https://documenter.getpostman.com/view/8094264/Uz5GpcMu`
- Collection API Postman example: `https://www.postman.com/collections/bf7371e42ac451c315e7`
- Admin: `http://127.0.0.1:8000/admin/`
### Test
- Test case location:`togo/todo/todo_api/test.py`
- Run Test: `python manage.py test`
- List TestCase:
  - UserTestCases:
    - create user success
    - create user without username
    - create user with username is exist
    - create user with username exceed 50
    - create user with short password (<4)
    - create user with number_todo_number is less than 0
    - create user with number_todo_number is not integer
    - create user with number_todo_number is none

    - retrieve user success
    - retrieve user id not exist

    - update user success
    - update user with number_todo_number is less than 0
    - update user with number_todo_number is not integer
    - update user with number_todo_number is none

  - TaskTestCases:
    - create todo task success
    - create todo task without task name
    - create todo task with completed is not boolean
    - create todo task exceed daily limit (number_todo_number)

    - retrieve list todo task by user id
    - retrieve todo task by task id

    - update todo task success
    - update todo task with empty task name
    - update todo task with completed is not boolean

### Step quickly to run the Togo project
- `pip install -r requirements.txt`: ``make install`
- `python manage.py makemigrations`: `make makemigrations`
- `python manage.py migrate`: `make migrate`
- `python manage.py runserver`: `make run`
-  Create new user to login Togo project: `https://documenter.getpostman.com/view/8094264/Uz5GpcMu#e00d3bb9-9fd8-4252-99cb-7cb5b572329b`
- `python manage.py test`: `make test`
