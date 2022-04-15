from django.urls import path

from . import views
from todo import views

urlpatterns = [
    path(
            'api/todo', 
            views.TodoTaskCreateAPI.as_view(),
            name='todo'
        )
]