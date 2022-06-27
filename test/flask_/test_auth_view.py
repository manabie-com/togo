import os
import sys
import unittest
from unittest import mock

import pytest
from flask_ import init as api_init

os.environ["APP_CONFIG_DEFAULT"] = "Mock"
sys.path.append(os.getcwd())


@pytest.mark.usefixtures("client")
@pytest.mark.usefixtures("test_client")
@pytest.mark.usefixtures("add_mock_services_to_db")
class AuthViewTest(unittest.TestCase):
    maxDiff = None

    def setUp(self):
        RUN_MODE = os.getenv("APP_CONFIG_DEFAULT", "Mock")
        app = api_init.factory(RUN_MODE, "test_app")
        app.config["TESTING"] = True
        with app.app_context():
            self.client = app.test_client()

    @mock.patch("flask_.blueprints.auth.view.controller.post_signup")
    def test_post_signup(self, mock_post_signup):
        mock_post_signup.return_value = {}
        request_data = {"username": "user1", "password": "test"}
        response = self.client.post("/auth/signup", json=request_data)
        self.assertEqual(response.status_code, 200)

    @mock.patch("flask_.blueprints.auth.view.controller.post_signin")
    def test_post_signin(self, mock_post_signin):
        mock_post_signin.return_value = {}
        request_data = {"username": "user1", "password": "test"}
        response = self.client.post("/auth/signin", json=request_data)
        self.assertEqual(response.status_code, 200)
