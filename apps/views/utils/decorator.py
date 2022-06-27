from functools import wraps
from django.core.exceptions import ObjectDoesNotExist
from apps.tasks.task_limited_each_day import pick_limit_for_user, callback_get_limit_task
from apps.exceptions.status_code import Code500
from apps.models.schedule import Schedule
from apps.views.utils.constant import LIMIT_TASK
from togo.logger import togo_task_pick_limit_manually_logger
from celery import Celery

app = Celery(__name__)


def retry_get_limit_task(func):
    @wraps(func)
    def execute(*args, **kwargs):
        try:
            return func(*args, **kwargs)
        except ObjectDoesNotExist:
            (user, date_) = args if args else (kwargs.get('user'), kwargs.get('date'))
            schedule = Schedule.objects.create(limit=LIMIT_TASK, user=user, date=date_)
            togo_task_pick_limit_manually_logger.warning(
                "something wrong with Schedule celery - create manually schedule: ",
                schedule.id)
            pick_limit_for_user.apply_async((str(date_),), link=callback_get_limit_task.s(str(date_)))
            return schedule
        except Exception as e:
            raise Code500

    return execute
