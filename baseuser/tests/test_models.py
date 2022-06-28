from django.test import TestCase
from baseuser.models import BaseUser

from utils import encrypting


class TestModels(TestCase):
    def setUp(self):
        self.user = BaseUser.objects.create(
            username="user_test_001", password="Aa123456"
        )

    def test_user_is_not_superuser_on_creation(self):
        print("Running test_user_is_not_superuser_on_creation")
        self.assertNotEquals(self.user.is_superuser, True)

    def test_encrypt_and_decrypt_user_id(self):
        print("Running test_encrypt_and_decrypt_user_id")
        encrypted_number = encrypting.encrypt(self.user.id)
        self.assertNotIsInstance(encrypted_number, int)
        decrypted_number = encrypting.decrypt(encrypted_number)
        self.assertEquals(decrypted_number[0], self.user.id)
