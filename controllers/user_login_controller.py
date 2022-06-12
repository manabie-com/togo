import os
import jwt
from config import config
from model.mysqldb.task_model import TaskModel
from model.mysqldb.user_model import UserModel


class LoginController:
    def login(self, data):
        username = data.get("username")
        password = data.get("password")
        user_info = UserModel().get_user_by_name(username=username)
        result = {"token": ""}
        if user_info:
            if password == user_info.get("password"):
                token = jwt.encode(
                            payload=data,
                            key=config.PASSWORD
                        )
                result["status"] = 200
                result["message"] = "SUCCESS!"
                result["token"] = token
            else:
                result["status"] = 400
                result["message"] = "password is incorrect!"
        else:
            result["status"] = 400
            result["message"] = "username not found!"
        return result
