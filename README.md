### Notes
To make it run:
- `go run main.go`
- Import Postman collection from docs to check example

---

### Functional requirement checklist:
- [x] Right now a user can add as many tasks as they want, we want the ability to **limit N tasks per day**. For example, users are limited to create only 5 tasks only per day, if the daily limit is reached, return 4xx code to the client and ignore the create request.
	- More under Developer Notes

### Non-functional requirements checklist:
- [x] **A nice README on how to run, what is missing, what else you want to improve but don't have enough time**
- [x] **Consistency is a MUST**
- [x] Fork this repo and show us your development progress by a PR
- [ ] Write integration tests for this project
- [x] Make this code DRY
	- More under Developer Notes
- [ ] Write unit test for the services layer
- [ ] Change from using SQLite to Postgres with docker-compose
- [x] This project includes many issues from code to DB structure, feel free to optimize them
	- More under Developer Notes
- [ ] Write unit test for storages layer
- [ ] Split services layer to use case and transport layer
	- More under Developer Notes

---

#### Developers Notes
- I initially intended to use Python for this exercise but instead decided to take the opportunity to get a feel of coding in Go
- Limit user to N tasks per day
	- Additional functions:
		```
		tasks.go
			addTaskAllowed
		db.go
			RetrieveUserTaskLimit
			CountTasks
		```
- Refactoring for DRY/optimizations/consistency
	- [x] Function comments at the beginning of the function
	- [x] Helper function for sending responses
	- [ ] Move authentication related functions into separate module
	- [ ] Variable names
		I followed the pattern for now as there might be company-specific convention that I might not be aware of.
	- [ ] Additional endpoints under /tasks instead of checking what kind of method was called to allow expansion of functionalities
- Split services layer to use case and transport layer
	- Use case layer functions:
		addTask
		addTaskAllowed
		createToken
		getAuthToken
		listTasks
		userIDFromCtx
		validToken
		value
	- Transport layer functions:
		sendCodeResponse
		sendOKResponse
	    ServeHTTP
- Possible enhancements:
	- Delete task(s)
	- List tasks on a given date range
	- Categorizing of tasks as "Done" or "Pending"