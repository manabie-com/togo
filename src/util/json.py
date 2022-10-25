""" Implement a customer json encoder """

from json import JSONEncoder
from datetime import datetime
from decimal import Decimal


class EnhancedEncoder(JSONEncoder):
    """
    Enhanced encoder to encode datetime, Decimal object
    """

    def default(self, o):
        """
        Overriding default function of JSON encoder to dealing with datetime & decimal
        :param o: object value to be encoded
        :return: ways to encode datetime, decimal & default values
        """
        if isinstance(o, datetime):
            return o.strftime("%Y-%m-%d %H:%M:%S")
        elif isinstance(o, Decimal):
            return float(o)
        else:
            return JSONEncoder.default(self, o)
