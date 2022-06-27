import os
import sys
import unittest
from unittest import mock

import pytest
from flask_ import init as api_init
import jwt

os.environ["APP_CONFIG_DEFAULT"] = "Mock"
sys.path.append(os.getcwd())


@pytest.mark.usefixtures("client")
@pytest.mark.usefixtures("test_client")
@pytest.mark.usefixtures("add_mock_services_to_db")
class TaskViewTest(unittest.TestCase):
    maxDiff = None

    def setUp(self):
        RUN_MODE = os.getenv("APP_CONFIG_DEFAULT", "Mock")
        app = api_init.factory(RUN_MODE, "test_app")
        app.config["TESTING"] = True
        with app.app_context():
            self.client = app.test_client()

    @mock.patch("flask_.blueprints.task.view.controller.post_task")
    def test_post_task(self, mock_post_task):
        mock_post_task.return_value = {}
        user_id = "610913b828522a470b822f80"
        jwt_token = jwt.encode(
            payload={"user_id": user_id}, key="secret", algorithm="HS256"
        )
        access_token = {"Authorization": f"Bearer {jwt_token}"}
        response = self.client.post("/task", headers=access_token)
        self.assertEqual(response.status_code, 200)

    @mock.patch("flask_.blueprints.task.view.controller.post_task")
    def test_post_task_no_jwt(self, mock_post_task):
        mock_post_task.return_value = {}

        response = self.client.post("/task")
        self.assertEqual(response.status_code, 401)
