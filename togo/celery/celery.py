from __future__ import absolute_import
import os
import django

from celery import Celery
from django.conf import settings
from apps.tasks.task_limited_each_day import pick_limit_for_user
from togo.celery.celery_setting import ONE_DAY_SECONDS
from datetime import date

os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'togo.settings')
django.setup()

app = Celery('togo', enable_utc=False)
app.config_from_object('django.conf:settings', namespace='CELERY')
app.autodiscover_tasks(lambda: settings.INSTALLED_APPS)


@app.on_after_configure.connect
def setup_periodic_tasks(sender, **kwargs):
    sender.add_periodic_task(ONE_DAY_SECONDS, pick_limit_for_user.s(str(date.today())), expires=10)


if __name__ == '__main__':
    app.start()
