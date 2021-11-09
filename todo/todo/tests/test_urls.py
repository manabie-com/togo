from django.test import SimpleTestCase
from django.urls import reverse, resolve
from todo.views import login, tasks

class TestUrls(SimpleTestCase):

    def test_login_url_is_resolves(self):
        url = reverse('login')
        self.assertEquals(resolve(url).func, login)

    def test_tasks_url_is_resolves(self):
        url = reverse('tasks')
        self.assertEquals(resolve(url).func, tasks)
