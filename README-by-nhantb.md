### To execute the unit test for the services layer
```cgo
$ go test ./internal/services
```

### To execute the unit test for the storages layer
```cgo
$ docker rm -f $(docker ps -a -q); docker-compose up db storages
```

### To execute the integration test for this server
```cgo
$ docker rm -f $(docker ps -a -q); docker-compose up db integration
```

### To start the server and check the API with postman
```cgo
$ docker rm -f $(docker ps -a -q); docker-compose up db app
```

### Something new
- Implement validation for the limitation on creating new task
- Add the integration test for the workflow following sequence diagram
- Add the unit test for 'services' layer
- Add the unit test for 'storages' layer
- Add 'config' package to centralize the way to get environment
configuration (mockup variables, via docker-compose, etc)
using for deployment/integration test, etc
- Add 'utils' package for using as common func further.\
For example:
  - Make outbound HTTP request
  - Jsonify for objects
- Add 'model' package to centralize the input/output models
and something liked that
- Change to using Postgres with Docker-compose
- ToDoService.Store has been changed to be interface
to be able to execute unit test with a mockup database
  
### Need to be optimized
- I prefer another project structure or would if the project was
written like this:
  </br>
  </br>
  ```
  github.com/manabie-com/togo
  |__ action (for the use case)
  |__ api (for the transport, middleware, etc )
  |__ client (for the outbound API client)
  |__ conf (to centralize the configuration, as mentioned above)
  |__ model (the in/out model and something like that will be stored here)
       |__ enum
       |__ error_code
       |__ ... 
  
  ```