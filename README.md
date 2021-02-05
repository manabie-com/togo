`https://github.com/manabie-com/togo/tree/v0.0.1`

# Requirements
Right now a user can add many task as they want, we want ability to limit N task per day.

Example: users are limited to create only 5 task only per day, if imit reached, return 4xx code to client and ignore the create request.
#### Backend requirements
- Write integration tests for this project
- Make this code DRY
- Write unit test for `services` layer
- Change from using `SQLite` to `Postgres` with `docker-compose`
#### Frontend requirements
- Implement in React with hooks
- A login interface
- A list/create tasks interface
- Nice state management mechanism
#### Optional requirements
- Write unit test for `storages` layer

# Instruction

```
	├── ...
	├── docs
	|   ├── ...
	│   └── togo_pg.postman_collection.json -- Doc API Change from using `SQLite` to `Postgres` with `docker-compose`
	├── fe -- The test of FE
	├── internal
	│   ├── services
	│   │   ├── task.go
	│   │   └── task_test.go -- Write unit test for `services` layer
	│   └── storages
	│      ├── postgres -- Change from using `SQLite` to `Postgres` with `docker-compose`
	│      │   ├── db.go
	│      │   └── db_test.go -- Write unit test for `storages` layer
	│      └── sqlite
	│         ├── db.go
	│         └── db_test.go -- Write unit test for `storages` layer
	├── main_test.go Write integration tests for this project
	└── main.go
```

- folder scripts: Sample data postges (Change from using `SQLite` to `Postgres` with `docker-compose`)


## Backend requirements reply

### Requirements:
* docker >= 17.12.0+
* docker-compose

### Setup portgres (Change from using `SQLite` to `Postgres` with `docker-compose`)
- This project run command `docker-compose up -d`
  
#### Access to postgres:
* `localhost:5432`
- **Username:** postgres (as a default)
- **Password:** changeme (as a default)

#### Access to PgAdmin: 
- **URL:** `http://localhost:5050`
- **Username:** pgadmin4@pgadmin.org (as a default)
- **Password:** admin (as a default)

#### Add a new server in PgAdmin:
- **Host name/address** `postgres`
- **Port** `5432`
- **Username** as `POSTGRES_USER`, by default: `postgres`
- **Password** as `POSTGRES_PASSWORD`, by default `changeme`

### Right now a user can add many task as they want, we want ability to limit N task per day.

*I am using golang's global variable. When the project is started, it will be init with the default value which has already set. But the variable will be reset in value when we stop or start project*

**Alternativeway**

- Database usage for storage (Not recommended)

- Use the `Redis Incr` (High Priority)

```go

var limitNtaskPerday = make(map[string]int, 0)
var maxLimitNtaskPerday int = 5

func (s *ToDoService) limitNtaskToday(resp http.ResponseWriter, req *http.Request) bool {
	id, ok := userIDFromCtx(req.Context())
	if !ok {
		resp.WriteHeader(http.StatusBadRequest)
		formatOutput(resp, map[string]string{
			"error": "unknow",
		})
		return false
	}

	key := getCurrentDate() + "_" + id
	if limitNtaskPerday[key] >= maxLimitNtaskPerday {
		resp.WriteHeader(http.StatusLocked)
		formatOutput(resp, map[string]string{
			"error": "created tasks over limit per day",
		})
		return false
	}
	return true
}

func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
	// code ...

	key := getCurrentDate() + "_" + userID
	limitNtaskPerday[key]++
	formatOutput(resp, map[string]*storages.Task{
		"data": t,
	})
}


```
  
### Backend requirements run test
- Go inside of directory,  `cd be`
- Run integration tests for this project `go test -v main_test.go`
- Run unit test for `services` + `storages` layer `go test -v ./...`

### Backend requirements make this code DRY

- #1

**Original code**

```go

func (s *ToDoService) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	id := value(req, "user_id")
	if !s.Store.ValidateUser(req.Context(), id, value(req, "password")) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := s.createToken(id.String)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]string{
		"data": token,
	})
}

```


```go

func formatOutput(resp http.ResponseWriter, v interface{}) {
	json.NewEncoder(resp).Encode(v)
}

```

**To become**

```go

func (s *ToDoService) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	id := value(req, "user_id")
	if !s.Store.ValidateUser(req.Context(), id, value(req, "password")) {
		resp.WriteHeader(http.StatusUnauthorized)
		formatOutput(resp, map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := s.createToken(id.String)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		formatOutput(resp, map[string]string{
			"error": err.Error(),
		})
		return
	}

	formatOutput(resp, map[string]string{
		"data": token,
	})
}

```

- #2

**Original code**

```go

func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
	
	// code ...
	now := time.Now()
	userID, _ := userIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = now.Format("2006-01-02")

	// code ...
}

```

```go

func getCurrentDate() string {
	now := time.Now()
	return now.Format("2006-01-02")
}

```

**To become**

```go

func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
	
	// code ...
	userID, _ := userIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = getCurrentDate()

	// code ...
}

```

## Frontend requirements reply
* Start API `go run main.go`
* Go inside of directory,  `cd fe`
* Run this command `npm install`
* Start project `npm start`