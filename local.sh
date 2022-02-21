export TODO_DB_URL="postgres://postgres:abcd1234@localhost:54321/todo1?sslmode=disable"
cd $PWD/cmd/

go run todo.go
