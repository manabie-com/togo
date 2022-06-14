import os
import unittest
import requests
import json

from controllers.task_controller import TaskController

CREATE_TASK = "http://127.0.0.1:8000/task"
GET_TASK = "http://127.0.0.1:8000/task"


class testTask(unittest.TestCase):
    def test_create_task_true(self):
        body = {
                    "name": "task16",
                    "content": "content1"
                }
        user = {
                    "id": 1,
                    "username": "ngoc",
                    "password": "abcd",
                    "max_todo": 5
                }
        result = TaskController().create_task(data=body, user_info=user)
        self.assertEqual(result["status"], 200)

    def test_create_task_false(self):
        body = {
            "name": "task6",
            "content": "content6"
        }
        user = {
            "id": 1,
            "username": "ngoc",
            "password": "abcd",
            "max_todo": 5
        }
        result = TaskController().create_task(data=body, user_info=user)
        self.assertEqual(result["status"], 400)

    def test_get_task(self):
        user = {
                    "id": 1,
                    "username": "ngoc",
                    "password": "abcd",
                    "max_todo": 5
                }
        result = TaskController().list_task(user_info=user)
        self.assertEqual(result["status"], 200)


if __name__ == "__main__":
    unittest.main()

