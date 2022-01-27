### Prerequisite
Golang, Docker must be installed.

### Source code structure
```
|-api_service/
|-account_service/
|-session_service/
|-todo_service/
|-docs/
    |-api_references.md
|-db/
    |-schema.sql
|-docker-compose.yaml
```

Explaination:
- **api_service/**: Root service, contains all API methods.
- **account_service/**: Sub-service, handling account-related processes.
- **session_service/**: Sub-service, handling session-related processes.
- **todo_service/**: Sub-service, handling todo-related processes.
- **docs/**: Contains application documents.
  - api_references.md: API documentation.
- **db/**: Contain database structure .
  - schema.sql: database schema.

### How to run
1. Run ```docker-compose -f docker-compose.yaml up``` to start MySql and Redis containers.
2. Go to each services folder, use command ```cp .env.example .env``` to generate ```.env``` file for each services.
  - Same service must have the same port. Ex: api_service/.env ```ACCOUNT_SERVICE_PORT``` and account_service/.env ```ACCOUNT_SERVICE_PORT``` must be identical.
  - session_service/.env ```ACCESS_SECRET``` is arbitrary.
3. Go to each services folder, use command ```go run main.go``` or pre-build with ```go build main.go``` and ```./main```.
4. Follow api_references.md to call APIs.
