run-dev:
	go run main.go

run-test:
	go test ./services/... ./controllers/...

create-mocks:
	mockgen -source=repositories/task-repository.go -destination=mocks/repositories/task-repository_mock.go -package=repositories_mocks
	mockgen -source=repositories/user-repository.go -destination=mocks/repositories/user-repository_mock.go -package=repositories_mocks
	mockgen -source=services/task-service.go -destination=mocks/services/task-service_mock.go -package=services_mocks
	mockgen -source=controllers/task-controller.go -destination=mocks/controllers/task-controller_mock.go -package=controllers_mocks