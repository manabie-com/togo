WorkDirectory:


    ├─┬ togo
    │ ├─┬ apps                        # Directory for api using rest-api
    │ │ ├─┬ exceptions                  # Directory for exceptions
    │ │ ├─┬ models                      # Directory for all table
    │ │ ├─├──┬ migrations               # Directory for migrations
    │ │ ├─├──┬ models                   # Directory for Models
    │ │ ├─┬ urls                        # Directory for Controllers
    │ │ ├─┬ views                       # Directory for Views
    │ │ ├─┬ serializers                 # Directory for Manage request and response api
    │ ├─┬ logging                       # Directory for manage logs
    │ ├─┬ test                          # Directory for Unit test
    │ ├─┬ toogo                         # Directory for Settings system
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