# Run code locally ubuntu

- Run this code is easier with pycharm IDE

- Install docker to run postgresql
>$ sudo apt-get install docker-ce docker-ce-cli containerd.io
- Check docker is installed yet
>$ docker --version
- Run postgresql
>$ sudo docker run -d --name postgres -e POSTGRES_HOST_AUTH_METHOD='trust' -e POSTGRES_PASSWORD='1234' -e POSTGRES_DB= -p 5432:5432 postgres
- Check docker container is running
>$ sudo docker ps -a
- Install python 3.8
>$ sudo apt update
>$ sudo apt install python3.8 python3.8-dev python3.8-venv
- CD into the project
- Create virtual environment:
>$ python3.8 -m venv venv
- Activate virtual environment:
>$ source ./venv/bin/activate
- Install required libraries:
>$ cd <-path-to-project->
> 
>$ pip install -r requirements.txt 
- To start the code run:
>$ python app.py
```
  database loading...
  database loaded
   * Serving Flask app 'app' (lazy loading)
   * Environment: production
     WARNING: This is a development server. Do not use it in a production deployment.
     Use a production WSGI server instead.
   * Debug mode: on
   * Restarting with stat
   * Debugger is active!
   * Debugger PIN: 316-686-129
   * Running on http://127.0.0.1:5000/ (Press CTRL+C to quit)
```
- Verify that you run successfully at <http://localhost:5000>
- Can use file data_sample.json to import the sample data into database with curl:
>$ cd <-path-to-project>
> 
>$ curl -X POST http://localhost:5000/api/api/demo
>   -H 'Content-Type: application/json'
>   -d @data_sample.json

## A sample “curl” command to call your API

>$ curl -X POST http://localhost:5000/api/task/add
>   -H 'Content-Type: application/json'
>   -d '{
	"title": <-title->,
	"description": <-description->,
	"user_id": <-user_id->,
	"date": <-date->"
}'

## How to run your unit tests locally
- Install requirement
> $ cd test

> $ pip install -r requirements-test.txt

- From project run command
> $ coverage run -m pytest .\test\
- Report the coverage of test on the code
> $ coverage report -m

## Error
- Not matching distribution found for psycopg2
  - open requirements.txt, find and replace psycopg2 with psycopg2-binary

## What do you love about your solution?
- I think this project should be upgraded to be better.
- In this case, I think the database can easy to add more queries if we need without writing the raw SQL query.
- Python is a good support with many libraries

## What else do you want us to know about however you do not have enough time to complete?
- I think should add swagger to this assignment in order to easier to use the API.
- I think should add a formatting check as the black library
- I think should add an API folder to route more professional
- I think should add a logger to check and write the log for the trace error process
