import uuid

from django.conf import settings
from django.db import models

# Create your models here.

TASK_STATUS = (
    ('PD', 'Pending'),
    ('DO', 'DONE'),
)


class BaseModel(models.Model):
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    class Meta:
        abstract = True
        ordering = ['-created_at']


class SettingTask(BaseModel):
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    date = models.DateField(auto_now_add=True)
    user = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.CASCADE, related_name="settings")
    limit = models.IntegerField(default=5)

    class Meta:
        unique_together = ('user', 'date')


class Task(BaseModel):
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    setting = models.ForeignKey(SettingTask, on_delete=models.CASCADE, related_name='setting_tasks')
    is_accepted = models.BooleanField(default=True)
    name = models.CharField(max_length=150, null=False, blank=False)
    description = models.TextField(null=True, blank=True)
    status = models.CharField(choices=TASK_STATUS, default='PD', max_length=10)

    def __str__(self):
        return self.name
