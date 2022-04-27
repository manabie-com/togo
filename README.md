Run docker-compose up for start app

go to link swagger: localhost:3002/v1

User login: 
    + username: user1
    + password: Test@123

curl : curl -X 'POST' \
  'http://localhost:3002/api/tasks' \
  -H 'accept: */*' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6InVzZXIxIiwidXNlcm5hbWUiOiJ1c2VyMSIsImNyZWF0ZWRBdCI6IjIwMjItMDQtMjdUMTM6MjE6NTEuNDA1WiIsInVwZGF0ZWRBdCI6IjIwMjItMDQtMjdUMTM6MjE6NTEuNDA1WiIsImlhdCI6MTY1MTA2NTg5NSwiZXhwIjoxNjUxMTUyMjk1fQ.JM-ibJrPqCPNf33zOl-8S4cuaa13J8bBKwu9baT3YnE' \
  -H 'Content-Type: application/json' \
  -d '{
  "tasks": [
    {
      "title": "test1",
      "description": "testing",
      "note": "test",
      "watchers": [
        "1"
      ],
      "excutors": [
        "2"
      ]
    }
  ]
}'