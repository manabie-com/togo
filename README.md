### Application Specification
- How to limit N tasks per day => I created a new table named "configurations" for input the maximum capacity per day. If the day is not exists, it means 0 and the user can't create new task.
- When the user submits a new task, I will check total task which was created in this day and compare with the capacity in the configuration table.


### Completed requirements:
There're something that I have done in this test
- I split application to service, transport, repository layer
- I tried using a new DI Framework, google wire.
- Create new API for configuration capacity per day.
- I remove token filter in Serve HTTP in old code, and apply the JwT filter in Gin Framework.


### Non-completed requirements:
Because of my limitation of time, so many of requirements below I couldn't complete. 
- Fork this repo and show us your development progress by a PR. I have just fork your repository and submit pull request then.
- I haven't yet completed testing function. I have not found the best solution for testing with google wire.
- Change from using SQLite to Postgres with docker-compose
- Password is being stored as raw text in DB.

### Running my application

	go run main.go
	
- Import postman file: docs/togo_new_api.json, I added 2 APIs for create configuration and get configuration.
- Call API create configuration firstly to limit task per day