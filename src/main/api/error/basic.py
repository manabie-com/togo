from werkzeug.exceptions import HTTPException
from src.main.util import logger


class HTTPError(HTTPException):
    def __init__(self, code: int, description: str):
        self.code = code
        self.description = description


def default_error_handler(error: HTTPError):
    logger.error(str(error), exc_info=error)
    return {"error": error.name, "description": error.description}, error.code
