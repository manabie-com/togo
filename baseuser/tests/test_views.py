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
        self.list_url = "/api/users/"
        self.detail_url = "/api/users/{0}/"

    def tearDown(self):
        del self.client
        del self.registration_url
        del self.login_url
        del self.list_url
        del self.detail_url

    def test_POST_registration(self):
        print("Running test_POST_registration")
        response = self.__login(
            UserDataTest.USER_USERNAME_1, UserDataTest.USER_PASSWORD
        )
        self.assertEquals(response.status_code, 401)

        response = self.__register_new_user(
            UserDataTest.USER_USERNAME_1, UserDataTest.USER_PASSWORD
        )
        self.assertEquals(response.status_code, 201)

    def test_POST_login(self):
        print("Running test_POST_login")
        response = self.__register_new_user(
            UserDataTest.USER_USERNAME_1, UserDataTest.USER_PASSWORD
        )
        response = self.__login(
            UserDataTest.USER_USERNAME_1, UserDataTest.USER_PASSWORD
        )
        self.assertEquals(response.status_code, 200)

    def test_POST_registration_and_login(self):
        print("Running test_POST_registration_and_login")
        response = self.__register_new_user(
            UserDataTest.USER_USERNAME_1, UserDataTest.USER_PASSWORD
        )
        self.assertEquals(response.status_code, 201)

        response = self.__login(
            UserDataTest.USER_USERNAME_1, UserDataTest.USER_PASSWORD
        )
        self.assertEquals(response.status_code, 200)

    def test_GET_list_users(self):
        print("Running test_GET_list_users")
        user = self.__register_and_login(
            UserDataTest.USER_USERNAME_1, UserDataTest.USER_PASSWORD
        )
        admin = self.__register_and_login(
            UserDataTest.ADMIN_USERNAME, UserDataTest.ADMIN_PASSWORD, True
        )

        user_auth_header = self.__get_auth_header(user["access"])
        admin_auth_header = self.__get_auth_header(admin["access"])

        user_get_list_response = self.__get_list_of_users(user_auth_header)
        self.assertEquals(user_get_list_response.status_code, 403)

        admin_get_list_response = self.__get_list_of_users(admin_auth_header)
        self.assertEquals(admin_get_list_response.status_code, 200)

    def test_GET_detail_user_by_current_user(self):
        print("Running test_GET_detail_user_by_current_user")
        user = self.__register_new_user(
            UserDataTest.USER_USERNAME_1, UserDataTest.USER_PASSWORD
        )
        login = self.__login(UserDataTest.USER_USERNAME_1, UserDataTest.USER_PASSWORD)

        user_auth_header = self.__get_auth_header(login.data["access"])
        user_detail = self.__get_detail_of_user(user.data["id"], user_auth_header)
        self.assertEquals(user_detail.status_code, 200)

    def test_GET_detail_user_by_admin_user(self):
        print("Running test_GET_detail_user_by_admin_user")
        user = self.__register_new_user(
            UserDataTest.USER_USERNAME_1, UserDataTest.USER_PASSWORD
        )
        admin_login = self.__register_and_login(
            UserDataTest.ADMIN_USERNAME, UserDataTest.ADMIN_PASSWORD, True
        )

        admin_auth_header = self.__get_auth_header(admin_login["access"])
        response = self.__get_detail_of_user(user.data["id"], admin_auth_header)
        self.assertEquals(response.status_code, 200)

    def test_GET_detail_other_user_by_current_user(self):
        print("Running test_GET_detail_other_user_by_current_user")
        user_1_login = self.__register_and_login(
            UserDataTest.USER_USERNAME_1, UserDataTest.USER_PASSWORD
        )
        user_2 = self.__register_new_user(
            UserDataTest.USER_USERNAME_2, UserDataTest.USER_PASSWORD
        )

        user_1_auth_header = self.__get_auth_header(user_1_login["access"])
        response = self.__get_detail_of_user(user_2.data["id"], user_1_auth_header)
        self.assertEquals(response.status_code, 403)

    def __register_new_user(self, username, password, is_superuser=False):
        response = self.client.post(
            self.registration_url,
            {"username": username, "password": password, "is_superuser": is_superuser},
        )
        return response

    def __login(self, username, password):
        response = self.client.post(
            self.login_url, {"username": username, "password": password}
        )
        return response

    def __register_and_login(self, username, password, is_superuser=False):
        self.__register_new_user(username, password, is_superuser)
        response = self.__login(username, password)
        return response.data

    def __get_auth_header(self, access_token):
        return {"HTTP_AUTHORIZATION": "Bearer " + access_token}

    def __get_list_of_users(self, auth_header):
        return self.client.get(self.list_url, **auth_header)

    def __get_detail_of_user(self, user_id, auth_header):
        return self.client.get(self.detail_url.format(user_id), **auth_header)
