genmock:
	moq -pkg user -out internal/usecase/user/UserServiceMock.go internal/usecase/user Service
	moq -pkg user -out internal/usecase/user/UserCacheMock.go internal/usecase/user Cache
	moq -pkg auth -out internal/usecase/auth/AuthServiceMock.go internal/usecase/auth Service
	moq -pkg auth -out internal/usecase/auth/UserRepositoryMock.go internal/usecase/auth Repository
	moq -pkg task -out internal/usecase/task/TaskServiceMock.go internal/usecase/task Service
	moq -pkg task -out internal/usecase/task/TaskRepositoryMock.go internal/usecase/task Repository

run:
	docker-compose up --build 
