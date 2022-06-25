from django.db import models
from django.contrib.auth.models import User
from django.utils import timezone


class Schedule(models.Model):
    objects = models.Manager()
    date = models.DateField(default=timezone.now)
    limit = models.IntegerField(blank=False, null=False)
    user = models.ForeignKey(User, related_name='schedules', blank=False, null=False, on_delete=models.CASCADE)

    class Meta:
        db_table = 'schedule'
        unique_together = (('date', 'user'),)
        index_together = (('date', 'user'),)
        # app_label = 'apps.models.models.schedule.Schedule'
