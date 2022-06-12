import unittest
import requests
import json

USER_REGISTER = "http://127.0.0.1:8000/user_register"
USER_LOGIN = "http://127.0.0.1:8000/login"


class testUser(unittest.TestCase):
    def test_register_true(self):
        user = {
                    "username": "hung",
                    "password": "aaaaa",
                }
        result = requests.post(url=USER_REGISTER, json=user)
        print(result.text)
        self.assertEqual(result.status_code, 200)

    def test_register_false(self):
        user = {
                    "username": "ngoc",
                    "password": "abcd",
                    "max_todo": "abcd"
                }
        result = requests.post(url=USER_REGISTER, json=user)
        print(result.text)
        self.assertEqual(result.status_code, 400)

    def test_login_true(self):
        user = {
            "username": "ngoc",
            "password": "aaaaa"
        }
        result = requests.post(url=USER_LOGIN, json=user)
        print(result.text)
        self.assertEqual(result.status_code, 200)

    def test_login_false(self):
        user = {
            "username": "ngoc",
            "password": "1234"
        }
        result = requests.post(url=USER_LOGIN, json=user)
        print(result.text)
        self.assertEqual(result.status_code, 400)


if __name__ == "__main__":
    unittest.main()

