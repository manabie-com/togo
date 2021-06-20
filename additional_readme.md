How to run and debug the application:
-Download and install Go: https://golang.org
-Run the command "go run main.go" in the terminal to run the application
-Register as new user if you do not have an account yet
-Get new access token and when it expires with the login API
-Add valid token in the header "Authorization"
-Debug other APIs with their required parameters

What I completed:
-Working limit (5) tasks per day per user
-Proper CRUD (including update, detail, delete tasks and users)
-Additional functions that might come in handy
-Break code into structured model for management
-Refactor some redundant code

What needs further improvement:
-Integrate with Postgresql instead of MySQL
-Proper authorization
-Integrate some sort of query builder to automate the query process
-Use access token to log user out when it expires
-Unit test for functions 
-Swagger document
-Change form submission to request body or header for security
-Better handling of routing in URL
-Improve security of password when dealing with authentication and roles (when developed)