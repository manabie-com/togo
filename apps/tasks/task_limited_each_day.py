from celery import Celery

app = Celery(__name__)


@app.task
def test1(message_id):
    print(message_id)
    return message_id
@app.task
def test2(message_id):
    print(message_id)
    return message_id
