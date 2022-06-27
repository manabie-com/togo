from mongoengine.document import Document
from mongoengine.fields import IntField, ListField, StringField


class Service(Document):
    app = StringField(required=True)
    code = IntField(required=True)
    type = StringField(default="worker")
    queues = StringField(required=True)
    concurrency = IntField(default=4)
    backend = StringField(required=True)
    broker = StringField(required=True)
    loglevel = StringField(default="INFO")
    accept_content = StringField()
    result_serializer = ListField(StringField())
    task_serializer = StringField()

    meta = {"db_alias": "app-db"}
