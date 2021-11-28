# Togo
### Starting the appplication
You can run the test with
```bash
go test "./test"
```

Then, to run the code locally
```bash
go run "./cmd"
```

### APIs
There are only 1 API to create todo tasks.
```bash
curl --location --request POST 'http://localhost:9000/api/v1/todo' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "make breakfast",
    "user_id": 1
}'
```

### Project layout
The project follows [the golang standard layout](https://github.com/golang-standards/project-layout) 
with some modification to keep the layout as simple as possible. By separating logic to its own module, 
I hope to increase the maintainability of the code base.

The "cmd" folder contains the "main.go" file which is the stating point of the application.
The folder contains code for dependencies injection and linking endpoints with its respective HTTP handler.

The "internal" folder are for storing code and modules that can not be imported from outside. 
It's sub folders are:
 - The "service" folder contains the top level HTTP handler.
 - The "store" folder contains abstraction for keeping the application states.
 - The "model" folder contains domain entities,which in this case are todo tasks.
 - Others folder are utility folders: logging, tracing, convert, etc.

### Architecture
Because the requirements are a simple CURD service, or in this case a C ( create ) service. 
I choose a simple 2 layers architecture:

- Handler layer (a.k.a controllers) is responsible for parsing between external requests and internal domain entities.
it does it best to become a wall separate the chaos outside world and the peaceful internal world.

- Storage layer (a.k.a repository) has the same responsibility with the handler layer,
which is to connect the internal with the outside world. Normally, this layer is for keeping the service state, 
 in other words, connect to a database.

The service currently only connect to an in memory database, because connecting to a real database will
make running a simple project, which is only used for interviewing, become complicated. However, I have ensured
that connecting to a real database is simple, which is just adding another store implements the store interface.

### Requirements

- Implement one single API which accepts a todo task and records it
  - There is a maximum **limit of N tasks per user** that can be added **per day**.
- Write integration (functional) tests
- Write unit tests
- Choose a suitable architecture to make your code simple, organizable, and maintainable
- Write a concise README
  - How to run your code locally?
  - A sample “curl” command to call your API
  - How to run your unit tests locally?
  - What do you love about your solution?
  - What else do you want us to know about however you do not have enough time to complete?
