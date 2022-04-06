### How to run the code locally
- Prerequisite:
    - Docker Desktop
- First, open terminal in the project directory
- Then, we need to deploy the `docker-compose` by using
```bash
docker-compose up --build
```
- Third, execute the following command to auto-generate the DB schema
```bash
docker exec assignment_app_1 python manage.py recreate_db
```
- Now the app is live =))
### CURL command
```bash
curl --location --request POST 'http://localhost:5000/api/todo' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "titlemain",
    "description": "b",
    "limit": 1,
    "todos": [{
        "title": "title1",
        "description":"desc1"
    },
    {
    
        "title": "title2",
        "description":"desc2",
        "todos": [
            {
                "title": "title3",
                "description": "desc3"
            }
        ]
    }]
}'
```
### Notes: Please import the command to Postman because it's quite hard to execute it right away
### How to run your unit tests locally?
- Unit tests: 
```bash
    docker exec assignment_app_1 pytest -v tests/unit
```
- Integration tests:
```bash
    docker exec assignment_app_1 pytest -v tests/integration
```
- All tests
```bash
    docker exec assignment_app_1 pytest -v tests/
```
### What do you love about your solution?
- First, I apply `Domain driven design` with 3 layers:
    - Domain model: Refected by ORM models, contain core business
    - service layer: implement business logic, deal with domain model and repository
        -  services: implement logic for domain model
        - unit of work: execute commit/rollback steps to DB based on services return value. After finished, unit of work will be closed
    - adapters: handle data consistency
        - repository: keep CRUD execution on database
- Second, for database design I chose [Composite design pattern](https://refactoring.guru/design-patterns/composite) to handle todo list relationship
    - A todo list can contain multiple todo items, using `todo_list_id` foreign key
    - A todo list can have one parent todo list, based on `parent_id`
    - A todo list can contain other smaller todo lists
    - Todo list and todo items models inherit from one abstract model
- Third, I used Docker container so no need to setup anything in local
- Fourth, I tried to apply Test driven development in writing the tests (not fully completed) by writing the simple failed testcase for the function before implementing, then implement the function until the testcase is succesfully and refactor the test code later.