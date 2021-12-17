# Manabie Application Assignment
## Simple API for creating new daily tasks
The detailed description of the assignment can be found [here](https://github.com/manabie-com/togo)  
### How to run the app locally
#### Prerequisites:
- NodeJS
- 
In your terminal, to install all the required dependencies, run:
```bash
npm install
```
#### Schemas
Table `users` to store user information.
```sql
CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  `name` VARCHAR,
  email VARCHAR NOT NULL,
  created_at TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
)
```
Table `tasks` to store daily created tasks.
```sql
CREATE TABLE IF NOT EXISTS tasks (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title VARCHAR NOT NULL,
  detail VARCHAR,
  due_at TIMESTAMP,
  created_at TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  reporter_id INTEGER,
  assignee_id INTEGER,
  FOREIGN KEY (reporter_id)
    REFERENCES users (id),
  FOREIGN KEY (assignee_id)
    REFERENCES users (id)
)
```
Table `user_tasks` to keep track of the number of tasks a user has created per day.
```sql
CREATE TABLE IF NOT EXISTS user_tasks (
    created_date TEXT NOT NULL,
    reporter_id INTEGER NOT NULL,
    task_count INTEGER NOT NULL
)
```
#### Start the server
In your terminal, run:\
```
npm start
```
#### Create a user
In your terminal, to create a user id, run:\
On Windows:
```bash
curl --header "content-type: application/json" --request POST --data "{\"name\": \"frank\", \"email\": \"frank@mail.com\"}" http://localhost:3000/users
```
On Mac:
```bash
curl --header "content-type: application/json" --request POST --data '{"name": "frank", "email": "frank@mail.com"}' http://localhost:3000/users
```
Expected output:
```json
{"id": 1}
```
You have successfully created a user with id (auto generated), name and email.
#### Add a new task
In your terminal, to add a new task, run:\
On Windows:
```console
curl --header "content-type: application/json" --request POST --data "{\"title\": \"get grocery\", \"detail\": \"buy eggs and ham\", \"due_at\": \"2021-12-31 23:59:59\", \"reporter_id\": 1}" http://localhost:3000/tasks
```
On Mac:
```console
curl --header "content-type: application/json" --request POST --data '{"title": "get grocery", "detail": "buy eggs and ham", "due_at": "2021-12-31 23:59:59", "reporter_id": 1}' http://localhost:3000/tasks
```
Expected output:
```json
{"id": 1}
```
To see the task count:\
```console
curl http://localhost:3000/tasks/count | json
```
The `| json` tag is there to make the output more readable. The output should be:
```json
{
  "results": [
    {
      "created_date": "2021-12-17",
      "reporter_id": 1,
      "task_count": 1
    }
  ]
}
```
#### When you have reached your daily task limit, which is preset to `5`
