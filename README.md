### How to run your code locally?

- B1. Clone project
- B2. go into folder challenge\target
- B3. open command line screen at that folder
- B4. run command to execute app java
   - command: java -jar challenge-0.0.1-SNAPSHOT.jar
- B5 use command curl the other cmd screen to test

### A sample “curl” command to call your API

- curl -i -X POST localhost:8080/api/login -H "Content-Type: application/json" -d "{\\"username\\":\\"hungnk\\",\\"password\\":\\"admin123\\"}"
- As soon as run this command it will generate Jwt Token
- assign to "Authorization": "Bearer {#token}" to api createTask, createUser called after login;
  - Ex call api after login: 
    - curl -i -X POST localhost:8080/api/createTask -H "Content-Type: application/json" -H "Authorization: Bearer {#token}" -d "{\\"id\\":\\"M01\\",\\"content\\":\\"doning01\\",\\"userId\\":\\"hungnk\\",\\"createdDate\\":\\"2022-06-24\\" }"
    - {#token} is token received at the time login
    - create user with the following curl: 
	- curl -i -X POST localhost:8080/api/createUser -H "Content-Type: application/json" -H "Authorization: Bearer {#token}" -d "{\\"username\\":\\"trungst\\",\\"password\\":\\"admin1234\\",\\"maxLimitTodo\\": 2}"
    - You can check the database at: http://localhost:8080/h2-console
	   - JDBC URL: jdbc:h2:mem:chanllengesdb
	   - User Name: sa
	   - no Password

### How to run your unit tests locally?

- Currently, I can't run unit tests in cmd, if you want to run then open java IDE like Eclipse, Intellij... I apologize for the inconvenience.

### What do you love about your solution?

- Currently this project I am using MVC pattern, it divides the structure very clearly so it is easy to change and maintain.

### What else do you want us to know about however you do not have enough time to complete?

   - I still can't configure unit test to run on cmd, I feel pretty bad for this.
