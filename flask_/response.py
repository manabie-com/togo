from helpers.response import AbstractResponse


class HttpResponse(AbstractResponse):
    def __init__(self, status_code=200, data=[], message="", **kwargs):
        self.status_code = status_code
        self.data = data["data"]
        self.message = message
        for k, v in kwargs.items():
            self.__dict__[k] = v

    def keys(self):
        return self.__dict__.keys()

    def __getitem__(self, key):
        if key not in self.__dict__:
            raise KeyError
        return self.__dict__[key]

    def __iter__(self):
        yield {"message": self.message, "data": self.data}
        yield self.status_code
