import datetime

from django.contrib.auth.models import User
from django.utils import timezone
from rest_framework.test import APITestCase

from todo.models import TodoTask
from todo.models import UserProfile

# Create your tests here.
class TodoTaskTest(APITestCase):
    def setUp(self):
        """
        Set up initial data to be used for tests
        """
        self.user = User.objects.create_user(
            username='test_user', 
            email='test_user@test.com', 
            password='TestPW123!')
        self.user.save()

        self.profile = UserProfile.objects.create(
            user=self.user, limit=3)
        self.profile.save()

        self.url = 'http://127.0.0.1:8000/api/todo'

    def test_create_task(self):
        """
        Test creation of task
        """
        task = {
            'title': 'Test creation of task',
            'text': 'Here we`re testing if creating a task works',
            'user': self.profile.id
        }

        response = self.client.post(self.url, task)
        self.assertEqual(201, response.status_code)

    def test_create_task_no_user(self):
        """
        Test creation of a task with invalid user/user profile
        """
        task = {
            'title': 'Test creation of task without a user',
            'text': 'Now let`s test what happens if the user doesn`t exist',
            'user': 0
        }

        response = self.client.post(self.url, task)
        self.assertEqual(400, response.status_code)

    def test_daily_limit_over(self):
        """
        Test creation of a number of tasks
        exceeding the daily limit of user
        """
        for i in range(self.profile.limit):
            task = {
                'title': f'Task #{i}',
                'text': 'Testing daily limit',
                'user': self.profile.id
            }
            self.client.post(self.url, task)

        task['title'] = 'Final task'
        task['text'] = 'This should yield a 429'

        response = self.client.post(self.url, task)
        self.assertEqual(429, response.status_code)

    def test_daily_limit_under(self):
        """
        Test creation of a number of tasks 
        less than the daily limit of user
        """
        for i in range(self.profile.limit-1):
            task = {
                'title': f'Task #{i}',
                'text': 'Testing daily limit',
                'user': self.profile.id
            }
            self.client.post(self.url, task)

        task['title'] = 'Final task'
        task['text'] = 'This should NOT yield a 429'

        response = self.client.post(self.url, task)
        self.assertEqual(201, response.status_code)

    def test_task_created_today(self):
        """
        Test if the daily limit does not 
        consider tasks created earlier than today
        """

        # Create tasks in database
        for i in range(self.profile.limit):
            data = {
                'title': f'Task #{i}',
                'text': 'Testing daily limit',
                'user': self.profile
            }
            
            # Manually change date_created
            yesterday = timezone.now() - datetime.timedelta(days=1)
            task = TodoTask(**data)
            task.save()
            task.date_created=yesterday
            task.save()

        # Post a task
        task = {
                'title': 'Final Task',
                'text': 'This should NOT yield a 429',
                'user': self.profile.id
        }

        response = self.client.post(self.url, task)
        self.assertEqual(201, response.status_code)