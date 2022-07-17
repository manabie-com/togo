# Togo
Test API for the Akaru Backend Engineer Coding Challenge that allows a user to create tasks.

## About the project
The Togo API is built using the following:
- Python 3.8.5
- Django 3.2.7
- Django REST Framework 3.12.4
- PostgreSQL 13.0

## Setting up
You may build a Docker container which contains the Django server as well as a Postgres database:
```
docker-compose -f docker-compose.yml up --build
```
Once this is successful, you may access the API at `localhost:8000/api/`.

Alternatively, you may run the project outside of Docker, using `virtualenv`.
```
virtualenv env
source env/bin/activate
pip3 install -r requirements.txt
python3 manage.py runserver
```
Note that this will use SQLite as the default database instead of Postgres.

## Creating an admin user
You may optionally create a superuser so that you can change the task limits for each user. While the Docker container is running, execute the following command in another terminal:
```
docker exec -it <CONTAINER_ID> python manage.py createsuperuser
```
An interactive console will guide you through setting up the credentials of the admin user.

You may obtain the container ID by executing `docker ps` and looking for the ID of the `web-1` container.

## Running the tests
```
docker exec -it <CONTAINER_ID> python manage.py test
```

## Using the API endpoints

You may test the API through the following Postman collection:

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/3110431-083cc236-df41-4cb7-88b0-deaaf6939eff?action=collection%2Ffork&collection-url=entityId%3D3110431-083cc236-df41-4cb7-88b0-deaaf6939eff%26entityType%3Dcollection%26workspaceId%3Dd75aa22d-c1db-40db-a193-1fee1cdf4e3c#?env%5BLocal%5D=W3sia2V5IjoiSE9TVCIsInZhbHVlIjoiaHR0cDovL2xvY2FsaG9zdDo4MDAwIiwiZW5hYmxlZCI6dHJ1ZSwidHlwZSI6ImRlZmF1bHQiLCJzZXNzaW9uVmFsdWUiOiJodHRwOi8vbG9jYWxob3N0OjgwMDAiLCJzZXNzaW9uSW5kZXgiOjB9LHsia2V5IjoiUFJPRCIsInZhbHVlIjoiaHR0cDovL3NncDEuaXZhbmJhbGluZ2l0Lm1lOjgwMDAiLCJlbmFibGVkIjp0cnVlLCJ0eXBlIjoiZGVmYXVsdCIsInNlc3Npb25WYWx1ZSI6Imh0dHA6Ly9zZ3AxLml2YW5iYWxpbmdpdC5tZTo4MDAwIiwic2Vzc2lvbkluZGV4IjoxfSx7ImtleSI6IlRPS0VOIiwidmFsdWUiOiJleUowZVhBaU9pSktWMVFpTENKaGJHY2lPaUpJVXpJMU5pSjkuZXlKMGIydGxibDkwZVhCbElqb2lZV05qWlhOeklpd2laWGh3SWpveE5qVTNPVEV5TkRrMUxDSnBZWFFpT2pFMk5UYzVNVEl4T1RVc0ltcDBhU0k2SWpCaU9XTTVOamRoTm1ZelpUUTJPV0k1WldOa1pHTm1aR1E0WTJNNU4yTmtJaXdpZFhObGNsOXBaQ0k2TVgwLllnSW00Vi1rVE5ocG9NQlQ3ak1oakFOUk83b3BGMEl6XzNvdkhwVEtTZ1UiLCJlbmFibGVkIjp0cnVlLCJzZXNzaW9uVmFsdWUiOiJleUowZVhBaU9pSktWMVFpTENKaGJHY2lPaUpJVXpJMU5pSjkuZXlKMGIydGxibDkwZVhCbElqb2lZV05qWlhOeklpd2laWGh3SWpveE5qVTRNVFl4Tmpnd0xDSnBZWFFpT2pFMk5UZ3dOelV5T0RBc0ltcDBhU0k2SW1KbE5EQTFOelE1T1ROaS4uLiIsInNlc3Npb25JbmRleCI6Mn0seyJrZXkiOiJUQVNLX0lEIiwidmFsdWUiOiIxIiwiZW5hYmxlZCI6dHJ1ZSwic2Vzc2lvblZhbHVlIjoiMSIsInNlc3Npb25JbmRleCI6M31d)

The API is also deployed live on my personal server at `http://sgp1.ivanbalingit.me:8000/api/`.

### User Signup ```POST /api/users/```
```
curl -X POST http://localhost:8000/api/users/ \
  -H "Content-Type: application/json" \
  -d '{"username": "your_username", "email": "email@email.com", "password": "password"}'
```

### User Login ```POST /api/token/```
```
curl -X POST http://localhost:8000/api/token/ \
  -H "Content-Type: application/json" \
  -d '{"username": "your_username", "password": "password"}'
```
This returns something like ```{"refresh": "REFRESH_TOKEN_HERE", "access": "ACCESS_TOKEN_HERE"}```. You may use the access token for endpoints requiring authentication.
The token lasts for 1 day. You can get a new access token through supplying the refresh token to ```/api/token/refresh/```.

### Get Tasks for User ```GET /api/tasks/```
```
curl -L -X GET http://localhost:8000/api/tasks/ \
  -H "Authorization: Bearer ACCESS_TOKEN_HERE"
```

### Create Task for User ```POST /api/tasks/```
```
curl -X POST http://localhost:8000/api/tasks/ \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ACCESS_TOKEN_HERE" \
  -d '{"name": "Create tasks API"}'
```

### Update Task for User ```PUT /api/tasks/<TASK_ID>/```
```
curl -X PUT http://localhost:8000/api/tasks/<TASK_ID>/ \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ACCESS_TOKEN_HERE" \
  -d '{"name": "Create tasks API and push to repo"}'
```

### Delete Task for User ```DELETE /api/tasks/<TASK_ID>/```
```
curl -L -X DELETE http://localhost:8000/api/tasks/<TASK_ID>/ \
  -H "Authorization: Bearer ACCESS_TOKEN_HERE"
```

## Notes
- I have chosen to use Python/Django REST Framework since it is the platform I am most comfortable with. I have created the structure of the project so that it is simple to understand from the point of view of another person. Since the API has simple requirements, a complex file structure is not needed right now.
- Given more time, the following could have been implemented:
  - Modularize the code for the viewsets, serializers, models, and tests so that each component of the project (e.g. User, Task) has its own file for better maintainability.
  - Write further tests that will catch possible edge cases when using the API (e.g. deleting a task that does not belong to you).
  - Deploy the Docker container in AWS Lightsail instead of my own server.
  - Utilize GitHub Actions more for continuous integration.
