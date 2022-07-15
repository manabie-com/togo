from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from django.contrib.auth.models import User
from togo.models import UserProfile
from togo.serializers import UserProfileSerializer


class UserView(APIView):
    """
    API endpoint that allows users to be viewed or edited.
    """
    def get(self, request):
        response_dict = { "message": "Success.", "users": UserProfileSerializer(UserProfile.objects.all(), many=True).data }
        return Response(response_dict, status=status.HTTP_200_OK)

    def post(self, request):
        user = User.objects.create(**request.data)
        user_profile = UserProfile.objects.get(user=user)
        response_dict = { "message": "User successfully created.", "user": UserProfileSerializer(user_profile).data }
        return Response(response_dict, status=status.HTTP_200_OK)