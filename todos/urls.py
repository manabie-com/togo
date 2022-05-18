from django.urls import path, include

from todos.views import TaskView, TaskDetailView, TaskSetting

urlpatterns = [
    path('auth/', include('authentication.urls')),
    path('tasks', TaskView.as_view(), name="tasks"),
    path('tasks/<int:pk>', TaskDetailView.as_view(), name="detail_task"),
    path('task-settings', TaskSetting.as_view(), name="task-settings"),
]
