export TODO_DB_URL="postgres://postgres:abcd1234@localhost:54321/todo_tasks?sslmode=disable"
echo $PWD
cd $PWD/cmd/

go run todo.go
