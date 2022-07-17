from django.contrib import admin
from togo.models import *

@admin.register(UserProfile)
class UserProfileAdmin(admin.ModelAdmin):
    list_display = ("username", "email", "task_limit", "is_superuser")

    def username(self, obj):
        return obj.user.username
    
    def email(self, obj):
        return obj.user.email

    def is_superuser(self, obj):
        return obj.user.is_superuser