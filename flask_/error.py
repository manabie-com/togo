import os
import traceback
from functools import wraps

from celery_.services.exceptions import ConflictError, ObjectNotFoundError
from flask_.exceptions import APIAuthError
from flask_jwt_extended.exceptions import NoAuthorizationError, WrongTokenError
from helpers.common import config_str_to_obj
from jwt.exceptions import ExpiredSignatureError, InvalidSignatureError
from marshmallow.exceptions import ValidationError
from mongoengine.errors import DoesNotExist, NotUniqueError

_config = config_str_to_obj("flask_.config", os.getenv("APP_CONFIG_DEFAULT", "Config"))


def handle_exception(func):
    @wraps(func)
    def handle(*args, **kwargs):
        try:
            return func(*args, **kwargs)
        except InvalidUsage as iu:
            raise iu
        except APIAuthError as aae:
            raise InvalidUsage(
                code=aae.status_code, message=aae.error_msg, payload=str(aae)
            )
        except (
            ExpiredSignatureError,
            NoAuthorizationError,
            WrongTokenError,
            InvalidSignatureError,
        ) as jwt_auth_error:
            raise InvalidUsage(code=401, message=str(jwt_auth_error))
        except FileNotFoundError as fe:
            raise InvalidUsage(code=404, payload=str(fe))
        except ObjectNotFoundError as not_found:
            raise InvalidUsage(
                code=404, message=not_found.message, payload=not_found.payload
            )
        except DoesNotExist as dne:
            raise InvalidUsage(code=404, message=str(dne))
        except ConflictError as conflict:
            raise InvalidUsage(
                code=conflict.code,
                message=conflict.message,
                payload=conflict.payload,
            )
        except ValidationError as ve:
            raise InvalidUsage(code=422, message="Validation failed", payload=str(ve))
        except NotUniqueError:
            raise InvalidUsage(code=400, message="Username already exists")
        except Exception as e:
            raise InvalidUsage(code=500, payload=str(e))

    return handle


class InvalidUsage(Exception):
    _status_code = 400
    _default_error_msg = {
        400: "Bad request",
        401: "Permission denied",
        403: "Account suspended",
        404: "Not Found",
        422: "Database key error",
        500: "Internal generic error",
        503: "Service down",
    }

    def __init__(self, code=None, message=None, payload=None, tags=None):
        Exception.__init__(self)
        self.stack_trace = traceback.format_exc()
        self.payload = payload
        self.message = message
        self.tags = tags or {}

        if code:
            self._status_code = code
        if not message:
            self.message = self._default_error_msg[self._status_code]

    def to_dict(self):
        rv = dict({"payload": self.payload} or ())
        rv["message"] = self.message
        return rv
