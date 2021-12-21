from daily_task_app import app
from task_daily_dao_module import TaskDaily
from task_dao_module import Task
from flask import json

TASKS_TABLE = 'task'
TASK_DAILY_TABLE = 'tasks_daily'
task = Task(TASKS_TABLE)
task_daily = TaskDaily(TASK_DAILY_TABLE)

def setup():
    drop_table_tasks = 'DROP TABLE IF EXISTS {}'.format(TASKS_TABLE)
    drop_table_task_daily = 'DROP TABLE IF EXISTS {}'.format(TASK_DAILY_TABLE)

    task.execute_query(drop_table_tasks)
    task_daily.execute_query(drop_table_task_daily)

    create_new_task_table = "CREATE TABLE {}(id text PRIMARY KEY, task_name text, project_id text, created_at text, updated_at text)".format(TASKS_TABLE)
    create_new_task_daily_table = "CREATE TABLE {}(id text PRIMARY KEY, project_id text, task_id text, task text, owner text, created_at text, updated_at text, deadline_at text)" .format(TASK_DAILY_TABLE)
    
    task.execute_query(create_new_task_table)
    task_daily.execute_query(create_new_task_daily_table)

def test_insert_and_get_data():     
    setup()

test_insert_and_get_data()