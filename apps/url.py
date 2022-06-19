from django.urls import include, re_path

from rest_framework import permissions
from drf_yasg.views import get_schema_view
from drf_yasg import openapi

schema_view = get_schema_view(
   openapi.Info(
      title="API",
      default_version="v1",
      description="Description: The system api developing for tiktok and facebook",
      contact=openapi.Contact(email="trinct1412@gmail.com"),
       # terms_of_service="https://www.google.com/policies/terms/",
       # license=openapi.License(name="BSD License"),
   ),
   public=True,
   permission_classes=(permissions.AllowAny,),
)


urlpatterns = [
    # define url swagger
    re_path(r'^docs/$', schema_view.with_ui('swagger', cache_timeout=0), name='schema-swagger-ui'),
    re_path(r'^redocs/$', schema_view.with_ui('redoc', cache_timeout=0), name='schema-redoc'),

    # define url api apps
    re_path(r'^users/', include('apps.urls.user')),
]
