from datetime import datetime
import mysql.connector as connector
from config import config


class UserModel:
    def __init__(self):
        self.db = None
        self.db_cursor = None
        try:
            self.db = connector.connect(host=config.HOST, user=config.USER_NAME, passwd=config.PASSWORD, db=config.DATABASE, charset='utf8')
            self.db_cursor = self.db.cursor()
        except Exception as e:
            print(e)
            print("Can't connect to db...")

    def get_user_by_name(self, username):
        query_cd = "SELECT * FROM test.user WHERE username = '{}'".format(username)
        self.db_cursor.execute(query_cd)
        inf = self.db_cursor.fetchall()
        result = {}
        if inf:
            result = {
                "id": inf[0][0],
                "password": inf[0][1],
                "max_todo": inf[0][2],
                "username": inf[0][3],
            }
        self.db_cursor.close()
        return result

    def insert(self, info_data):
        query_cd = "INSERT INTO test.user (password, username, max_todo) VALUES ('{}', '{}',{})".format(info_data.get(
            "password"), info_data.get("username"), info_data.get("max_todo"))
        print(query_cd)
        try:
            self.db_cursor.execute(query_cd)
            self.db.commit()
            self.db_cursor.close()
            return True
        except Exception as e:
            return str(e)

    def update(self, info_data):
        query_cd = "UPDATE test.user SET max_todo = {} WHERE username = '{}'".format(info_data.get("max_todo"),
                                                                                     info_data.get("username"))
        try:
            self.db_cursor.execute(query_cd)
            self.db.commit()
            self.db_cursor.close()
            return True
        except Exception as e:
            return e


# a = UserModel()
# data = {
#     "username": "nguyen duc quan",
#     "password": "quan",
#     "max_todo": 100
# }
# r = a.get_user_by_name("nguyen duc quan")
# print(r)
