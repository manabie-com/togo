from sql_util import SqliteUtil
from my_config import TABLE_USERS, TABLE_TASKS, TABLE_DAILY
from user_dao_module import User
import json


sql_create_users_table = "CREATE TABLE {}(id text PRIMARY KEY, username text, created_at text, updated_at text)".format(TABLE_USERS)
sql_create_tasks_table = "CREATE TABLE {}(id text PRIMARY KEY, task_name text, project_id text, created_at text, updated_at text)".format(TABLE_TASKS)
sql_create_tasks_daily_table = "CREATE TABLE {}(id text PRIMARY KEY, project_id text, task_id text, task text, owner text, created_at text, updated_at text, deadline_at text)".format(TABLE_DAILY)

sql_util = SqliteUtil()

sql_util.drop_table(TABLE_USERS)
sql_util.drop_table(TABLE_TASKS)
sql_util.drop_table(TABLE_DAILY)

sql_util.execute_sql(sql_create_users_table)
sql_util.execute_sql(sql_create_tasks_table)
sql_util.execute_sql(sql_create_tasks_daily_table)

# prepare user
user = {
  "id": "id1",
  "username": "hoaipham",
  "created_at": "2021-12-19 22:10:15",
  "updated_at": "2021-12-19 22:10:15"
}

json_object = json.loads(json.dumps(user))
user = User(TABLE_USERS)
user.insert_user(json_object)
