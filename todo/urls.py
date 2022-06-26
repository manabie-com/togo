from django.urls import path

from todo.views import TasksView, TaskView

urlpatterns = [
    path("", TasksView.as_view(), name="tasks"),
    path("<str:id>/", TaskView.as_view(), name="task"),
]