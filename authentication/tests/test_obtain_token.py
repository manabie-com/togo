from django.urls import reverse
from faker import Faker
from rest_framework.test import APITestCase
from rest_framework import status

from authentication.tests import UserFactory, COMMON_PASSWORD, check_token_success


class ObtainTokenTestCase(APITestCase):

    @classmethod
    def setUpClass(cls):
        super(ObtainTokenTestCase, cls).setUpClass()
        cls.faker_obj = Faker()
        cls.user_exists = UserFactory.create()
        cls.url_api = reverse('token_obtain_pair')
        cls.user_build = UserFactory.build()

    def test_obtain_token_success(self):
        print("TEST OBTAIN TOKEN SUCCESS")
        data = {
            'email': self.user_exists.email,
            'password': COMMON_PASSWORD
        }
        response = self.client.post(self.url_api, data)

        self.assertTrue(check_token_success(response.status_code, response.data))
        self.assertEqual(response.data['user']['id'], self.user_exists.id)

    def test_obtain_token_fail(self):
        print("TEST OBTAIN TOKEN FAIL")

        data = {
            'email': self.user_build.email,
            'password': self.user_build.password
        }
        response = self.client.post(self.url_api, data)

        self.assertTrue(response.status_code, status.HTTP_401_UNAUTHORIZED)
