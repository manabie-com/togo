from django.urls import reverse
from rest_framework import status
from rest_framework_simplejwt.tokens import RefreshToken

from authentication.tests import UserFactory, TestCaseBase, check_token_success


class TokenRefreshTestCase(TestCaseBase):

    @classmethod
    def setUpClass(cls):
        super(TokenRefreshTestCase, cls).setUpClass()
        cls.api_url = reverse('token_refresh')
        cls.user_exists = UserFactory.create()
        cls.user_build = UserFactory.build()
        cls.refresh_obj = RefreshToken.for_user(cls.user_exists)

    def test_token_refresh_success(self):
        print("TEST TOKEN REFRESH SUCCESS")

        data = {
            'refresh': str(self.refresh_obj),
        }
        response = self.client.post(self.api_url, data)

        self.assertTrue(check_token_success(response.status_code, response.data))
        self.assertEqual(response.data['user']['id'], self.user_exists.id)

    def test_token_refresh_fail(self):
        print("TEST TOKEN REFRESH FAIL")

        data = {
            'refresh': self.faker_obj.pystr()
        }
        response = self.client.post(self.api_url, data)

        self.assertTrue(response.status_code, status.HTTP_401_UNAUTHORIZED)
