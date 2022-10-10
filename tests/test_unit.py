import pytest
from flask_login import AnonymousUserMixin
from flask_login import login_user, login_manager
from app import create_app, db
from app import init_routes
from app import Users
from werkzeug.security import generate_password_hash

import jwt
import datetime


# util for decoding the user token
def decode_token(username):
    token = jwt.encode({"public_id": username,
                        "exp": datetime.datetime.utcnow() + datetime.timedelta(minutes=30)},
                       "secret_key")
    return token


# test the registration endpoint
def test_registration(client):
    sample_user = {"name": "test", "password": "password", "limit_per_day": 1}
    result = client.post("/register", json=sample_user)

    assert result.status_code == 200
    assert result.get_json()["message"] == "registration success."


# test the login endpoint
def test_login(client):

    result = client.post("/login", auth=("sample_user", "sample_password"))

    assert result.status_code == 200
    assert "token" in result.get_json()

## TODO unable to do fixture testing using anonymous user
# test the create_todo endpoint
# def test_create_todo(client):
#
#     payload = {"todo": "sample todo task"}
#     result = client.post("/todo", json=payload, headers={"x-access-token": decode_token("sample_user")})
#
#     assert result.status_code == 200
#     assert result.get_json()["message"] == "new task added."
