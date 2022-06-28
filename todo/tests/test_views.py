from django.test import TestCase, Client
from django.utils import timezone

from utils import encrypting


class UserDataTest:
    USER_USERNAME_1 = "user_test_001"
    USER_USERNAME_2 = "user_test_002"
    USER_PASSWORD = "Aa123456"
    ADMIN_USERNAME = "admin"
    ADMIN_PASSWORD = "admin"


class TestViews(TestCase):
    def setUp(self):
        self.client = Client()
        self.registration_url = "/api/users/registration/"
        self.login_url = "/api/login/"
        self.list_url = "/api/tasks/"
        self.detail_url = "/api/tasks/{0}/"
        self.user_detail_url = "/api/users/{0}/"
        self.task_title = "test get task detail {0}"
        self.task_description = "get task detail description"
        self.tokens = self.__register_and_login(
            UserDataTest.USER_USERNAME_1, UserDataTest.USER_PASSWORD
        )
        self.admin_tokens = self.__register_and_login(
            UserDataTest.ADMIN_USERNAME, UserDataTest.ADMIN_PASSWORD, True
        )

    def tearDown(self):
        del self.client
        del self.registration_url
        del self.login_url
        del self.list_url
        del self.detail_url
        del self.task_title
        del self.task_description
        del self.tokens
        del self.admin_tokens

    def test_GET_task_list_without_authentication(self):
        print("Running test_GET_task_list_without_authentication")
        response = self.client.get(self.list_url)
        self.assertEquals(response.status_code, 401)

    def test_GET_task_detail_without_authentication(self):
        print("Running test_GET_task_detail_without_authentication")
        response = self.client.get(self.detail_url)
        self.assertEquals(response.status_code, 401)

    def test_POST_add_new_task_without_authentication(self):
        print("Running test_POST_add_new_task_without_authentication")
        response = self.client.post(
            self.list_url,
            {
                "title": self.task_title.format(timezone.now()),
                "description": self.task_description,
            },
        )
        self.assertEquals(response.status_code, 401)

    def test_POST_add_new_task_with_authentication(self):
        print("Running test_POST_add_new_task_with_authentication")
        user_auth_header = self.__get_auth_header(self.tokens["access"])

        list_of_tasks_before_add_task = self.__get_list_of_tasks(user_auth_header)
        response = self.__add_task(
            self.task_title.format(timezone.now()),
            self.task_description,
            user_auth_header,
        )
        list_of_tasks_after_add_task = self.__get_list_of_tasks(user_auth_header)

        self.assertEquals(response.status_code, 201)
        self.assertGreater(
            len(list_of_tasks_after_add_task.data),
            len(list_of_tasks_before_add_task.data),
        )

    def test_GET_task_list_with_authentication(self):
        print("Running test_GET_task_list_with_authentication")
        user_auth_header = self.__get_auth_header(self.tokens["access"])
        response = self.__get_list_of_tasks(user_auth_header)
        self.assertEquals(response.status_code, 200)

    def test_GET_task_detail_with_authentication(self):
        print("Running test_GET_task_detail_with_authentication")
        user_auth_header = self.__get_auth_header(self.tokens["access"])

        for i in range(2):
            self.__add_task(
                self.task_title.format(timezone.now()),
                self.task_description,
                user_auth_header,
            )

        list_of_tasks = self.__get_list_of_tasks(user_auth_header)
        task = list_of_tasks.data[0]

        response = self.__get_detail_of_task(task["id"], user_auth_header)

        self.assertEquals(response.status_code, 200)
        self.assertEquals(
            encrypting.decrypt(task["id"])[0],
            encrypting.decrypt(response.data["id"])[0],
        )

    def test_DELETE_task_with_authentication(self):
        print("Running test_DELETE_task_with_authentication")
        user_auth_header = self.__get_auth_header(self.tokens["access"])

        for i in range(2):
            self.__add_task(
                self.task_title.format(i + 1), self.task_description, user_auth_header
            )

        list_of_tasks_before_delete_task = self.__get_list_of_tasks(user_auth_header)
        task = list_of_tasks_before_delete_task.data[0]
        self.__delete_task(task["id"], user_auth_header)
        list_of_tasks_after_delete_task = self.__get_list_of_tasks(user_auth_header)

        self.assertGreater(
            len(list_of_tasks_before_delete_task.data),
            len(list_of_tasks_after_delete_task.data),
        )

    def test_POST_add_task_exceed_maximum_tasks_per_day_scenario_1(self):
        print("Running test_POST_add_task_exceed_maximum_tasks_per_day_scenario_1")
        tasks = list()

        user = self.__register_new_user(
            UserDataTest.USER_USERNAME_2, UserDataTest.USER_PASSWORD
        )
        user_login_response = self.__login(
            UserDataTest.USER_USERNAME_2, UserDataTest.USER_PASSWORD
        )

        user_auth_header = self.__get_auth_header(user_login_response.data["access"])
        maximum_task_per_day = user.data["maximum_task_per_day"]

        for i in range(maximum_task_per_day + 10):
            tasks.append(
                self.__add_task(
                    self.task_title.format(i + 1),
                    self.task_description,
                    user_auth_header,
                )
            )

        response = self.__add_task(
            self.task_title.format(timezone.now()),
            self.task_description,
            user_auth_header,
        )

        self.assertEquals(response.status_code, 400)
        self.assertNotEquals(maximum_task_per_day, len(tasks))

    def test_POST_add_task_exceed_maximum_tasks_per_day_scenario_2(self):
        print("Running test_POST_add_task_exceed_maximum_tasks_per_day_scenario_2")
        tasks = list()

        user = self.__register_new_user(
            UserDataTest.USER_USERNAME_2, UserDataTest.USER_PASSWORD
        )
        user_login_response = self.__login(
            UserDataTest.USER_USERNAME_2, UserDataTest.USER_PASSWORD
        )

        user_auth_header = self.__get_auth_header(user_login_response.data["access"])
        maximum_task_per_day = user.data["maximum_task_per_day"]

        for i in range(maximum_task_per_day + 10):
            tasks.append(
                self.__add_task(
                    self.task_title.format(i + 1),
                    self.task_description,
                    user_auth_header,
                )
            )
        response = self.__add_task(
            self.task_title.format(timezone.now()),
            self.task_description,
            user_auth_header,
        )
        self.assertEquals(response.status_code, 400)

        response = self.__delete_task(tasks[0].data["id"], user_auth_header)
        self.assertEquals(response.status_code, 204)

        response = self.__add_task(
            self.task_title.format(timezone.now()),
            self.task_description,
            user_auth_header,
        )

        admin_auth_header = self.__get_auth_header(self.admin_tokens["access"])

        maximum_task_per_day -= 1
        response = self.__update_user_maximum_tasks_per_day(
            user.data["id"], maximum_task_per_day, admin_auth_header
        )
        self.assertEqual(response.status_code, 400)

        maximum_task_per_day += 5
        response = self.__update_user_maximum_tasks_per_day(
            user.data["id"], maximum_task_per_day, admin_auth_header
        )
        self.assertEquals(response.status_code, 200)

        response = self.__add_task(
            self.task_title.format(timezone.now()),
            self.task_description,
            user_auth_header,
        )
        self.assertEquals(response.status_code, 201)

    def test_POST_user_cannot_get_list_tasks_of_other(self):
        print("Running test_POST_user_cannot_get_list_tasks_of_other")
        user_1 = self.__register_and_login(
            UserDataTest.USER_USERNAME_1, UserDataTest.USER_PASSWORD
        )
        user_2 = self.__register_and_login(
            UserDataTest.USER_USERNAME_2, UserDataTest.USER_PASSWORD
        )

        user_1_auth_header = self.__get_auth_header(user_1["access"])
        user_2_auth_header = self.__get_auth_header(user_2["access"])
        user_1_number_of_tasks, user_2_number_of_tasks = 3, 5

        for i in range(user_1_number_of_tasks):
            self.__add_task(
                self.task_title.format(timezone.now()),
                self.task_description,
                user_1_auth_header,
            )
        for i in range(user_2_number_of_tasks):
            self.__add_task(
                self.task_title.format(timezone.now()),
                self.task_description,
                user_2_auth_header,
            )

        user_1_list_tasks_response = self.__get_list_of_tasks(user_1_auth_header)
        user_2_list_tasks_response = self.__get_list_of_tasks(user_2_auth_header)

        self.assertEquals(len(user_1_list_tasks_response.data), user_1_number_of_tasks)
        self.assertEquals(len(user_2_list_tasks_response.data), user_2_number_of_tasks)
        self.assertNotEqual(
            len(user_1_list_tasks_response.data), len(user_2_list_tasks_response.data)
        )

    def test_GET_user_cannot_get_detail_task_of_other(self):
        print("Running test_GET_user_cannot_get_detail_task_of_other")
        new_user = self.__register_and_login(
            UserDataTest.USER_USERNAME_2, UserDataTest.USER_PASSWORD
        )
        user_1_auth_header = self.__get_auth_header(self.tokens["access"])
        new_user_auth_header = self.__get_auth_header(new_user["access"])

        task = self.__add_task(
            self.task_title.format(timezone.now()),
            self.task_description,
            new_user_auth_header,
        )

        response = self.__get_detail_of_task(task.data["id"], user_1_auth_header)
        self.assertEquals(response.status_code, 404)

        response = self.__get_detail_of_task(task.data["id"], new_user_auth_header)
        self.assertEquals(response.status_code, 200)

    def __register_new_user(self, username, password, is_superuser=False):
        data = {
            "username": username,
            "password": password,
            "is_superuser": is_superuser,
        }
        response = self.client.post(self.registration_url, data)
        return response

    def __login(self, username, password):
        data = {"username": username, "password": password}
        response = self.client.post(self.login_url, data)
        return response

    def __register_and_login(self, username, password, is_superuser=False):
        self.__register_new_user(username, password, is_superuser)
        response = self.__login(username, password)
        return response.data

    def __get_auth_header(self, access_token):
        return {"HTTP_AUTHORIZATION": "Bearer " + access_token}

    def __update_user_maximum_tasks_per_day(self, user_id, maximum_task, auth_header):
        return self.client.put(
            self.user_detail_url.format(user_id),
            {"maximum_task_per_day": maximum_task},
            content_type="application/json",
            **auth_header
        )

    def __delete_task(self, task_id, auth_header):
        return self.client.delete(self.detail_url.format(task_id), **auth_header)

    def __get_list_of_tasks(self, auth_header):
        return self.client.get(self.list_url, **auth_header)

    def __get_detail_of_task(self, task_id, auth_header):
        return self.client.get(self.detail_url.format(task_id), **auth_header)

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
