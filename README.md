## Manabie Togo app

### Notes
This is a simple backend for a todo service, right now this service can handle login/list/create simple tasks, to make it run:
- `java -jar togo-app-1.0.jar`
- Import Postman collection from docs to check example

### Functions I have done 
- The Togo application is written by Java + Spring framework.
- The main difference between my version and orginal version is: my app is written by 100% Java, it is nicer, cleaner, the layers are more seperater, easier to maintain, and reuse.
- Authenticate / Authorization JWT 
- Add a new JWT filter chain to handle everything about Authenticate / Authorization JWT, and the REST API Controller does not worry any thing about the security (better than the original Golang source code, the security is mixing with the Controller).
- The /login url is allowed for everyone, but others should authenticated (login first, then make a call later).
- The user contains the **encrypt** password (not the raw text as the orginal).
- Is is only support the Postgress connection (server: localhost / port 5432 / username postgres / password 123456), and does not support SQLite.
- Login request: get username/password, then return the JWT token to user (I sign it by SignatureAlgorithm.HS512, with SecretKey). 
- List task request: get created date, then return the tasks at created_date (must pass login first, and has Authorization header).
- Add task request: get content, then if under limit, create a new task and add to database / if over limit, throw HTTP Error 429 Too many request.
- Some unit tests (but I do not enoush time to finish everything).
- Fork the orginal repo.
- The source code is commented, and easy to read and understand.

### Functions I do not yet done
- Not yet finsih do unit tests.
- Not yet write docker-compose.
- I think the DB structure contain some small issues, but in practically, it can works, and can solve what I want.
