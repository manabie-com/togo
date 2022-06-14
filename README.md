### Requirements

1. MySQL Server, create a database called `todo_db`.
2. [Node.js](https://nodejs.org/en/download/) installation. You can use the latest LTS version as of this writing (16.15.1).

### How to setup in local machine

Make sure you're in the root folder of the project.

```bash
$ cd togo
```

Use the package manager [npm](https://www.npmjs.com/) to install dependencies. Run this command in your terminal.

```bash
$ npm install
```

Configure MySQL connection host, user and password in this file:

```
togo/config/ormconfig.ts
```

Run migration scripts.

```bash
npm run typeorm migration:run
```

Run the server.

```bash
$ npm run start
```

### Usage

Once the server is running, you may now send POST requests to `/api/todos`.
For this iteration please use userId=10172512 when sending the requests. This is a sample cURL command for the API:

```bash
curl -X POST localhost:5000/api/todos \
  -H "content-type: application/json" \
  -d '{"task":"Some task","userId": 10172512}'
```

### How run the tests

Run the following command in your terminal.

```bash
$ npm run test
```

### What I love about my solution

During the building of the solution, it made me appreciate the creation of tests even more. Thanks for this challenge!
