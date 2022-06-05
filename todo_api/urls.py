
from django.urls import path, include
from django.contrib import admin
from .views import (
    UserDetailAPI,
    RegisterUserAPI,
    TodoListAPI,
    TodoDetailApi
)


urlpatterns = [
    path('admin/', admin.site.urls),
    path('user/', UserDetailAPI.as_view(), name='user-get'),
    path('user/register/', RegisterUserAPI.as_view(), name='user-create'),
    path('todo/', TodoListAPI.as_view(), name='todo-list'),
    path('todo/<int:todo_id>/', TodoDetailApi.as_view(), name='todo-detail'),
]