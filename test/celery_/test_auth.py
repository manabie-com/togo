import jwt
import os
import sys

from database.models.user import User
from celery_.response import ServiceResponse
from flask_.error import InvalidUsage
from test.celery_.base_test_case import BaseTestCase

sys.path.append(os.getcwd())


class AuthTest(BaseTestCase):
    def setUp(self):
        super().add_mock_services_to_db()
        from celery_.services import auth_service

        self.auth_service = auth_service

        self.user = User(username="user1", password="test_password")
        self.user.hash_password()
        self.user.save()

    def tearDown(self):
        super().drop_collection()
        User.drop_collection()

    def test_post_signup(self):
        signup_data = {"username": "user2", "password": "test_password"}

        result = self.auth_service.post_signup.s(signup_data).apply().get()
        username = User.objects.get(id=result.get("user_id")).username

        self.assertIs(type(result), ServiceResponse)
        self.assertEqual(username, signup_data.get("username"))

    def test_post_signin(self):
        signin_data = {"username": "user1", "password": "test_password"}

        result = self.auth_service.post_signin.s(signin_data).apply().get()
        jwt_token = jwt.encode(
            payload={"user_id": str(self.user.id)}, key="secret", algorithm="HS256"
        )
        self.assertIs(type(result), ServiceResponse)
        self.assertEqual(result.get("bearer_token"), f"Bearer: {jwt_token}")

    def test_post_signin_fail(self):
        signin_data = {"username": "user1", "password": "wrong_password"}

        task = self.auth_service.post_signin.s(signin_data).apply()

        self.assertIs(type(task.result), InvalidUsage)
        self.assertEqual(task.result.payload, "Login fail")
        self.assertEqual(task.result.message, "Your password is not correct")
