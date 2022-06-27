from django.test import TestCase, Client
from django.utils import timezone

from utils import encrypting


class UserDataTest:
    USER_USERNAME = "user_test_001"
    USER_PASSWORD = "Aa123456"


class TestViews(TestCase):
    def setUp(self):
        self.client = Client()
        self.registration_url = "/api/users/registration/"
        self.login_url = "/api/login/"
        self.list_url = "/api/tasks/"
        self.detail_url = "/api/tasks/{0}/"
        self.tokens = self.__register_and_login(
            UserDataTest.USER_USERNAME, UserDataTest.USER_PASSWORD
        )

    def test_GET_task_list_without_authentication(self):
        response = self.client.get(self.list_url)
        self.assertEquals(response.status_code, 401)

    def test_GET_task_detail_without_authentication(self):
        response = self.client.get(self.detail_url)
        self.assertEquals(response.status_code, 401)

    def test_POST_add_new_task_without_authentication(self):
        response = self.client.post(
            self.list_url,
            {"title": "Add a new task", "description": "Test add a new task"},
        )
        self.assertEquals(response.status_code, 401)

    def test_POST_add_new_task_with_authentication(self):
        auth_header = self.__get_auth_header(self.tokens["access"])
        task_title = "test get task detail {0}".format(timezone.now())
        task_description = "get task detail description"

        list_of_tasks_before_add_task = self.__get_list_of_tasks(auth_header)
        response = self.__add_task(task_title, task_description, auth_header)
        list_of_tasks_after_add_task = self.__get_list_of_tasks(auth_header)

        self.assertEquals(response.status_code, 201)
        self.assertGreater(
            len(list_of_tasks_after_add_task.data),
            len(list_of_tasks_before_add_task.data),
        )

    def test_GET_task_list_with_authentication(self):
        auth_header = self.__get_auth_header(self.tokens["access"])
        response = self.__get_list_of_tasks(auth_header)
        self.assertEquals(response.status_code, 200)

    def test_GET_task_detail_with_authentication(self):
        auth_header = self.__get_auth_header(self.tokens["access"])
        task_title = "test get task detail {0}".format(timezone.now())
        task_description = "get task detail description"

        for i in range(2):
            self.__add_task(task_title, task_description, auth_header)

        list_of_tasks = self.__get_list_of_tasks(auth_header)
        task = list_of_tasks.data[0]

        response = self.client.get(self.detail_url.format(task["id"]), **auth_header)

        self.assertEquals(response.status_code, 200)
        self.assertEquals(
            encrypting.decrypt(task["id"])[0],
            encrypting.decrypt(response.data["id"])[0],
        )

    def test_DELETE_task_with_authentication(self):
        auth_header = self.__get_auth_header(self.tokens["access"])
        task_title = "test get task detail %s"
        task_description = "get task detail description"

        for i in range(2):
            self.__add_task(task_title % str(i + 1), task_description, auth_header)

        list_of_tasks_before_delete_task = self.__get_list_of_tasks(auth_header)
        task = list_of_tasks_before_delete_task.data[0]
        self.__delete_task(task["id"], auth_header)
        list_of_tasks_after_delete_task = self.__get_list_of_tasks(auth_header)

        self.assertGreater(
            len(list_of_tasks_before_delete_task.data),
            len(list_of_tasks_after_delete_task.data),
        )

    def test_POST_add_task_exceed_maximum_tasks_per_day(self):
        pass

    def __register_new_user(self, username, password):
        response = self.client.post(
            self.registration_url, {"username": username, "password": password}
        )
        return response

    def __login(self, username, password):
        response = self.client.post(
            self.login_url, {"username": username, "password": password}
        )
        return response

    def __register_and_login(self, username, password):
        self.__register_new_user(username, password)
        response = self.__login(username, password)
        return response.data

    def __get_auth_header(self, access_token):
        return {"HTTP_AUTHORIZATION": "Bearer " + access_token}

    def __delete_task(self, task_id, auth_header):
        return self.client.delete(self.detail_url.format(task_id), **auth_header)

    def __get_list_of_tasks(self, auth_header):
        return self.client.get(self.list_url, **auth_header)

    def __add_task(self, task_title, task_description, auth_header):
        response = self.client.post(
            self.list_url,
            {
                "title": task_title,
                "description": task_description,
            },
            **auth_header
        )
        return response
