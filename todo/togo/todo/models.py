from django.db import models

class User(models.Model):
    id = models.CharField(primary_key=True, max_length=100, null=False)
    password = models.CharField(max_length=100, null=False)
    max_todo = models.IntegerField(null=False)

class Task(models.Model):
    id = models.CharField(primary_key=True, max_length=100, null=False)
    content = models.CharField(max_length=300, null=False)
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    created_date = models.DateField(null=False)
