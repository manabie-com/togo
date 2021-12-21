from sqlite3 import connect
import unittest
import sqlite3
from task_daily_dao_module import TaskDaily
from datetime import datetime, timedelta
from sql_util import SqliteUtil
import json


class TaskDailyDaoUnittest(unittest.TestCase):

    def setup(self, TASK_DAILY_TABLE, task_daily):
        drop_table = 'DROP TABLE IF EXISTS {}'.format(TASK_DAILY_TABLE)
        task_daily.execute_query(drop_table)
        create_new_table = "CREATE TABLE {}(id text PRIMARY KEY, project_id text, task_id text, task text, owner text, created_at text, updated_at text, deadline_at text)" .format(TASK_DAILY_TABLE)
        task_daily.execute_query(create_new_table)

    def compare_object(self, expect_result, actual_result):
        self.assertEqual(expect_result.get('id'), actual_result.get('id'))
        self.assertEqual(expect_result.get('task_id'), actual_result.get('task_id'))
        self.assertEqual(expect_result.get('task'), actual_result.get('task'))
        self.assertEqual(expect_result.get('owner'), actual_result.get('owner'))
        self.assertEqual(expect_result.get('project_id'), actual_result.get('project_id'))
        self.assertEqual(expect_result.get('created_at'), actual_result.get('created_at'))
        self.assertEqual(expect_result.get('updated_at'), actual_result.get('updated_at'))
        self.assertEqual(expect_result.get('deadline_at'), actual_result.get('deadline_at'))

    def test_insert_task_daily_success(self):
        # setup database and create table before testing insert task
        TASK_DAILY_TABLE = 'tasks_daily'
        task_daily = TaskDaily(TASK_DAILY_TABLE)
        self.setup(TASK_DAILY_TABLE, task_daily)

        actual_list = task_daily.get_all_tasks_daily()
        expect_size_list = 0
        self.assertEqual(expect_size_list, len(actual_list))

        now = datetime.now().strftime('%Y-%m-%d %H:%M:%S')
        deadline = (datetime.now() + timedelta(days=1)).strftime('%Y-%m-%d %H:%M:%S')
        task_daily_object = {
            'id': 'id3',
            'task_id': 'task_id1',
            'task': 'task1',
            'owner': 'hoaipham',
            'project_id': 'project1',
            'created_at': now,
            'updated_at': now,
            'deadline_at': deadline
        }
        task_daily.insert_task_daily(task_daily_object)
        actual_list = task_daily.get_all_tasks_daily()
        expect_size_list = 1
        self.assertEqual(expect_size_list, len(actual_list))
        self.compare_object(task_daily_object, actual_list[0])
    
    
    def test_insert_task_daily_failed(self):
        TASK_DAILY_TABLE = 'tasks_daily'
        task_daily = TaskDaily(TASK_DAILY_TABLE)
        self.setup(TASK_DAILY_TABLE, task_daily)

        now = datetime.now().strftime('%Y-%m-%d %H:%M:%S')
        deadline = (datetime.now() + timedelta(days=1)).strftime('%Y-%m-%d %H:%M:%S')
        task_daily_object1 = {
            'id': 'id1',
            'task_id': 'task_id1',
            'task': 'task1',
            'owner': 'hoaipham',
            'project_id': 'project1',
            'created_at': now,
            'updated_at': now,
            'deadline_at': deadline
        }
        task_daily.insert_task_daily(task_daily_object1)
        task_daily_object2 = {
            'id': 'id2',
            'task_id': 'task_id2',
            'owner': 'hoaipham',
            'project_id': 'project1',
            'created_at': now,
            'updated_at': now,
            'deadline_at': deadline
        }
        self.assertRaises(Exception,task_daily.insert_task_daily(task_daily_object2))

    





if __name__ == '__main__':
    unittest.main()


        
