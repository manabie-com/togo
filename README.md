### Requirements

- Java 11
- Maven

### Unit test

Run `mvn test` to execute unit test with Junit

### Build

Run `mvn clean install` to build a jar file

### Run

Run `mvn spring-boot:run` for a dev server

### CURL commands

- Create an account by `curl --location --request GET 'http://localhost:8080/api/accounts' --header 'Authorization: Bearer ${token}--header 'Content-Type: application/json' --data-raw '{"firstName": "tu","lastName": "le","password": "letu","uid": "letu","taskLimit": 10}'`. Remember that only admin user (username: admin, password: admin) can triger this endpoint

- Create a task by `curl --location --request GET 'http://localhost:8080/api/tasks/2' --header 'Authorization: Bearer ${token} --header 'Content-Type: application/json' --data-raw '{"title": "Create project plan","notes": "This can be done tomorow"}'`. To trigger this endpoint, regular user (the account that created above) should be used (username: letu, password: letu)

### Login 

Note that `${token} is Bearer token that received when a user login into the application. Information to get a token: 

- Callback url: http://localhost:8080

- Auth url: http://localhost:8080/oauth/authorize

- Access token url: http://localhost:8080/oauth/token

- Client id: myclient1

- Client secret: mysecret1

- Scope: user_info