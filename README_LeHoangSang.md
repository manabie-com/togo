### How to run
* Install [docker][1] on MacOS or [docker][2] on Windows
* If you are using Windows, you need to install [docker-compose][3]
* Close this repo and `cd` to this repo folder
* Run the integration test by ` docker-compose up integration_tests && docker stop postgres_for_test && docker rm postgres_for_test`
* Run the unit test for services layer by `docker-compose up unit_tests_services_layer && docker stop postgres_for_test && docker rm postgres_for_test`
* Run the unit test for storage layer by `docker-compose up unit_tests_storage_layer && docker stop postgres_for_test && docker rm postgres_for_test`
* Run the main app: `docker-compose up app`

### Project structure
```
├── README.md
├── README_LeHoangSang.md
├── scripts // scripts for init db
├── internal
│   ├── config //includes function get database config from environment variables
│   ├── integration_tests // define scenario for integration tests
│   ├── model // struct for login success response, error response, get tasks response,...
│   ├── services // includes handler and unit test for services layer
│   ├── storages
│   │   ├── dbinterface // this is interface includes 3 function: RetrieveTasks, AddTask, ValidateUser
│   │   ├── entities.go // includes struct for user and task
│   │   ├── pg // this implement dbinterface and override 3 function in dbinterface to action with PostgresDB 
│   │   ├── sqlite // // this implement dbinterface and override 3 function in dbinterface to action with Sqlite
│   │   └── test // unit test for storages layer
├── main.go
├── scripts
```

### Something new
* Because in storages layer, type `PostgresDB` and type `LiteDB` are implemented from `DBInterface`, so that struct `ToDoService` is changed from:
```go
    type ToDoService struct {
        JWTKey string
        Store  *sqllite.LiteDB
    }
```
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;to
```go  
    type ToDoService struct {
        JWTKey string
        Store  dbinterface.DBInterface
    }
```
With these structures (base on strategy design pattern), we easily create ToDoService with behaviour for Postgres or Sqlite by NewToDoServices function
```go
func NewToDoServices(jwt string, driverName string,dbInfo string) (*ToDoService,error) {
	db,err:=dbinterface.NewDB(driverName,dbInfo)
	if err != nil {
		return nil,err
	}
	if driverName == config.DBType.Sqlite {
		return &ToDoService{
			JWTKey: jwt,
			Store:  &sqllite.LiteDB{
				DB: db,
			},
		},nil
	}
	return &ToDoService{
		JWTKey: jwt,
		Store:  &pg.PostgresDB{
			DB: db,
		},
	},nil
}
```
* Use global config variable for request path (`/login`, `/tasks`), JWT (get from env), db type (`postgres`,`sqlite`),... Detail for this info is coded in package config
* Change the way for creating response body from using map[string]string to use struct which is defined in package model
* Validate the limit of tasks in day in add task function
* Add integration tests for a workflow in this project
* Add unit test for `services` layer
* Add unit test for `storages` layer
* Change from using SQLite to Postgres with docker-compose
* Change data type of return values in function `RetrieveTasks` and function `AddTask` in package storages from `error` to `config.ErrorInfo`
to help us create `HTTP Responcode` more exactly because in old code, `services layer` call `storage layer` and if it receives error, it only returns status code 500 to client

 
  
[1]: https://docs.docker.com/docker-for-mac/install/
[2]: https://docs.docker.com/docker-for-windows/install/
[3]: https://dockerlabs.collabnix.com/intermediate/workshop/DockerCompose/How_to_Install_Docker_Compose.html




