# TODO API
A simple API for tracking TODO records written in Golang (Gin, Gorm, Testify).
When a TODO request is post to API with an user id which hasn't been registered in the database, the API will create a new user with that ID along with it's TODO task.
If user id is exist in database and user hasn't reached daily limit (default 8), the TODO task will be registered.

# How to run
To run the project, first install the dependencies
> go mod tidy

Configurate database's information in **.env** file copied from **.env.example**. The database should be created in order for the API to connect and migrate
```
file .env
DB_HOSTNAME=localhost
DB_USERNAME=root
DB_PASSWORD=
DB_PORT=3306
DB_NAME=todo_db
```
Run command start API
> go run main.go

If the database configuration is correct and the API is able to connect, migration will be automatically run.

# Example cURL command
```
curl --location --request POST "http://localhost:8080/api/to-do" \
--header "Content-Type: application/x-www-form-urlencoded" \
--data-urlencode "user_id=1" \
--data-urlencode "task_detail=Hello World TODO"
```
This command will create a post request to `localhost:8080/api/to-do` registering for user id = 1 with TODO task is `Hello World TODO`

# How to test
When testing we will use a test database, i.e `todo_db_test` to perform database operations. Please specify the DB configuration in the `.env` came along within each test package in order for the test to work.
## Unit tests
To run a unit test, use the example command below:
```
go test <path_to_test_package> -count=1
i.e
go test ./test/unit/form -count=1
```
Flag `-count=1` is to prevent test caching

## Functional tests
To run functional test, use the example command below:
```
go test <path_to_test_package> -count=1
i.e
go test ./test/functional -count=1
```

# What I love about this solution
This is the most challenging assignment I have ever been on. Through out the process of creating, structuring and testing this API, I have learn so much more about Golang, struct, interfaces, packages, testing, concurrencies,... I am aware that I'm allowed to use other languages which I'm familiar with to give a solution to this assignment, but still I prefer stepping out of comfort zone and prove that I am eager to learn new things.
I love most in this project is the `validator` package which I implemented as a global package so that I can use wherever to validate my structs, and also the `ErrorJSON` struct that can be used as json to display message (before I was using a single string and tried to display it as JSON, not very wise).
Then the best thing in this project is testing, I was able to write tests which executed seamlessly, even though I have zero experienced in TDD or any other kind of unit/integration/functional testings.

# One more thing
My tests run fine when conducted seperately (run command on each of testing package), but will fail when it come to run in all tests `go test ./test/unit/...`, this is because I stil haven't figured it out what is the best way to structuring test packages and how they should be coded. It would be great if you can advice me on how to make my TODO project better and how should I code testing.
Lastly, thank you for letting me join this assignment, I enjoyed it.
