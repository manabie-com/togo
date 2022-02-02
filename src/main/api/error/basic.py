from werkzeug.exceptions import HTTPException


class HTTPError(HTTPException):
    def __init__(self, code: int, description: str):
        self.code = code
        self.description = description


def default_error_handler(error: HTTPError):
    return {"error": error.name, "description": error.description}, error.code
