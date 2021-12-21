## This is simple application to add tasks daily for user
## **1. Architecture:**
- Microservice
- Language: python
- Framework: flask
- Database: sqlite3
## **2. Design:** 
### 2.1 Tables in database: users, tasks and tasks_daily
### Modules: 3 module 
- users: to store users info 
- tasks: to store tasks info
- tasks_daily: to store tasks of user daily
### 2.2 Process when calling api:
- **Step 1**: calling api to add task 
- **Step 2**: process api \
&nbsp; - Check request's validation: number tasks inserted today < N(maximum task per day). \
&nbsp; - Insert data into 2 tables: tasks and tasks_daily.
- **Step 3**: return success or failed

### 2.3 Test: 2 types test - unit test and integration test
- sqlite3 supports to stored database in memory while running application, so I would like to create database and tables in memory to test instead mocking database and tables as per usual
- I suppose module users and tasks are tested, so I just write tests for tasks daily module 
## **3. Before run application, please setup env:**
- install python3 and python3-pip
- install libs: 
	> `pip install -r requirements.txt`
- create env: 
	> `python3 -m venv my_project_env`
- active env:
	> `source my_project_env/bin/activate`
- prepare data before running app:
	> `python3 prepare_data.py`
## **4. Run application:**
- Run app by command: 
	> `python3 daily_task_app.py`

- "curl" command to call API to add task: 
	> `curl -X POST http://127.0.0.1:5000/daily-task -H 'Content-Type: application/json' -d '{"project_id":"project1", "task":"task_hihi","owner":"hoaipham", "deadline_at": "2021-12-19 15:36:22"}'`
## **5. Run unit test**
- To run test, please replace this line in `my_config.py`
	> `DB_CONNECTION_STR = ':memory:'`
- Run test: 
	> `python3 unittest_task_daily.py`

## **6. Run integration test**
- Not yet finished

If there are any questions, feel free to ask me via hoaipham011997@gmail.com. 
