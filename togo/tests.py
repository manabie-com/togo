from rest_framework.test import APITestCase
from rest_framework import status
from django.test import TestCase
from django.contrib.auth.models import User
from togo.models import *

class UserTestCase(TestCase):
    # Test if UserProfile for user is automatically created upon registration
    def test_create_user_profile(self):
        user = User.objects.create(username="test_user", password="test_password")
        self.assertIsNotNone(UserProfile.objects.filter(user=user).first())        

class UsersAPITestCase(APITestCase):
    # Test if user can register properly
    def test_user_registration(self):
        request_data = {
            "username": "test_user",
            "password": "test_password"
        }
        response = self.client.post("/api/users/", request_data, format="json")
        self.assertEqual(response.status_code, status.HTTP_201_CREATED)
        self.assertEqual(User.objects.count(), 1)
        self.assertEqual(UserProfile.objects.count(), 1)
        self.assertEqual(User.objects.get().username, "test_user")

    # Test if user can get token using authentication endpoint
    def test_user_authentication(self):
        User.objects.create_user(username="test_user", password="test_password")
        request_data = {
            "username": "test_user",
            "password": "test_password"
        }
        response = self.client.post("/api/token/", request_data, format="json")
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        self.assertIn("access", response.data)
        self.assertIn("refresh", response.data)

class TasksAPITestCase(APITestCase):
    # Create test user and authenticate the client
    def setUp(self):
        self.user = User.objects.create_user(username="test_user", password="test_password")
        self.client.force_authenticate(user=self.user)

    # Test if task can be created properly
    def test_task_create(self):
        request_data = {
            "name": "Test task name"
        }
        response = self.client.post("/api/tasks/", request_data, format="json")
        self.assertEqual(response.status_code, status.HTTP_201_CREATED)
        self.assertEqual(response.data["task"]["name"], "Test task name")
        self.assertEqual(Task.objects.filter(user=self.user).count(), 1)

    # Test if creating tasks beyond task limit will return an error
    def test_task_create_limit(self):
        request_data = {
            "name": "Test task name"
        }
        
        # Maximize task limit
        task_limit = UserProfile.objects.get(user=self.user).task_limit
        for i in range(task_limit):
            self.client.post("/api/tasks/", request_data, format="json")

        # Next task creation should fail
        response = self.client.post("/api/tasks/", request_data, format="json")
        self.assertEqual(response.status_code, status.HTTP_403_FORBIDDEN)
        self.assertEqual(Task.objects.filter(user=self.user).count(), task_limit)

    # Test task list retrieval for user
    def test_task_read(self):
        task1 = Task.objects.create(user=self.user, name="Test task 1 name")
        task2 = Task.objects.create(user=self.user, name="Test task 2 name")
        response = self.client.get("/api/tasks/", format="json")
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        self.assertEqual(len(response.data["tasks"]), 2)
        self.assertIn({"id": task1.id, "name": "Test task 1 name"}, response.data["tasks"])
        self.assertIn({"id": task2.id, "name": "Test task 2 name"}, response.data["tasks"])

    # Test if task updates properly
    def test_task_update(self):
        task = Task.objects.create(user=self.user, name="Test task before update")
        request_data = {
            "name": "Test task after update"
        }
        
        # Update of task must be successful
        response = self.client.put("/api/tasks/{}/".format(task.id), request_data, format="json")
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        
        # Retrieving task list should return updated name
        response = self.client.get("/api/tasks/", format="json")
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        self.assertIn({"id": task.id, "name": "Test task after update"}, response.data["tasks"])

    # Test if task deletes properly
    def test_task_delete(self):
        task = Task.objects.create(user=self.user, name="Test task name")
        response = self.client.delete("/api/tasks/{}/".format(task.id), format="json")
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        self.assertEqual(Task.objects.count(), 0)
