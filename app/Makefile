run:
	go run main.go

mock_gen:
	mockgen -package mock_task_db -destination repo/mongo/mock_task_db/task.go github.com/manabie-com/togo/app/repo/mongo/task Repository
	mockgen -package mock_user_db -destination repo/mongo/mock_user_db/user.go github.com/manabie-com/togo/app/repo/mongo/user Repository 
		
test:
	go test -v -cover ./...
