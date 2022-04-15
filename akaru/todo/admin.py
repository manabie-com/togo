from django.contrib import admin
from django.contrib.auth.admin import UserAdmin as BaseUserAdmin
from django.contrib.auth.models import User

from todo.models import UserProfile
from todo.models import TodoTask


class UserProfileInline(admin.StackedInline):
    model = UserProfile
    can_delete = False
    verbose_name_plural = 'user profiles'


class UserAdmin(BaseUserAdmin):
    inlines = (UserProfileInline,)


# Register your models here.
admin.site.unregister(User)
admin.site.register(User, UserAdmin)
admin.site.register(TodoTask)