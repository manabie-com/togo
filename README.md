### How to run
* docker-compose up
* make dev
* Import Postman exported collection from `/api/postman/todo.postman_collection.json` for testing

### What I did
* Write unit test for `services` layer -> /usecase/interactor/*.go
* Make this code DRY
* Change from using `SQLite` to `Postgres` with `docker-compose`
* Optimization
    * Apply `Clean Architecture`
    * Timing out long-running requests with `3 seconds` threshold

### Further improvements
* Write more tests
* Add pagination for fetching tasks
* i18n supports for user-friendly error message
* Custom error codes on the server side
* Implements a caching layer using `Redis` above primary repositories for performance enhancements
* Full-text searching supports using `Elasticsearch`
* Move common packages inside `/infrastructure` to a new git repo for reusing later
* A gRPC server for internal inter-services communication in a microservices-based system (maybe a gRPC-gateway for HTTP requests coming from client side)
* Documentation following OpenAPI
* Better Makefile and OS-targeted builds
* For enterprise purpose, this `To Do` app might have
    * Usage plans, for example
        * `Free`: max 10.000 tasks
        * `Individual`: max 1.000.000 tasks
        * `Business`: unlimited tasks can be created
    * Task namespaces
        * Period-basis such as `Daily`, `Weekly`, `Monthly`, ...
        * User'custom namespaces
    * Reminder
    * Task co-creation
