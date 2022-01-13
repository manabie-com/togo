### Start-End date: 12/01/2022 01:45:12 - 14/01/2022 01:21:20
## How to run your code locally?
```
yarn start
```
## A sample “curl” command to call your API
- **API:**

  POST: `localhost:3000/user/:userId/task`

  body:
  ```
  {
    taskId: number
  }
  ```

- **Sample “curl”**
  ```
  curl -d "{\"taskId\": 6}" -H "Content-Type: application/json" -X POST localhost:3000/user/2/task
  ```

## How to run your unit tests locally?
```
yarn test
```

## What do you love about your solution?
I made it as simple as possible for this assignment
- Lightweight DB: **lowdb** (no need to setup DB server)

- Typescript framework: **nestjs**

- nestjs default testing framework: **jest**

  file location: `test\app.controller.spec.ts`

## What else do you want us to know about however you do not have enough time to complete?
NO
