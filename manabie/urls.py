from django.urls import path, include
from rest_framework_simplejwt import views as jwt_views


urlpatterns = [
    path(
        "api/login/", jwt_views.TokenObtainPairView.as_view(), name="token_obtain_pair"
    ),
    path(
        "api/token/refresh/", jwt_views.TokenRefreshView.as_view(), name="token_refresh"
    ),
    path("api/tasks/", include("todo.urls"), name="tasks"),
    path("api/users/", include("baseuser.urls"), name="users"),
]
