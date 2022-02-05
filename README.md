# ToDo Application


## Setup instructions

Create a new .env file based on the given `.env.example`. 

Afterwards, initialize the database by doing:

```
make serve
make migrate_up

```

Run the curl command to test the application:

```
curl -XPOST -H "Content-type: application/json" -d '{
	"title": "Test Task",
	"content": "This is the exam",
	"is_complete": false,
	"username": "roandayne"
}' 'http://localhost:8080/api/tasks'
```

## Testing
To test the application run `make run_test`

## Delete Database
To delete database run `make migrate_down`