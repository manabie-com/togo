from django.db import models

class User(models.Model):
    id = models.AutoField(primary_key=True)
    username = models.CharField(max_length=100, null=False)
    password = models.CharField(max_length=100, null=False)
    max_todo = models.IntegerField(default=5)

class Task(models.Model):
    id = models.AutoField(primary_key=True)
    content = models.CharField(max_length=300, null=False)
    user_id = models.ForeignKey(User, on_delete=models.CASCADE)
    date_created = models.DateTimeField(auto_now_add=True)