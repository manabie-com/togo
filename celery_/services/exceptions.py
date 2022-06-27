class BaseException(Exception):
    code = 0
    description = None

    def __init__(self, message, payload=None):
        self.message = message
        self.payload = payload
        super().__init__(self.message, self.payload)


class ConflictError(BaseException):
    code = 409


class ObjectNotFoundError(BaseException):
    code = 404
