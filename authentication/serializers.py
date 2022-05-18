from django.conf import settings
from rest_framework import serializers
from rest_framework_simplejwt.authentication import JWTAuthentication
from rest_framework_simplejwt.serializers import TokenObtainPairSerializer, TokenRefreshSerializer, \
    TokenVerifySerializer
from rest_framework_simplejwt.tokens import RefreshToken

from authentication.models import User


class UserSerializer(serializers.ModelSerializer):

    class Meta:
        model = User
        fields = ['id', 'email', 'password', 'first_name', 'last_name']
        extra_kwargs = {
            'password': {'write_only': True},
        }


class MyTokenObtainPairSerializer(TokenObtainPairSerializer):

    def validate(self, attrs):
        data = super().validate(attrs)
        refresh = self.get_token(self.user)
        user_serializer = UserSerializer(self.user)
        data['user'] = user_serializer.data
        data['token_expire_at'] = refresh.access_token.get('exp')
        data['refresh_token_expire_at'] = refresh.get('exp')
        return data


class MyTokenRefreshSerializer(TokenRefreshSerializer):

    def validate(self, attrs):
        refresh = RefreshToken(attrs['refresh'])

        data = {'access': str(refresh.access_token)}

        if settings.SIMPLE_JWT['ROTATE_REFRESH_TOKENS']:
            if settings.SIMPLE_JWT['BLACKLIST_AFTER_ROTATION']:
                try:
                    # Attempt to blacklist the given refresh token
                    refresh.blacklist()
                except AttributeError:
                    # If blacklist app not installed, `blacklist` method will
                    # not be present
                    pass

            refresh.set_jti()
            refresh.set_exp()

            data['refresh'] = str(refresh)

        jwt_obj = JWTAuthentication()
        user = jwt_obj.get_user(refresh.access_token)
        user_serializer = UserSerializer(user)
        data['user'] = user_serializer.data
        data['token_expire_at'] = refresh.access_token.get('exp')
        data['refresh_token_expire_at'] = refresh.get('exp')
        return data


class MyTokenVerifySerializer(TokenVerifySerializer):

    def validate(self, attrs):
        data = super().validate(attrs)
        jwt_object = JWTAuthentication()
        validated_token = jwt_object.get_validated_token(attrs.get('token'))
        user = jwt_object.get_user(validated_token)
        user_serializer = UserSerializer(user)
        data['user'] = user_serializer.data
        return data
