from __future__ import absolute_import
import os
from celery import Celery
from django.conf import settings
from apps.tasks.task_limited_each_day import pick_limit_for_user
from .celery_setting import ONE_DAY_SECONDS

os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'togo.settings')
include_tasks = (
    'apps.tasks.task_limited_each_day'
)

app = Celery(include=include_tasks, enable_utc=False, backend='django-db', broker='redis://localhost:6379/0')
app.config_from_object('django.conf:settings')
app.autodiscover_tasks(lambda: settings.INSTALLED_APPS)


@app.on_after_configure.connect
def setup_periodic_tasks(sender, **kwargs):
    sender.add_periodic_task(ONE_DAY_SECONDS, pick_limit_for_user.s(), expires=10)


if __name__ == '__main__':
    app.start()
