from django.urls import path
from . import views

urlpatterns = [
    path('login', views.login, name="login"),
    path('login/', views.login, name="login"),
    path('tasks', views.tasks, name="tasks"),
    path('tasks/', views.tasks, name="tasks")
]
