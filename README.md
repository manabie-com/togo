### Overview
This is a simple backend for a good old todo service, right now this service can handle login/list/create simple tasks.  

The repo was forked to implement additional requirements and improvement from:
https://github.com/manabie-com/togo

### How to start
- Start the service: `make start`
- Unit tests: `make test-unit`
- Integrate tests:
  - Start Postgres: `make postgres-start`
  - Init Postgres DB: `make postgres-init`
  - Start tests: `make test-integrate`
- Import Postman collection from `docs` to check example.

### Todo / Comment list
- Split `services` layer to `use case` and `transport` layer
  - Can use `grpc-gateway` to expose this API layer. Pros: similar syntax if internal services also use `gRPC`. Cons: not enough time, not needed.
  - Use `github.com/gorilla/mux` instead
- Improve unit test for `services`
- Move config variable out of source code. Setup config profile on environment
- Add debug tests script for APIs top-down services logic. Currently using Postman collection for manual test.
- Add server graceful shutdown
- Login request need encrypt
- Need to fix SOLID
- Add database connection in sequence diagram
- Refactor `storages` to reduce duplicated query logic. Low priority since expected to use only one DB type.
- Reformat logging base on log framework. Add metrics.
