from __future__ import absolute_import, unicode_literals

from celery import Celery
from togo.settings import CELERY_BROKER_URL
from apps.views.utils.constant import LIMIT_TASK

app = Celery()
app.conf.broker_url = CELERY_BROKER_URL


@app.task
def pick_limit_for_user(date_):
    from togo.logger import togo_task_pick_limit_logger
    from apps.models.schedule import Schedule
    from django.contrib.auth.models import User
    try:
        data = []
        users = User.objects.all()
        for user in users:
            schedule = Schedule(user=user, limit=LIMIT_TASK, date=date_)
            data.append(schedule)
        if data:
            Schedule.objects.bulk_create(data)
    except Exception as e:
        togo_task_pick_limit_logger.exception(e)
        return False
    else:
        togo_task_pick_limit_logger.exception("schedule was created: ", data)
        return True


@app.task
def callback_get_limit_task(success, date_):
    from togo.logger import togo_task_pick_limit_logger
    from django.contrib.auth.models import User
    from apps.models.schedule import Schedule
    if not success:
        try:
            users_schedule_success = Schedule.objects.filter(date=date_).values_list('user', flat=True)
            users = User.objects.exclude(id__in=users_schedule_success)
            data = []
            for user in users:
                schedule = Schedule(user=user, limit=LIMIT_TASK, date=date_)
                data.append(schedule)
            Schedule.objects.bulk_create(data)
        except Exception as e:
            togo_task_pick_limit_logger.exception(e)
