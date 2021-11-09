from django.urls import path
from . import views

urlpatterns = [
    path('login/', views.login),
    path('tasks/', views.tasks)
]
