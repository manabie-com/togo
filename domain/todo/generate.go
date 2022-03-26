//go:generate mockgen  -destination=../../test/todomock/todo_mocks.go -package=todomock . UserRepo,TaskRepo
package todo
