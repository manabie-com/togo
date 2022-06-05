from django.contrib import admin

from . import models


class UserAdmin(admin.ModelAdmin):
    list_display = ('username','number_todo_limit')
class TodoAdmin(admin.ModelAdmin):
    list_display = (
        'title',
        'description',
        'status',
        'date',
        'priority',
        'tag'
    )

class TodoListAdmin(admin.ModelAdmin):
    list_display = (
        'id_user',
        'id_todo',
        'date'
    )

admin.site.register(models.Todo, TodoAdmin)
admin.site.register(models.CustomUser, UserAdmin)
admin.site.register(models.TodoList, TodoListAdmin)