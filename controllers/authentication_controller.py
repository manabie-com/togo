import jwt
from config import config
from model.mysqldb.task_model import TaskModel
from model.mysqldb.user_model import UserModel


class Authentication:
    def validate_token(self, token):
        data_info = jwt.decode(token, key=config.PASSWORD, algorithms=['HS256', ])
        user_info = UserModel().get_user_by_name(username=data_info.get("username"))
        if data_info.get("password") == user_info.get("password"):
            return user_info
        return False
