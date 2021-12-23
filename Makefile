test:
	docker exec tasks_api  go test ./app/... ./repository/... ./services/... ./controller/...

mock:
	docker exec tasks_api mockgen -destination ./mocks/controller/controller.go github.com/manabie-com/backend/controller/tasks I_TaskController
	
	docker exec tasks_api mockgen -destination ./mocks/repository/repository.go github.com/manabie-com/backend/repository I_Repository
	docker exec tasks_api mockgen -destination ./mocks/taskservice/taskservice.go github.com/manabie-com/backend/services/task I_TaskService
	docker exec tasks_api mockgen -destination ./mocks/taskservicevalidate/taskservicevalidate.go github.com/manabie-com/backend/services/task I_TaskServiceValidate
	docker exec tasks_api mockgen -destination ./mocks/userservicevalidate/taskservicevalidate.go github.com/manabie-com/backend/services/user I_UserServiceValidate