from functools import wraps
from django.core.exceptions import ObjectDoesNotExist

from apps.exceptions.status_code import Code500
from apps.models.models.schedule import Schedule
from apps.views.utils.constant import LIMIT_TASK
from apps.tasks.task_limited_each_day import pick_limit_for_user
from togo.logger.base import togo_task_pick_limit_logger
from django.contrib.auth.models import User
from celery import Celery

app = Celery(__name__)


@app.task
def callback_get_limit_task(success, date_):
    if not success:
        try:
            users_schedule_success = Schedule.objects.filter(date=date_).value_list('user', flat=True)
            users = User.objects.exclude(id=users_schedule_success)
            data = []
            for user in users:
                schedule = Schedule(user=user, limit=LIMIT_TASK)
                data.append(schedule)
            Schedule.objects.bulk_create(data)
        except Exception as e:
            raise Code500


def retry_get_limit_task(func):
    @wraps(func)
    def execute(*args, **kwargs):
        try:
            return func(*args, **kwargs)
        except ObjectDoesNotExist:
            (user, date_) = args if args else (kwargs.get('user'), kwargs.get('date'))
            schedule = Schedule.objects.create(limit=LIMIT_TASK, user=user)
            togo_task_pick_limit_logger.exception("something wrong with Schedule celery create generally schedule: ",
                                                  schedule.id)
            # pick_limit_for_user.apply_async(link=callback_get_limit_task.s())
            return schedule
        except Exception as e:
            raise Code500

    return execute
