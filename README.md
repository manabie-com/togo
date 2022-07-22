<div align="center">
	<h1>Togo</h1>
	<p>A project challenge from manabie.</p>
	<a href="/CHALLENGE_INSTRUCTIONS.md">Original Instructions can be found here</a>
</div>

### Architecture
- Hexagonal Architecture (a.k.a Ports and Adapters)

### How to run locally

- clone this repository
```bash
git clone git@github.com:vindecodex/togo.git
```

- `npm install`

- copy env variables
```bash
cp .env-sample .env
```

**Note: (FOR NONE DOCKER ONLY) Please create 'togo' database before running the program**

- npm run dev

### Run using Docker

**NOTE: Also close other running mysql server from your computer it will conflict with the mysql docker image**

```bash
docker-compose up --build
```

### Run test
`npm run test` or `npm t`

### Usage

#### Check If API runs correctly
```bash
curl http://localhost:3000
```
Should output **Welcome Togo**

---

#### Create Account - POST Request
You cannot create a task If you don't have an account.
```bash
curl -d 'username=test&password=test' http://localhost:3000/user
```
Should output `{"id":1,"username":"test","type":"BASIC"}`

---

#### Login - POST Request
Get a token to allow creating task
```bash
curl -d 'username=test&password=test' http://localhost:3000/authenticate
```
Should output
```bash
{
"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ0ZXN0IiwidHlwZSI6IkJBU0lDIiwiaWF0IjoxNjU4MzkyOTk1fQ.6LOrSehIyFGhYRULxOYvm2o1jLVFFzv3gkHSaZqYfBM"
}
```
Copy Token value to be used for posting task

---

#### Create Task - POST Request
**NOTE:** Newly created user has only BASIC account which means you can only post 5 task per day, you can subscribe to PREMIUM to have 10 task per day.
```bash
curl -d 'task=YOUR_TASK' -H "Authorization: Bearer TOKEN_HERE" http://localhost:3000/todo
```
Should say true if added successfully

---

#### Upgrade to premium - POST Request
**NOTE:** You need to re-authenticate after upgrading.
```bash
curl -H "Authorization: Bearer TOKEN_HERE" http://localhost:3000/upgrade
```
Should say true if upgrade success.

---

#### Get All Todo - GET Request
```bash
curl http://localhost:3000/todo
```

---

### Get All Todo By User Only - GET Request
```bash
curl http://localhost:3000/todo/user/{userId}
```
