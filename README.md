### How to run code locally?

- Enviroments
  - Apache Maven 3.8.5
  - Openjdk 17.0.2
  - Docker
- Run code
  - Clone repo: git clone https://github.com/long-lehoang/togo
  - Open togo directory: cd togo
  - Run mysql on docker: docker-compose up -d
  - Build project: mvn clean install
  - Run: mvn spring-boot:run
  
### A sample “curl” command to call API

- Register user: curl -X POST -H "Content-Type: application/json" -d "{ \\"username\\":\\"long\\",\\"password\\":\\"Password1!\\"}" "http://localhost:8080/user/register"
- Login: curl -X POST -H "Content-Type: application/json" -d "{ \\"username\\":\\"long\\",\\"password\\":\\"Password1!\\"}" "http://localhost:8080/auth/login"
- Add Task: curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer {*Token*}" -d "{ \\"title\\":\\"Interview\\",\\"description\\":\\"Manabie interview\\"}" "http://localhost:8080/task"
- Delete task: curl -X DELETE -H "Content-Type: application/json" -H "Authorization: Bearer {*Token*} " -d "{ \\"title\\":\\"Interview\\",\\"description\\":\\"Manabie interview\\"}" "http://localhost:8080/task?id={id}"

### How to run your unit tests locally?

- Just run command in todo directory: mvn test

### What do you love about your solution?

- Nothing

### What else do you want us to know about however you do not have enough time to complete?

- Add Flyway to manage script
- Add Update Task API
- Update Delete Api with isDelete.