from django.db import models


# User model for neater data

# The functionality to insert to this table is outside the scope of the requirement
# We are only expecting select statements to be used
class User(models.Model):
    id = models.AutoField(primary_key=True)
    username = models.CharField(max_length=20,unique=True)
    daily_limit = models.IntegerField()
    task_today = models.IntegerField(default=0)

# UserTask model for neater data

# Associates a particular task to a user
class UserTask(models.Model):
    id = models.AutoField(primary_key=True)
    user = models.ForeignKey(User, on_delete=models.CASCADE, default=1)
    added_time = models.DateTimeField()
    is_active = models.BooleanField(default=True)

# Task model for the actual task data to be sent via API

# Requirement only specified the "task"
# But additional fields were created with expansions in mind
class Task(models.Model):
    task = models.ForeignKey(UserTask, on_delete=models.CASCADE, default=1)
    title = models.CharField(max_length=50)
    description = models.CharField(max_length=100,null=True)
    start_time = models.DateTimeField(null=True)
    end_time = models.DateTimeField(null=True)