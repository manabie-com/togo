The application is restructured as a Django application from the original Go Lang codebase.

 **Setting-up Django**
Use the virtual environment provided. In the root directory, execute:
-  `source env/bin/activate`
-  `pip install -r requirements.txt`

This should take care of the dependencies needed to run the Django application.

**Running PostgreSQL in Docker**
In the root directory, a Makefile is provided to setup the docker, execute:
-  `make postgres`
-  `make createdb`
  
**Migrating Data**
Go inside the `todo` directory and execute the following in sequence:
-  `python manage.py makemigrations`
-  `python manage.py migrate`
-  `python manage.py loaddata todo/fixtures/init_data.json`

The commands above will create a schema based on the defined models in `todo/models.py` and load initial data for testing from the `fixture` folder.

**Running the Server**
Finally, to run the server, inside the todo folder, execute:
-  `python manage.py runserver 5050`

By default, it will run on port `8000`, but for the ease of testing, please specify the port number as `5050`.

**Running the Test Scripts**
I have few test scripts prepared inside the `todo/tests` folder. To run the tests, inside the todo folder, execute:
-  `python manage.py test`

**Sample API tests**
I modified the login to use POST method instead of GET, I added my revised postman collection on the root directory of the project.
-  POST `localhost:8000/login` payload: `{"user_id":"firstUser","password":"example"}`
-  GET `localhost:8000/tasks`
-  POST `localhost:8000/tasks` payload: `{"content":"sample content"}`


### Functional requirement:

Upon adding of new todo, the system now checks if the maximum todo of the user for the day is reached. If the user already created N todo, the system will return a 400 error.

N todo is defined on the max_todo column in the user table.
  

### Non-functional requirements:

I did most part of the non-functional requirement checklist, but here are the things I wished I can improve on my work:

 - I wish I could explore more on creating testing. As seen on my commits, I did the test on the later part, and obviously not in a TDD manner. This is something that I am working on and would like to improve on.
 - Django has a build in module for authentication, but I decided to do it from scratch. With that I made security compromises. Given more time, I think I can play around authentication and session management more.
 - On PostgreSQL and docker, I just pulled the Postgres image and create a container out of it. I provided the command on the Makefile. I could have done it as docker-compose.
 - I noticed that the API fails in Django due to trailing backslash issues, so I defined both urls with and without backslash, but I think there are better ways to fix this.

Some of the modifications I did from the original code that I this are reasonable:
 
 - I changed the login from GET request to POST request and put the data in the request body.
 - I added an incremental integer key on users. Using text as key can be confusing and prone to conflict especially if it will be coming from the user's input.
 - I changed some data types on the schema.
 - I restructured the handler from the original Go lang codebase to how I'll do it in Django.