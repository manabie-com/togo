### Overview
- Togo app helps we can add the tasks and assign for user easily. See the steps below to get started

### How to run
- Run on the local:
```console
cd togo
npm install
npm run start
```
- You can test APIs using `curl`:
  - First, you should create a user 
```cosole
curl -d "{\"username\":\"adminTest\",\"password\":\"adminTest\"}" -H "Content-Type: application/json" -X POST "http://localhost:3000/api/users/register"
```
  - Login user
```cosole
curl -d "{\"username\":\"adminTest\",\"password\":\"adminTest\"}" -H "Content-Type: application/json" -X POST "http://localhost:3000/api/users/register"
```
  - After logined, you will get a token as below
```json
{
  "user": {
    "_id": "62998a3350c907ff52c228ad",
    "username": "admintest",
    "limit": 0,
  },
  "token": "JWT-eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWludGVzdCIsImlhdCI6MTY1NDIyOTc5MiwiZXhwIjoxNjYyMDA1NzkyfQ.LGiVR2GS2Blh3BeED44uMTOYn8q8eiOEaOs43nZnYjo"
}
```
   -  You have to add token for authentication once call APIs
```cosole
curl -d "{\"name\":\"This is the task sample\",\"username\":\"adminTest\"}" -H "Content-Type: application/json" -H "accept: */*" -H "Authorization: JWT-eyJhbGciOiJIUzI1NiIs...." -X POST "http://localhost:3000/api/tasks"
```
- Run the testcases:
```cosole
npm run test
```
Actually, TDD is a new style for me. It gives me very interesting experiences, gives me more perspectives. TTD helps me feel more confident when coding.

For this project. I feel it still lacks some things. For example, there should be logging, paging. Also need Docker for ci/cd, Swagger for document apis easier to understand.

Thank for your spending time to review my project. Hopefully, we will have a face to face meeting to discuss more details about this project , or any in-depth information.