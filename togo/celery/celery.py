from __future__ import absolute_import
import os
from celery import Celery
from django.conf import settings
from apps.tasks.task_limited_each_day import test1, test2
from celery.schedules import crontab

os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'togo.settings')
include_tasks = (
    'apps.tasks.task_limited_each_day'
)

app = Celery(include=include_tasks, enable_utc=False, broker='redis://localhost:6379')
app.config_from_object('django.conf:settings')
app.autodiscover_tasks(lambda: settings.INSTALLED_APPS)


@app.on_after_configure.connect
def setup_periodic_tasks(sender, **kwargs):
    sender.add_periodic_task(
        crontab(hour=1, minute=16, day_of_week=1),
        test.s('Happy Mondays!'),
    )


@app.task
def test(arg):
    print(arg)
    return arg


if __name__ == '__main__':
    app.start()
