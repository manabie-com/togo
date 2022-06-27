from mongoengine.document import Document
from mongoengine.fields import IntField, ObjectIdField


class UserTask(Document):
    user_id = ObjectIdField(required=True, unique=True)
    request_number_per_day = IntField()
    limit = IntField()

    meta = {
        "db_alias": "app-db",
        "auto_create_index": False,
        "index_background": True,
        "indexes": ["user_id"],
    }
