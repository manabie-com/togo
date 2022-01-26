### Prerequisite
Golang, Redis must be installed

### Source code structure
|-api_service/
|-account_service/
|-session_service/
|-todo_service/
|-docs/
    |-api_references.md
|-db/
    |-schema.sql

Explaination:
- **api_service**: Root service, contains all API methods.
- **account_service**: Sub-service, handling account-related processes.
- **session_service**: Sub-service, handling session-related processes.
- **todo_service**: Sub-service, handling todo-related processes.
- **docs**: Contains application documents
  - api_references.md: API documentation
- **db**: Contain database structure 
  - schema.sql: database schema

### How to run
1. Run sql schema (MySql)
2. Config enviroment ```.env``` file based on ```.env.example```
3. Go to all services folder, use command ```go run main.go```
4. Follow api_references.md to call APIs
