# go-blog-service
Blog service with Golang programming language

# Clean Architecture + SOLID principle + Dependency Injection



# Usage
In Mac: set env by run command

`export BLOG_TEST_DB_URL="username:password@/dbName?parseTime=true"`
`export BLOG_DB_URL="postgres://postgres:abcd1234@golangvietnam.com:54321/todos_tasks"`
`export BLOG_DB_URL="port=54321 host=golangvietnam.com user=postgres password=abcd1234 dbname=todos_tasks sslmode=disable"`

example:
    `export BLOG_TEST_DB_URL="root:password@tcp(localhost:23306)/blog?parseTime=true"`

run command: `go test ./...`

