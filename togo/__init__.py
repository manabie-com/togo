from __future__ import absolute_import, unicode_literals
from togo.celery.celery import app as celery_app
from apps.tasks.task_limited_each_day import app as task_set_limited

__all__ = ('celery_app', 'task_set_limited')
