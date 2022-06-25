from django.conf.urls import url
from ..views.user_detail_task import CreateDetail

from rest_framework_simplejwt.views import (
    TokenObtainPairView,
)

urlpatterns = [
    url(r'^assignment/$', CreateDetail.as_view(), name='assignment'),
    url(r'^token/$', TokenObtainPairView.as_view(), name='token')
]
