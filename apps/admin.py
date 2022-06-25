from django.contrib import admin
from apps.models.schedule import Schedule
from apps.models.detail import Detail
from apps.models.task import Task


class AuthorAdmin(admin.ModelAdmin):
    pass


admin.site.register(Schedule, AuthorAdmin)
admin.site.register(Detail, AuthorAdmin)
admin.site.register(Task, AuthorAdmin)
