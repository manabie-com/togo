# ToDo Application


## Setup instructions

Create a new .env file based on the given `.env.example`. Load the created env to your shell:

```
source .env
```

Afterwards, initialize the database by doing:

```
make db_start
make createdb
make migrateup
make run
```

Run the curl command to test the application:

```
curl -XPOST -H "Content-type: application/json" -d '{
	"title": "Test Task",
	"content": "This is the exam",
	"is_complete": false,
	"fullname": "Roan Dino"
}' 'http://localhost:8080/api/tasks'
```

## Testing
To test the application run `make test_app`

## Delete Database
To delete database run `make migratedown`