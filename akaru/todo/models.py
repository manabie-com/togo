from re import L
from django.contrib.auth.models import User
from django.dispatch import receiver
from django.db import models
from django.db.models.signals import post_save


# Create your models here.
class UserProfile(models.Model):
    user = models.OneToOneField(User, on_delete=models.CASCADE)
    limit = models.PositiveIntegerField()
    
    def __str__(self) -> str:
        return self.user.username

class TodoTask(models.Model):
    date_created = models.DateTimeField(auto_now_add=True)
    title = models.CharField(max_length=30)
    text = models.CharField(max_length=300)
    user = models.ForeignKey('UserProfile', on_delete=models.CASCADE)
