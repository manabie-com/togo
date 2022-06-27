WorkDirectory:


    ├─┬ togo
    │ ├─┬ apps                          # Directory for api using rest-api
    │ │ ├─┬ exceptions                  # Directory for exceptions
    │ │ ├─┬ migrations                  # Directory for migrations
    │ │ ├─┬ models                      # Directory for Models
    │ │ ├─┬ urls                        # Directory for Controllers
    │ │ ├─┬ views                       # Directory for Views
    │ │ ├─┬ serializers                 # Directory for Manage request and response api
    │ │ ├─┬ tasks                       # Directory for Manage celery task
    │ ├─┬ logging                       # Directory for manage logs
    │ ├─┬ test                          # Directory for Unit test
    │ ├─┬ togo                          # Directory for Settings system
    │ │ ├─┬ celery                      # Directory for Config celery
    ├─ docker-compose.yml               # Docker build django
    ├─ Dockerfile                       # Docker build sh
    ├─ entrapoint.sh                    # Docker build sh
    ├─ README.md                        # Readme file 
    ├─ requirements.txt                 # Package file requirements
    ├─ .gitignore                       # Git-ignore file
    ├─ manage.py                        # Run django


Requiments:

    - python: 3.8.6
    - django: 3.2.5
    - redis : 7.0.0

Run Project:

    1. mkdir logging public public/media public/static public/staticfiles
    2. pip install virtualenv
    3. python -m virtualenv env
    4. source env/bin/activate
    5. pip install -r requirements.txt
    6. brew install redis
    7. brew services start redis
    8. celery -A togo.celery worker -B -l info
    9. python manage.py runserver

Run Project By Docker:

    mkdir logging public public/media public/static public/staticfiles
    pip freeze > requirements.txt
    docker-compose up

Tutorial Curl :
    
    1. First you need call api to get Token:
       - curl -X POST "http://localhost:8000/api/users/token/" -H  "accept: application/json" -H  "Content-Type: application/json" -H  "X-CSRFToken: O8AESU8podSIZYorLrwc162N7ZdvVkdZpmouAROiTo9BzOSeGjRcBq9oljKG58i8" -d "{  \"username\": \"a\",  \"password\": \"1\"}"
    2. Second you change token below after beer:
       - curl -X POST "http://localhost:8000/api/users/assignment/" -H  "accept: application/json" -H  "Content-Type: application/json" -H  "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiZXhwIjoxNjU2MTMzNjk3LCJpYXQiOjE2NTYxMzMzOTcsImp0aSI6IjM2Y2I4ZjI5ZjU4MTQ5OTM4Zjk2MmQzY2YwN2M0M2QyIiwidXNlcl9pZCI6MX0.wdalVmB7jSka1gWiBCvtxgBqP9jrwNP6Ml9TI8SPsc4" -d "{  \"task\": 1,  \"user\": 2,  \"date\": \"2022-06-23\"}"

Tutorial Test :

    1. Unit test:
       - python manage.py test tests.unit_test.user_detail_task
    2. Intergration test:
       - python manage.py test tests.api_test.user_detail_task

Example Test Was Config in db.sqlite3:
    
    1. User:
       - user_test_1 : {'username':'a', 'password':'1'}
       - user_test_2 : {'username':'b', 'password':'1'}
    2. Intergration test:
       - task_test_1 : {'id':'1', 'name':'cooking'}
       - task_test_2 : {'id':'2', 'name':'booking'}
       - task_test_3 : {'id':'3', 'name':'reading'}
       - task_test_4 : {'id':'4', 'name':'learning'}
       - task_test_5 : {'id':'5', 'name':'bokking'}
       - task_test_6 : {'id':'6', 'name':'playing'}
       - task_test_7 : {'id':'7', 'name':'watching'}
       - task_test_8 : {'id':'8', 'name':'moving'}

###What do you love about your solution?
 - About design:
   - I choose the MVC design for the project because suitable architecture to make code simple, organizable, and maintainable.
   - I make the name file has the same for each the api because we can easy to find all the file need replace.
   - I choose db sqlite for project because it's so simple for mini project and can change another db maybe MySql in settings.py.
   - I choose redis for message queue to coordinator set the limit task for user per day because it's so familiar with me.
   
 - About the Project:
   - The celery will pick task set limit per day for each user.
   - If the limit for user does not set when assign the task celery will set task for user and create the job for set limit task for all another user.
   - About the index on db fit the problem so query in db will fast.
   - The logger will give system's error for debug.log and the togo_task for togo_task_pick_limit_debug.log.
   - I choose the Hard Delete because the task can assign again for user.
 
### What else do you want us to know about however you do not have enough time to complete?
 - About Requirement:
   - In fact, i have many the ideal for resolve the project.
   - The ideal above fit the project need to atomicity, consistency, isolation, durability like the System banking...
   - The Another ideal maybe develop Schedule like:
     - class Schedule:
       - taskmaster: user 
       - task: JSONField
     - task has structure like { "2022-02-02" : ([1,2,3,4,5], 2) }
     - task is [1,2,3,4,5]  
     - the limit task is 2
     - but it doesn't consistency.
   - Maybe consider NoSQL for the ideal above if the project need developed faster.
 - Write the command, input the example for db.