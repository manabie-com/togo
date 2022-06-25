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
    │ ├─┬ togo                         # Directory for Settings system
    │ │ ├─┬ celery                      # Directory for Config celery
    │ │ ├─┬ logger                      # Directory for Logger
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

Run Project:

    1. pip install virtualenv
    2. python -m virtualenv env
    3. source env/bin/activate
    4. pip install -r requirements.txt
    5. python manage.py runserver

Run Project By Docker:

    mkdir logging
    mkdir public
    mkdir public/media
    mkdir public/static
    mkdir public/staticfiles
    docker-compose up

Tutorial Curl :
    
    1. First you need call api to get Token:
       - curl -X POST "http://localhost:8000/api/users/token/" -H  "accept: application/json" -H  "Content-Type: application/json" -H  "X-CSRFToken: O8AESU8podSIZYorLrwc162N7ZdvVkdZpmouAROiTo9BzOSeGjRcBq9oljKG58i8" -d "{  \"username\": \"a\",  \"password\": \"1\"}"
    2 .Second you change token below after beer:
       - curl -X POST "http://localhost:8000/api/users/assignment/" -H  "accept: application/json" -H  "Content-Type: application/json" -H  "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiZXhwIjoxNjU2MTI5MDA3LCJpYXQiOjE2NTYxMjg3MDcsImp0aSI6IjdkMDBlNGM3YjBjMjRiYjk4M2Y1MjdiY2E0ZmM5OWY3IiwidXNlcl9pZCI6MX0.83ItP5Jr2E7zXrvwtEqUG4jCaSkM-LwBve8mDl2hDD0" -d "{  \"task\": 1,  \"user\": 2,  \"date\": \"2022-06-23\"}"