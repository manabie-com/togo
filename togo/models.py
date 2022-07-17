from django.contrib.auth.models import User
from django.db import models
from django.dispatch import receiver

class UserProfile(models.Model):
    user = models.OneToOneField(User, on_delete=models.CASCADE)
    task_limit = models.IntegerField(default=5)

class Task(models.Model):
    user = models.ForeignKey(User, on_delete=models.CASCADE)
    name = models.CharField(max_length=255, blank=True)

# Automatically create a UserProfile when a User is created
@receiver(models.signals.post_save, sender=User)
def create_user_profile_for_user(sender, instance, created, **kwargs):
    if created: UserProfile.objects.create(user=instance)