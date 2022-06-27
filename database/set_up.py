from database.models.service import Service
from helpers.common import config_str_to_obj
from mongoengine.connection import connect


def add_services_to_db(self):
    celery_config = config_str_to_obj("celery_.config", "Config")
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
