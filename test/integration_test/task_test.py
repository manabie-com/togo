import unittest
import requests
import json

CREATE_TASK = "http://127.0.0.1:8000/task"
GET_TASK = "http://127.0.0.1:8000/task"


class testTask(unittest.TestCase):
    def test_create_task_true(self):
        headers = {
            "Authorization": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VybmFtZSI6Im5hbSIsInBhc3N3b3J"
                             "kIjoiYWFhYWEifQ.f12_9TkDfd0Pxv1lH6MngaCxkZgsmQBS_oIVDE9m_dw"
        }
        body = {
                    "name": "an",
                    "content": "aaaaa",
                }
        result = requests.post(url=CREATE_TASK, json=body, headers=headers)
        print(result.text)
        self.assertEqual(result.status_code, 200)

    def test_create_task_false(self):
        headers = {
            "Authorization": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VybmFtZSI6Im5hbSIsInBhc3N3b3J"
                             "kIjoiYWFhYWEifQ.f12_9TkDfd0Pxv1lH6MngaCxkZgsmQBS_oIVDE9m_dw"
        }
        body = {
                    "name": "ngoc",
                    "content": "abcd"
                }
        result = requests.post(url=CREATE_TASK, json=body, headers=headers)
        print(result.text)
        self.assertEqual(result.status_code, 400)

    def test_get_task(self):
        headers = {
            "Authorization": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VybmFtZSI6Im5hbSIsInBhc3N3b3J"
                             "kIjoiYWFhYWEifQ.f12_9TkDfd0Pxv1lH6MngaCxkZgsmQBS_oIVDE9m_dw"
        }
        result = requests.get(url=CREATE_TASK, headers=headers)
        print(result.text)
        self.assertEqual(result.status_code, 200)


if __name__ == "__main__":
    unittest.main()

