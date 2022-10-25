wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

CompileDaemon --build="go build -o main main.go" --command=./main