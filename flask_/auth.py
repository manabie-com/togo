import os
from functools import wraps

import jwt
from flask import g, request
from flask_.exceptions import APIAuthError
from helpers import common

_config = common.config_str_to_obj(
    "celery_.config", os.getenv("APP_CONFIG_DEFAULT", "Config")
)


def auth_required():
    def wrapper(func):
        @wraps(func)
        def decorator(*args, **kwargs):
            authenticate()

            return func(*args, **kwargs)

        return decorator

    return wrapper


def authenticate():
    request_token = request.headers.get("Authorization")

    if not request_token:
        raise APIAuthError("Authorization token missing from headers.")

    if "Bearer " not in request_token:
        raise APIAuthError("Invalid Authorization token format.")

    token_data = jwt.decode(
        jwt=request_token.replace("Bearer ", ""),
        key="secret",
        algorithms=["HS256"],
    )
    g.user_id = token_data.get("user_id")
    g.jwt = request_token

    if not g.user_id:
        raise APIAuthError("Unable to get user_id from this token.")
