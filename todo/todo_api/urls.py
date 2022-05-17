# todo/todo_api/urls.py : API urls.py
from django.urls import path, include
from .views import (
    TodoListApiView,
    TodoDetailApiView,
    RegisterUserAPI,
    UserDetailAPI
)

urlpatterns = [
    path('user/', UserDetailAPI.as_view(), name='user-get'),
    path('user/register/', RegisterUserAPI.as_view(), name='user-create'),
    path('task/', TodoListApiView.as_view(), name='task-list'),
    path('task/<int:todo_id>/', TodoDetailApiView.as_view(), name='task-detail'),
]