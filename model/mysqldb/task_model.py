from datetime import datetime
import mysql.connector as connector
from config import config


class TaskModel:
    def __init__(self):
        self.db = None
        self.db_cursor = None
        try:
            self.db = connector.connect(host=config.HOST, user=config.USER_NAME, passwd=config.PASSWORD, db=config.DATABASE, charset='utf8')
            self.db_cursor = self.db.cursor()
        except Exception as e:
            print(e)
            print("Can't connect to db...")

    def get_task_by_user_id(self, user_id, from_time):
        query_cd = "SELECT * FROM test.task WHERE user_id = {} and created_date > {}".format(user_id, from_time)
        self.db_cursor.execute(query_cd)
        inf = self.db_cursor.fetchall()
        result = []
        for i in inf:
            result.append(i)
        self.db_cursor.close()
        return result

    def insert(self, info_data):
        query_cd = "INSERT INTO test.task (content, user_id, created_date, name) VALUES ('{}', {}, {}, '{}')".format(info_data.get(
            "content"), info_data.get("user_id"), info_data.get("created_date"),  info_data.get("name"))
        try:
            self.db_cursor.execute(query_cd)
            self.db.commit()
            self.db_cursor.close()
            return True
        except Exception as e:
            print(e)
            return str(e)

    def update(self, info_data):
        query_cd = "UPDATE test.task SET content = {} WHERE name = '{}'".format(info_data.get("content"),
                                                                                info_data.get("name"))
        try:
            self.db_cursor.execute(query_cd)
            self.db.commit()
            self.db_cursor.close()
            return True
        except Exception as e:
            return str(e)


# a = TaskModel()
# data = {
#     "username": "nguyen duc quan",
#     "password": "quan",
#     "max_todo": 100
# }
# r = a.insert(info_data=data)
# print(r)
