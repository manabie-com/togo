import os
import sys
import unittest

from database.models.service import Service
from helpers.common import config_str_to_obj
from mongoengine.connection import connect

sys.path.append(os.getcwd())
os.environ["APP_CONFIG_DEFAULT"] = "Mock"


class BaseTestCase(unittest.TestCase):
    maxDiff = None

    def add_mock_services_to_db(self):
        celery_config = config_str_to_obj("celery_.config", "Mock")
        connect(host=celery_config.APP_DB_URL, alias="app-db")
        Service.drop_collection()
        Service.objects.insert(
            [
                Service(
                    app="task_service",
                    code=1,
                    type="worker",
                    queues="task",
                    concurrency=4,
                    backend="redis://localhost",
                    broker="redis://localhost",
                    loglevel="INFO",
                    task_serializer="pickle",
                    accept_content=["pickle"],
                    result_serializer="pickle",
                ),
                Service(
                    app="auth_service",
                    code=2,
                    type="worker",
                    queues="auth",
                    concurrency=4,
                    backend="redis://localhost",
                    broker="redis://localhost",
                    loglevel="INFO",
                    task_serializer="pickle",
                    accept_content=["pickle"],
                    result_serializer="pickle",
                ),
            ]
        )

    def drop_collection(self):
        celery_config = config_str_to_obj("celery_.config", "Mock")
        connect(host=celery_config.APP_DB_URL, alias="app-db")

        Service.drop_collection()
