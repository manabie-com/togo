# TODO APIs Guideline

### 1. Installation:

    - OS for development: Windows 10.
    - Install docker desktop.

### 2. How to run the app on local machine:

#### Step 1:

    - Start docker on windows.

#### Step 2:

    - Cd to folder that contains source code of application(pulled from GIT).
    - Run this command: docker-compose up. Todo app will be setup and migrate automatically.
    - If the docker desktop asks for permission to access <togo\database> and <todo> folders then just allow it.
    - After above command done: 2 services will be created that are <backend> and <db>.

#### Step 3:

> How to create a super user in Todo application?
> You could follow one of these ways:

##### 1. Create by registration API:

    - If you want to create a user as super user, you can make a POST request to create a new user and add <is_superuser> field to payload.
    - Endpoint: POST /api/users/registration/
    - Example:
      - curl -X POST "http://localhost:8000/api/users/registration/" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"username\": \"admin\", \"password\": \"admin\", \"is_superuser\": true}"

##### 2. Create by django command:

    - Make sure you are in Todo project folder
    - Access docker container of <backend> and create super user:
      - docker-compose exec backend sh
      - run this command: python manage.py createsuperuser. Now you can set <username> and <password> for this super user.

#### Step 4:

    - At this step TODO APIs are ready for use.

### 3. In case you want to run migration manually(Optional):

1.  How to access docker container and run migration:
    - Access: `docker-compose exec backend sh`
    - Run migrations:
      - `python manage.py makemigrations`
      - `python manage.py migrate`
      - `python manage.py createsuperuser`
        - You can choose any USERNAME and PASSWORD
        - Example account: `username=admin & password=admin`

### 4. Util commands:

    - How to identify docker container ip address:
      - docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' <container_id>

### 5. Todo task APIs:

#### 1. Registration:

    - End point: POST /api/users/registration/
    - Create a new user

#### 2. Login:

    - End point: POST /api/login/
    - Need an user account(username & password) that are created by Registration API

#### 3. Todo and User:

    - End points:
      - GET | POST: /api/tasks/
      - GET | PUT | DELETE: /api/tasks/{task_id}/
      - GET: /api/users/
      - PUT: /api/users/{user_id}/
    - Only admin can change maximum of tasks of user per day

### 6. Curl guideline:

#### 1. Users:

    - POST:
      - curl -X POST "http://localhost:8000/api/users/registration/" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"username\": \"user_001\", \"password\": \"Aa123456\"}"
    - GET(Allow for admin user only):
      - curl -X GET "http://localhost:8000/api/users/" -H "Authorization: Bearer {access_token}"
    - PUT(Allow for admin user only):
      - curl -X PUT "http://localhost:8000/api/users/{user_id}/" -H "Authorization: Bearer {access_token}" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"maximum_task_per_day\": 15}"

#### 2. Login:

    - Normal user:
      - curl -X POST "http://localhost:8000/api/login/" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"username\": \"user_001\", \"password\": \"Aa123456\"}"
    - Admin user:
      - curl -X POST "http://localhost:8000/api/login/" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"username\": \"admin\", \"password\": \"admin\"}"
    - Get refresh_token:
      - curl -X POST "http://localhost:8000/api/token/refresh/" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"refresh\": \"{refresh_token}\"}"

#### 3. Tasks:

    - POST:
      - curl -X POST "http://localhost:8000/api/tasks/" -H "Authorization: Bearer {access_token}" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"title\": \"test adding task 001\", \"description\": \"test adding task 001\"}"
    - GET:
      - curl -X GET "http://localhost:8000/api/tasks/" -H "Authorization: Bearer {access_token}"
      - curl -X GET "http://localhost:8000/api/tasks/{task_id}/" -H "Authorization: Bearer {access_token}"
    - PUT:
      - curl -X PUT "http://localhost:8000/api/tasks/{task_id}/" -H "Authorization: Bearer {access_token}" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"title\": \"updating task 001\"}"
    - DELETE:
      - curl -X DELETE "http://localhost:8000/api/tasks/{task_id}/" -H "Authorization: Bearer {access_token}"

#### 4. Noted:

    - After login successfull <access_token> and <refresh_token> will be included in response.
    - Get <user_id> after make a request GET list of users by super user.
    - Get <task_id> after make a request GET list of tasks by current user.

### 7. Unit Test and Integration Test:

#### Assumptions:

> Only admin user can change maximum of task per day for normal user.

##### 1. How to run test:

    - Access docker container <backend>:
      - Cd to folder contains source code of Todo Application
      - Run command: `docker-compose exec backend sh`
    - Run all of tests in one time:
      - Run command: `python manage.py test`
    - Run specific module:
      - Run test for **baseuser** module: `python manage.py test baseuser`
      - Run test for **todo** module: `python manage.py test todo`

##### 2. User Scenario:

    - [ x ] Allow create a new user for any user.
    - [ x ] Username should be unique.
    - [ x ] Username and password fields should not be null.
    - [ x ] Username and password fields should not be blank.
    - [ x ] Can login with a created user(username / password).
    - [ x ] JWT pair tokens should be in response after login successful.
    - [ x ] Only admin can update <maximum_tasks_per_day> per day for other users.
    - [ x ] Only allow update <maximum_tasks_per_day>.
    - [ x ] Only admin can get list of users.
    - [ x ] Only admin can get detail of a user.
    - [ x ] Only current user can get detail of that user.
    - [ x ] The current user cannot get detail of others.

##### 3. Task Scenario:

    - [ x ] Authorize by JWT Access Token(Bearer as a prefix).
    - [ x ] Only authenticated user can call APIs(GET | POST | PUT | DELETE).
    - [ x ] GET list of tasks should be return list of tasks from the current user(assert fail if in the response includes others).
    - [ x ] GET single task should be return a task from the current user(assert fail if in the response includes others).
    - [ x ] PUT should be updating the task from the current user(assert fail if it can update from others).
    - [ x ] DELETE should be deleting the task from the current user(assert fail if trying to delete from others).
    - [ x ] POST should be creating a new task for the current user.
    - [ x ] Validate <number_of_current_tasks> smaller or equals <maximum_tasks_per_day> of the current user.
    - [ x ] New value of <maximum_task_per_day> should be greater or equal to the current value.

### 8. What do I love about my solution?

    - The Todo application can be installed and run automatically by docker.
    - Validate the number of tasks per day by datetime.
    - Hide <user_id> and <task_id> in url.

### 9. What else do I want you to know about however I do not have enough time to complete?

> If I have more time, I will

    - Implement Todo APIs as an asynchronous then it can handle multiple concurrent requests.
    - Research more about security.
    - Write more tests.
