from django.conf.urls import url
from ..views.user import ListUser

urlpatterns = [
    url(r'^users/$', ListUser.as_view()),
]
