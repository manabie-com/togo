# TODO APP
Implement simple application with limit adding task of the user

### Requirement
```bash
go version go1.16.6 darwin/amd64
```

### Local Build
The result of building creates at example-build folder
```bash
go build -o build/todo
```

## Adapter
Command line
```bash
Usage: ./todo start [PORT]
Example: ./todo start 3000
```

## TestAPI
```bash
curl -X POST http://localhost:3000/user/task/new-task
   -H "Content-Type: application/json"
   -d '{"task_name": "working", "user_id":1}'
```