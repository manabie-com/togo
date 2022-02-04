Make env file using the .env.example

Type `source .env` to your terminal

Type the following commands:
1. make db_start
2. make createdb
3. make migrateup
4. make run
5. curl -XPOST -H "Content-type: application/json" -d '{
	"title": "Test Task",
	"content": "This is the exam",
	"is_complete": false,
	"fullname": "Roan Dino"
}' 'http://localhost:8080/api/tasks'

To test the application run `make test_app`

To delete database run `make migratedown`