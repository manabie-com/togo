import os
import sys
import unittest
from unittest import mock

import pytest
from celery_.response import ServiceResponse
from flask_.response import HttpResponse

os.environ["APP_CONFIG_DEFAULT"] = "Mock"
sys.path.append(os.getcwd())


@pytest.mark.usefixtures("client")
@pytest.mark.usefixtures("test_client")
@pytest.mark.usefixtures("add_mock_services_to_db")
class AuthControllerTest(unittest.TestCase):
    def setUp(self):
        from flask_.blueprints.auth import controller

        self.auth_controller = controller

    @mock.patch("flask_.blueprints.auth.controller.chain")
    def test_post_signup(self, mock_chain):
        signup_data = {"username": "user1", "password": "test"}
        mock_chain().apply_async.return_value.get.return_value = ServiceResponse()

        result = self.auth_controller.post_signup(signup_data)
        self.assertIs(type(result), HttpResponse)

    @mock.patch("celery_.services.auth_service.post_signin")
    def test_post_signin(self, mock_service):
        signin_data = {"username": "user1", "password": "test"}
        mock_service.apply_async.return_value.get.return_value = ServiceResponse()

        result = self.auth_controller.post_signin(signin_data)
        self.assertIs(type(result), HttpResponse)
