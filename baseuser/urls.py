from django.urls import path

from baseuser.views import RegistrationAPIView, UserListAPIView, UserDetailAPIView

urlpatterns = [
    path("registration/", RegistrationAPIView.as_view(), name="user_registration"),
    path("", UserListAPIView.as_view(), name="user_list"),
    path("<str:id>/", UserDetailAPIView.as_view(), name="user_detail"),
]
