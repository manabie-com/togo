from celery.schedules import crontab

TASK_LIMITED_EACH_DAY = crontab(minute=0, hour=1, day_of_week=1)
