import random
from celery import Celery

app = Celery(__name__)


@app.task
def pick_limit_for_user():
    from togo.logger.base import togo_task_pick_limit_logger
    from apps.models.models.schedule import Schedule
    from django.contrib.auth.models import User
    try:
        data = []
        users = User.objects.all()
        for user in users:
            schedule = Schedule(user=user, limit=random.randint(0, 6))
            data.append(schedule)
        if data:
            Schedule.objects.bulk_create(data)
    except Exception as e:
        togo_task_pick_limit_logger.exception(e)
    else:
        togo_task_pick_limit_logger.exception("schedule was created: ", data)
