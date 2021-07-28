### Notes
- This is my first project with Golang so it's quite challenging to learn both syntax and framework. But I don't want to
use Java Spring Boot framework to have a chance to quickly learn this new language.
- I added auto generation of UUID for id column in database instead of generating from the code upon inserting.
- Handlers (API endpoints) are separated to `handlers` package to only handle things related to request, response, endpoints registration, etc.
- `db.go` now is the interface to help switching to other database easier. Only need to have a method to initialize and to implement remaining methods
- I changed "sign in" to POST request with username and password in body to enhance security
- I believe that Unit Test is a very good way to ensure each usecase in each component is working as expected. However, I am quite busy
with my work recently and don't have time to learn about Unit Testing in Go and write the test cases. I rely on integration testing
through Postman which can cover the actual flow of customer to ensure the app is working.
- If I have more time, I will do:
	- Unit Tests
	- hasing `password` to ensure password is not stored as plain text in database
	- using Viper for loading the configuration instead of hardcoding the values inside the app
	- add more APIs to follow RESTful standard
		- `GET /users/:id` to get `users` by `id`
		- `POST /users` to create new `users`
		- `PUT /users/:id` to update username, password, max_todo for a specific `users` 
		- `DELETE /users/:id` to delete a specific `users`
		- `GET /tasks/:id` to get `tasks` by `id`
		- `POST /tasks` to create new `tasks`
		- `PUT /tasks/:id` to update a specific `tasks`
		- `DELETE /tasks/:id` to delete a specific `tasks`
	- dockerize the whole

### Steps
This is a simple backend for a todo service, right now this service can handle login/list/create simple tasks, to make it run:
- ensure the `pg` folder and the migration script inside is accessible by `root` user so psql docker can run during setup
- `sudo docker-compose up`
-  go to http://localhost:5051/login with username `admin@manabie.com` and password `admin`
	- setup new server: hostname: `db`, port: `5432`, username: `postgres`, password `postgres`
	- if there is no table setup yet, run the migration script `migration_script.sql` in `pg` folder
- run the application using `go run main.go`
- Postman collections can be imported to run integration testing through two files using Postman Runner
	- collection file: togo.postman_collection.json
	- environment file: local.postman_environment.json