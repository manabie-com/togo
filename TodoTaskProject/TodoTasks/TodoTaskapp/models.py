from django.db import models
from django.contrib.auth.models import User
from datetime import datetime, timezone
from django.utils import timezone


# Create your models here.
class UserRecordTask(models.Model):
    user = models.ForeignKey(User, on_delete=models.CASCADE, blank=False, null=False)
    TaskTitle = models.CharField(max_length=50)
    TaskDescription = models.CharField(max_length=500)
    TaskDay = models.DateTimeField(default=datetime.now, blank=True)

    def get_record(self):
        return str(self.user) + ' task: ' + str(self.TaskTitle) + ' ' + str(self.TaskDay)

    def __repr__(self):
        return str(self.TaskTitle) + ' record by ' + str(self.user)        


class UserTaskAllow(models.Model):
    user = models.ForeignKey(User, on_delete=models.CASCADE, blank=False, null=False)
    task_allow = models.IntegerField()
    task_done = models.IntegerField()
    start_task_time = models.DateTimeField(default=datetime.now, blank=True)
    last_task_time = models.DateTimeField(default=datetime.now, blank=True)

    def get_taskallow(self):
        return str(self.user) + ' allow ' + str(self.task_allow) + ' task. And done ' + str(self.task_done) + ' task.'

    def __repr__(self):
        return str(self.task_allow) + ' is task limit for ' + str(self.user)

