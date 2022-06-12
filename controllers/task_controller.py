import jwt
import json
from flask import Flask, request, Response
from model.mysqldb.task_model import TaskModel
from model.mysqldb.user_model import UserModel
from datetime import datetime, time


class TaskController:
    def create_task(self, data, user_info):
        max_todo = user_info.get("max_todo")
        data_insert = {
            "user_id": user_info.get("id"),
            "content": data.get("content"),
            "name": data.get("name"),
            "created_date": int(datetime.now().timestamp())
        }
        all_task = TaskModel().get_task_by_user_id(user_id=user_info.get("id"),
                                                   from_time=datetime.combine(datetime.today(), time.min).timestamp())
        result = {}
        if len(all_task) < max_todo:
            result_insert = TaskModel().insert(data_insert)
            if result_insert is True:
                result["status"] = 200
                result["message"] = "SUCCESS!"
            else:
                result["status"] = 400
                result["message"] = result
        else:
            result["status"] = 400
            result["message"] = "Could not create more than {} tasks in a day!".format(max_todo)
        return result

    def list_task(self, user_info):
        all_task = TaskModel().get_task_by_user_id(user_id=user_info.get("id"), from_time=0)
        result = {
            "status": 200,
            "message": "SUCCESS!",
            "data": all_task
        }
        return result
