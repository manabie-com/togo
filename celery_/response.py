import json

from helpers import response


class ServiceResponse(response.AbstractResponse, dict):
    def __init__(self, message="", **kwargs):
        self.message = message
        for k, v in kwargs.items():
            self.__dict__[k] = v
        dict.__init__(self, **self.__dict__)

    def keys(self):
        return self.__dict__.keys()

    def __getitem__(self, key):
        if key not in self.__dict__:
            raise KeyError
        return self.__dict__[key]

    def __iter__(self):
        resp = {"data": {}}

        for k, v in self.__dict__.items():
            if k != "message":
                resp["data"][k] = v
        yield {"message": self.message}
        yield resp

    def default(self, o):
        return json.dumps(o.__dict__, cls=o)
