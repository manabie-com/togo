from django.contrib.auth.models import User
from rest_framework.generics import ListAPIView

from ..serializers.user import UserSerializer


class ListUser(ListAPIView):
    serializer_class = UserSerializer

    def get_queryset(self):
        return User.objects.all()