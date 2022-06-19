### Requirements

- NPM >= 8.11.0
- Nodejs >= v16.15.1

### Install

```bash
npm i
```

### Run (port: 3000)

```bash
npm start
```

### Test

```bash
npm test
```

### Demo
#### Accounts
- user A
  - access_token: userA:password
  - limit: 5
- user B
  - access_token: userA:password
  - limit: 10
- user C
  - access_token: userC:password
  - limit: 0

#### Get list task of user

```bash
curl --cookie "access_token=[access_token]" http://localhost:3000/api/tasks
```

#### Add task
```bash
curl --cookie "access_token=[access_token]" http://localhost:3000/api/task -H 'Content-Type: application/json' -d '{"content":"task content"}'
```
