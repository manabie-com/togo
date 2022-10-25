"""
Decorators as middlewares for flask
"""
from flask import request
from ..services import extract_token
from functools import wraps
from typing import Callable
from jwt import DecodeError
from src.api.error import HTTPError


def credentials_validation(func: Callable) -> Callable:
    """
    Validate request header authorization with bearer token.
    Using this decorator to verify the bearer token
    :param func:
    :return:
    """

    @wraps(func)
    def decorated(*args, **kwargs):
        """
        The decorated function, normally endpoint controller of flask
        """
        authorization = request.headers.get('authorization', None)
        if authorization:
            try:
                auth_type, auth_token = authorization.split(maxsplit=1)
                if auth_type.lower() == 'bearer' and extract_token(auth_token):
                    # extract token and forward to next route as `payload` argument
                    payload = extract_token(auth_token)
                    return func(*args, payload=payload, token=auth_token, **kwargs)
            except DecodeError:
                raise HTTPError(401, "Invalid authentication token!")

    return decorated
