from celery import Celery

from apps.views.utils.constant import LIMIT_TASK

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
            schedule = Schedule(user=user, limit=LIMIT_TASK)
            data.append(schedule)
        if data:
            Schedule.objects.bulk_create(data)
    except Exception as e:
        togo_task_pick_limit_logger.exception(e)
        return False
    else:
        togo_task_pick_limit_logger.exception("schedule was created: ", data)
        return True

