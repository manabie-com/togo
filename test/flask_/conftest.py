import os
from unittest.mock import patch

import pytest
from celery.app.task import Task
from database.models.service import Service
from flask_ import init as api_init
from helpers.common import config_str_to_obj
from mongoengine.connection import connect


@pytest.fixture
def add_mock_services_to_db():
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

    yield None

    celery_config = config_str_to_obj("celery_.config", "Mock")
    connect(host=celery_config.APP_DB_URL, alias="app-db")

    Service.drop_collection()


@pytest.fixture
def patch_apply_async():
    with patch(
        "celery.app.task.Task.apply_async",
        new=Task.apply,
    ):
        yield


@pytest.fixture
def client():
    RUN_MODE = os.getenv("APP_CONFIG_DEFAULT", "Mock")
    app = api_init.factory(RUN_MODE, "test_app")

    app.config["TESTING"] = True

    with app.app_context():
        yield client


@pytest.fixture
def test_client():
    RUN_MODE = os.getenv("APP_CONFIG_DEFAULT", "Mock")
    app = api_init.factory(RUN_MODE, "test_app")

    app.config["TESTING"] = True

    return app.test_client()
