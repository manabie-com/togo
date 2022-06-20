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

##### Current code coverage

File                     | % Stmts | % Branch | % Funcs | % Lines | Uncovered Line #s 
-------------------------|---------|----------|---------|---------|-------------------
All files                |   97.36 |       76 |    87.5 |   97.24 |                   
 togo                    |   90.47 |       50 |   33.33 |   90.47 |                   
  server.js              |   90.47 |       50 |   33.33 |   90.47 | 22,34             
 togo/DB/models          |     100 |      100 |     100 |     100 |                   
  Task.js                |     100 |      100 |     100 |     100 |                   
  User.js                |     100 |      100 |     100 |     100 |                   
 togo/repository         |     100 |    83.33 |     100 |     100 |                   
  Task.js                |     100 |    83.33 |     100 |     100 | 49                
  User.js                |     100 |      100 |     100 |     100 |                   
 togo/router             |     100 |      100 |     100 |     100 |                   
  constants.js           |     100 |      100 |     100 |     100 |                   
  errorConstants.js      |     100 |      100 |     100 |     100 |                   
  index.js               |     100 |      100 |     100 |     100 |                   
 togo/router/controller  |     100 |      100 |     100 |     100 |                   
  task.js                |     100 |      100 |     100 |     100 |                   
 togo/router/middlewares |   95.65 |    63.63 |     100 |   95.23 |                   
  auth.js                |   93.33 |    83.33 |     100 |    92.3 | 34                
  customeError.js        |     100 |       40 |     100 |     100 | 11-14             
 togo/router/routes      |     100 |      100 |     100 |     100 |                   
  task.js                |     100 |      100 |     100 |     100 |                   
 togo/utils              |     100 |      100 |     100 |     100 |                   
  index.js               |     100 |      100 |     100 |     100 |       

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
