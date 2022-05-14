from django.urls import path
from .views import UserRecordTasks, UserRecordTaskDetail, UserTaskAllows, UserTaskAllow_Detail


urlpatterns = [
    path('userrecords/', UserRecordTasks.as_view()),
    path('userrecord/<int:pk>/', UserRecordTaskDetail.as_view()),
    path('usertaskallows/', UserTaskAllows.as_view()),
    path('usertaskallow/<int:pk>/', UserTaskAllow_Detail.as_view())
]