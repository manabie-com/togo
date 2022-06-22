from django.conf.urls import url
from ..views.user_detail_task import CreateDetail

urlpatterns = [
    url(r'^assignment/$', CreateDetail.as_view()),
]
