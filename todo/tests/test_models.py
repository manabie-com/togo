from django.test import TestCase
from django.db import IntegrityError
from todo.models import Task
from baseuser.models import BaseUser

from utils import encrypting


class TestModels(TestCase):
    def setUp(self):
        self.task = Task(
            title="adding a new task", description="trying to add a new task"
        )
        self.user = BaseUser.objects.create(
            username="user_test_001", password="Aa123456"
        )
        self.number = 2

    def test_task_can_not_added_without_created_by_on_creation(self):
        print("Running test_task_can_not_added_without_created_by_on_creation")
        try:
            self.task = self.task.save()
            self.task.clean_fields()
        except:
            self.assertRaises(IntegrityError)

    def test_encrypt_and_decrypt_task_id(self):
        print("Running test_encrypt_and_decrypt_task_id")
        self.task = Task.objects.create(
            title="adding a new task",
            description="trying to add a new task",
            created_by=self.user,
        )
        encrypted_number = encrypting.encrypt(self.task.id)
        self.assertNotIsInstance(encrypted_number, int)
        decrypted_number = encrypting.decrypt(encrypted_number)
        self.assertEquals(decrypted_number[0], self.task.id)
