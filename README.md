# Todo Exercise

## Requirement:
````
Docker, Docker-compose
````

## Install
1. Clone git source
2. Run `docker-compose build`
3. Run `docker-compose up -d`


## Run unit test

```commandline
docker exec -it togo_backend_1 bash

python manage.py test authentication.tests todos.tests
```

## API document
https://documenter.getpostman.com/view/5858231/UyxnD4Yv