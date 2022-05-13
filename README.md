### How to run code locally?

- Environments
  - Docker
- Run code
  - Clone repo: git clone https://github.com/long-lehoang/togo
  - Open togo directory: cd togo
  - Run app: docker compose up -d
  
### A sample “curl” command to call API

- Register user: curl -X POST -H "Content-Type: application/json" -d "{ \\"username\\":\\"long\\",\\"password\\":\\"Password1!\\"}" "http://localhost:8080/user/register"
- Login: curl -X POST -H "Content-Type: application/json" -d "{ \\"username\\":\\"long\\",\\"password\\":\\"Password1!\\"}" "http://localhost:8080/auth/login"
- Add Task: curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer {*Token*}" -d "{ \\"title\\":\\"Interview\\",\\"description\\":\\"Manabie interview\\"}" "http://localhost:8080/task"
- Delete task: curl -X DELETE -H "Content-Type: application/json" -H "Authorization: Bearer {*Token*} " -d "{ \\"title\\":\\"Interview\\",\\"description\\":\\"Manabie interview\\"}" "http://localhost:8080/task?id={id}"

### How to run your unit tests locally?

- Just run command in todo directory: docker compose run test

### What do you love about your solution?

- Nothing

### What else do you want us to know about however you do not have enough time to complete?

- Add Flyway to manage script
- Add Update Task API
- Update Delete Api with isDelete.