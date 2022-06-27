from marshmallow import RAISE, Schema
from marshmallow.fields import Str


class SignInSchema(Schema):
    class Meta:
        unknown = RAISE

    username = Str()
    password = Str()


class SignUpSchema(Schema):
    class Meta:
        unknown = RAISE

    username = Str()
    password = Str()
