from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, String, Numeric

Base = declarative_base()


class BaseInterface:
    """ Interface contains some method """

    @classmethod
    def from_dict(cls, data: dict):
        """
        construct a orm object from a series instead using constructor
        :param data: Data to be assigned
        :return: the class instance.
        """
        obj = cls()
        for k, v in data.items():
            setattr(obj, str(k), v)
        return obj
