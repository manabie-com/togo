import os

from celery import Celery
from celery.schedules import crontab
from database.models.task import UserTask
from database.models.service import Service
from celery_.response import ServiceResponse
from flask_.error import InvalidUsage
from helpers.common import config_str_to_obj
from mongoengine.connection import connect

_config = config_str_to_obj("celery_.config", os.getenv("APP_CONFIG_DEFAULT", "Config"))
connect(host=_config.APP_DB_URL, alias="app-db")
_service = Service.objects.get(app="task_service")
celery_beat_schedule = {
    "make_reset_counter_request_at_0AM_job": {
        "task": "celery_.services.task_service.reset_counter_request",
        "schedule": crontab(hour=0, minute=0),
        "options": {"queue": "task"},
    }
}
_app = Celery(
    _service.app,
    broker=_service.broker,
    backend=_service.backend,
    task_serializer=_service.task_serializer,
    accept_content=_service.accept_content,
    result_serializer=_service.result_serializer,
    beat_schedule=celery_beat_schedule,
)


@_app.task()
def post_task(user_id):
    print(user_id)
    user_task = UserTask.objects.get(user_id=str(user_id))
    if user_task.request_number_per_day == user_task.limit:
        raise InvalidUsage(
            message="Reached the limit for number of requests already",
            payload="Bad request",
        )
    """
    To do anything
    """
    UserTask.objects(user_id=user_id).update(inc__request_number_per_day=1)
    return ServiceResponse()


@_app.task()
def initialize_user_task(post_signup_response):
    UserTask(
        user_id=post_signup_response.get("user_id"),
        request_number_per_day=0,
        limit=_config.USER_REQUEST_LIMIT,
    ).save()
    return ServiceResponse()


@_app.task()
def reset_counter_request():
    UserTask.objects().update(set__request_number_per_day=0)
    return ServiceResponse()
