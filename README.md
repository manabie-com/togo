# Requirements
1. Docker
2. NodeJs 14
# How to run?
1. Clone the repo `git@github.com:kiettanle/togo.git` or <https://github.com/kiettanle/togo>
2. Open a new terminal at the project's root directory, then run the command to install dependencies:
```
npm install
```
3. Open a new terminal at the project's root directory, then run the command:
```
docker-compose up
```
4. Open a new terminal at the project's root directory, then run this command to create the database:
```
docker exec -it mssql opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P Init123456 -Q "IF NOT EXISTS(SELECT * FROM sys.databases WHERE name = 'Todo')
BEGIN
    CREATE DATABASE [Todo]
END"
```
5. The run command or copy .env.example to .env
```
cp .env.example .env
```
6. Then run the migration
```
npm run migration:up
```
7. Seed sample data via:
```
npm run migration:seed
```

# How to Test API?
Open a new terminal at the project's root directory, then run the command:
```
npm run start:dev
```
## Swagger URL
<http://localhost:3000/_swagger/>

Test account:
```
username: user1
password: Init123456
```
```
username: user2
password: Init123456
```

## Postman collection
Check folder `postman`
# Sample CURL
Before running the pick task API, the user must log in first:
```
curl --location --request POST 'http://localhost:3000/auth/token' \
--header 'Content-Type: application/json' \
--data-raw '{
  "grant_type": "password",
  "username": "user1",
  "password": "Init123456",
  "rememberMe": true
}'
```

The result will look like this, copy your access_token and paste it into 
```
{
  "expires_in": 2592000,
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwicm9sZSI6IlVzZXIiLCJwZXJtaXNzaW9ucyI6WyJ0YXNrc19waWNrIiwidGFza3NfbGlzdCJdLCJ1c2VySWQiOiI0QUJDNDMzRS1GMzZCLTE0MTAtODg4Ny0wMDM3Q0E2RTZGNDIiLCJpYXQiOjE2NTM3MDU3ODksImV4cCI6MTY1NjI5Nzc4OSwiYXVkIjoiVG9kb0FwcCIsImlzcyI6ImxvY2FsaG9zdCIsInN1YiI6IjRBQkM0MzNFLUYzNkItMTQxMC04ODg3LTAwMzdDQTZFNkY0MiIsImp0aSI6InhXaU1vN082In0.r7z67EnkvQ68dItQsw7gmevPuBpRO_Aljd85em0N0Yg",
  "refresh_token": "OrWeiBfkg0DSWrOugjamQIql25arSgJ0yNLHsJQMTBGTn8JyYxlA33EaJR0ifmma",
  "token_type": "Bearer"
}
```
## Call pick task API, `taskId` you can copy from the database or call `GET /tasks` API to see the task list
```
curl --location --request POST 'http://localhost:3000/tasks/pick' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {{access_token}}' \
--data-raw '{
    "taskId": "4A86433E-F36B-1410-86BA-002BD86783D2"
}'
```

## API to get task list
```

curl --location --request GET 'http://localhost:3000/tasks' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {{access_token}}'
```
## Check your tasks list after picked:
```
curl --location --request GET 'http://localhost:3000/my-tasks' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {{access_token}}'
```
# How to run test?
* To run unit test only
```
npm run test:ut
```

* To run unit test only in watch mode
```
npm run test:ut:watch
```

* To run integration test only
```
npm run test:it
```

* To run integration test only in watch mode
```
npm run test:it:watch
```

* To run both unit test and integration test
```
npm run test
```

* To run test coverage
```
npm run test:cov
```

# To remove the docker container, run docker-compose down
```
docker-compose down
```

# Conclusion
## What do I love about my solution?
> - I worked with bare nodejs and it did not have any standard for structure, I can code as what I want, no standard, no structure. Ok fine,  it still run. But it will be a nightmare when the project is bigger and bigger and have many people join to that project.
 > - And when I worked with NestJs, I followed it's structure and feel it is a good Nodejs framework. We built the each module separated then merge to AppModule.
 > - In this project I used JWT to authenticate user, used typeorm to migrate and seed database, used swagger to create API docs. I also tried to create github action to run build and run test.

## What else do I want you to know?
> Write more test cases to increase test coverage.
> Write end to end test case.
> Update swagger for more detail include the request and response sample.
> Do performance testing using JMeter.
> Create script to build and push docker image to docker hub.
> Create script to deploy project to AWS or Digital Ocean or Self server.

---
_*I spent 3 days for this assignment, I have 2 days left. But I think it is enough for this project. I used to work like a machine and forgot my family. My health was down. I don't want it happen again. I should balance between my work and family. So I decided submit what I have done.*_

_I heard your company is using Golang. I have zero about Golang. But as my opinion learn the language is not difficult. The language can different syntax, code style... something else. If I can passed and I can learn with my passion for new technology and good instruction from your company. I think I can control Golang in the short time. But to be a `Golang monster` I think I need to work for many years. I hope to see you in the face to face round. Thank you for your attention._