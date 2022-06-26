from django.db import models
from baseuser.models import BaseUser


class Task(models.Model):

    priority_options = ((1, "High"), (2, "Medium"), (3, "Low"))
    status_options = ((1, "Ready"), (2, "Active"), (3, "Done"))

    title = models.CharField(max_length=150)
    description = models.TextField(max_length=500)
    created_by = models.ForeignKey(
        BaseUser, on_delete=models.CASCADE, related_name="tasks"
    )
    priority = models.PositiveIntegerField(choices=priority_options, default=3)
    status = models.PositiveIntegerField(choices=status_options, default=1)
    created_at = models.DateTimeField(auto_now_add=True)

    class Meta:
        db_table = "tasks"
        ordering = ("-created_at",)
