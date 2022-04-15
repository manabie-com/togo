# Đỗ Hoài Diễn
## Requirement
- NodeJS v14 
- Npm/Yarn
- Docker / MySql server
## Installation
Install package requirement
```sh
npm install
```
Start Docker or run MySql Server
Mysql info:
- Host: 127.0.0.1
- Username: root
- Password: 1
```sh
docker-compose up -d
```
Create database for your application
```sh
CREATE DATABASE togo CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```
Run NestJS in your local
```sh
npm run start:dev
```
Migrate Admin role
```sh
npm run migrate
```
Admin info: 
- Email: admin@admin.com
- Password: 1

Open Swagger in your browser
```sh
http://localhost:3000/api
```
## Testing
Run Unit test & Integration test
```sh
npm run test
```
## Swagger
You can testing with Swagger
Swagger link
```sh
http://localhost:3000/api
```
## Solution for this task
Because of user only has maximum limit task can be added per day and diffrent users can have different maximum daily limit, I will create role for user, each of users register has default max task can add per day (3) and admin can modify it.
To prevent race condition(user can add more then their maximum task per date), I designed 
<br />
Column `task_left`(number of task left user can add) with type: `unsigned` and decrease it before add task in table task 
<br />
Column: `lasted_date_task`: record the last task date
## Improvement
Because not enough time, this Back-end have some missing:
- Error exception is not detailed
- Test case is not detailed
- Missing automation testing
## Online server
If you have some error when start this app, you can use my server
```sh
http://54.179.70.143:3000/api
```
## Demo video
```sh
https://drive.google.com/file/d/1CKkprpdQyYroXvRMQ7OtroOtO7kj1WTi/view?usp=sharing
```
