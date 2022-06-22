from django.conf.urls import url
from ..views.user import CreateDetail

urlpatterns = [
    url(r'^assignment/$', CreateDetail.as_view()),
]
