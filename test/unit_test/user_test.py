import unittest
import requests
import json

from controllers.user_login_controller import LoginController
from controllers.user_register_controller import UserRegisterController

USER_REGISTER = "http://127.0.0.1:8000/user_register"
USER_LOGIN = "http://127.0.0.1:8000/login"


class testUser(unittest.TestCase):
    def test_register_true(self):
        user = {
                    "username": "an",
                    "password": "aaaaa",
                }
        result = UserRegisterController().register(data=user)
        self.assertEqual(result.get("status"), 200)

    def test_register_false(self):
        user = {
                    "username": "ngoc",
                    "password": "abcd"
                }
        result = UserRegisterController().register(data=user)
        self.assertEqual(result.get("status"), 400)

    def test_login_true(self):
        user = {
            "username": "ngoc",
            "password": "aaaaa"
        }
        result = LoginController().login(data=user)
        self.assertEqual(result.get("status"), 200)

    def test_login_false(self):
        user = {
            "username": "ngoc",
            "password": "1234",
            "max_todo": "abcd"
        }
        result = LoginController().login(data=user)
        self.assertEqual(result.get("status"), 400)


if __name__ == "__main__":
    unittest.main()

