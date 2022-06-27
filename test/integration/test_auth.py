import jwt
import os
import sys

import requests
from database.models.user import User
from celery_.response import ServiceResponse
from celery_ import init as service_init
from flask_ import init as api_init
from flask_.error import InvalidUsage
from test.integration.base_test_case import BaseTestCase

sys.path.append(os.getcwd())


class AuthTest(BaseTestCase):
    def setUp(self):
        RUN_MODE = os.getenv("APP_CONFIG_DEFAULT", "Mock")
        self.celery_client = service_init.factory(RUN_MODE)
        with self.celery_client:
            super().add_mock_services_to_db()
            self.app = api_init.factory(RUN_MODE, "test")

    def tearDown(self):
        with self.celery_client:
            super().drop_collection()

    def test_signin(self):
        self.app.config["TESTING"] = True
        with self.app.app_context():
            test_client = self.app.test_client()
            request_data = {"username": "user1", "password": "test"}
            response = test_client.post("/auth/signin", json=request_data)
            print(response)
            self.assertEqual(response.status_code, 200)
