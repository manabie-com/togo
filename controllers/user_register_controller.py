from model.mysqldb.user_model import UserModel


class UserRegisterController:
    def register(self, data):
        result = {}
        if "username" not in data:
            result["status"] = 400
            result["message"] = "username not found!"
            return result
        if "password" not in data:
            result["status"] = 400
            result["message"] = "password not found!"
            return result
        if "max_todo" not in data:
            data["max_todo"] = 5
        user_info = UserModel().insert(info_data=data)
        if user_info is True:
            result["status"] = 200
            result["message"] = "SUCCESS!"
        else:
            result["status"] = 400
            result["message"] = user_info
        return result
