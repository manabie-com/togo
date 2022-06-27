import jwt
import os

from celery import Celery
from database.models.user import User
from database.models.service import Service
from database.models.task import UserTask
from celery_.response import ServiceResponse
from flask_.error import InvalidUsage
from helpers.common import config_str_to_obj
from mongoengine.connection import connect

_config = config_str_to_obj("celery_.config", os.getenv("APP_CONFIG_DEFAULT", "Config"))
connect(host=_config.APP_DB_URL, alias="app-db")
_service = Service.objects.get(app="auth_service")
_app = Celery(
    _service.app,
    broker=_service.broker,
    backend=_service.backend,
    task_serializer=_service.task_serializer,
    accept_content=_service.accept_content,
    result_serializer=_service.result_serializer,
)


@_app.task()
def post_signup(signup_data):
    user = User(**signup_data)
    user.hash_password()
    user.save()
    return ServiceResponse(user_id=str(user.id))


@_app.task()
def post_signin(signin_data):
    user = User.objects.get(username=signin_data.get("username"))
    authorized = user.check_password(signin_data.get("password"))
    if not authorized:
        raise InvalidUsage(
            code=401, message="Your password is not correct", payload="Login fail"
        )

    jwt_token = jwt.encode(
        payload={"user_id": str(user.id)}, key="secret", algorithm="HS256"
    )
    return ServiceResponse(bearer_token=f"Bearer: {jwt_token}")
