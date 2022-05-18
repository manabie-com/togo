from django.urls import path

from authentication import views
from authentication.views import MyTokenObtainPairView, MyTokenRefreshView

urlpatterns = [
    path('registration', views.RegisterView.as_view(), name='registration'),
    path('token', MyTokenObtainPairView.as_view(), name='token_obtain_pair'),
    path('token/refresh', MyTokenRefreshView.as_view(), name='token_refresh'),
    path('profile', views.UserProfileView.as_view(), name="user_profile"),
    path('logout', views.LogoutView.as_view(), name="logout")
]