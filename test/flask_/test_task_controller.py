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
class TaskControllerTest(unittest.TestCase):
    def setUp(self):
        from flask_.blueprints.task import controller

        self.task_controller = controller

    @mock.patch("celery_.services.task_service.post_task")
    def test_post_task(self, mock_service):
        user_id = "610913b828522a470b822f80"
        mock_service.apply_async.return_value.get.return_value = ServiceResponse()

        result = self.task_controller.post_task(user_id)
        self.assertIs(type(result), HttpResponse)
