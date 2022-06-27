from flask_bcrypt import generate_password_hash, check_password_hash
from mongoengine.document import Document
from mongoengine.fields import StringField


class User(Document):
    username = StringField(required=True, unique=True)
    password = StringField(required=True, min_length=6)

    meta = {"db_alias": "app-db"}

    def hash_password(self):
        self.password = generate_password_hash(self.password).decode("utf8")

    def check_password(self, password):
        return check_password_hash(self.password, password)
