from django.db import models


class Task(models.Model):
    objects = models.Manager()
    name = models.CharField(max_length=255, blank=True, null=False)

    class Meta:
        db_table = 'task'
        app_label = 'apps.models.models.task.Task'
