from django.contrib import admin
from .models import UserTaskAllow, UserRecordTask

# Register your models here.

admin.site.register(UserRecordTask)
admin.site.register(UserTaskAllow)
