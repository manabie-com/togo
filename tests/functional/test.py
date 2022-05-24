import requests
import app
import datetime


def test_registration():
    sample_user = {"name": "test", "password": "password", "limit_per_day": 1}
    response = requests.post("http://127.0.0.1:5000/register", json=sample_user)

    assert response.status_code == 200
    assert response.json()["message"] == "registration success."

    # register non-unique user
    response = requests.post("http://127.0.0.1:5000/register", json=sample_user)
    assert response.status_code == 500


def test_login():
    response = requests.post("http://127.0.0.1:5000/login", auth=("test", "password"))

    assert response.status_code == 200
    assert "token" in response.json()

    # login non-existing user

    response = requests.post("http://127.0.0.1:5000/login", auth=("error", "error"))
    assert response.status_code == 401


def test_create_todo():
    response = requests.post("http://127.0.0.1:5000/login", auth=("test", "password"))
    token = response.json()["token"]

    sample_test = {"todo": "sample todo task"}
    response = requests.post("http://127.0.0.1:5000/todo", json=sample_test, headers={"x-access-token": token})

    assert response.status_code == 200
    assert response.json()["message"] == "new task added."

    # exceed allocated task per day
    sample_test = {"todo": "sample todo task"}
    response = requests.post("http://127.0.0.1:5000/todo", json=sample_test, headers={"x-access-token": token})

    assert response.status_code == 400
    assert response.text == "User has reached maximum todos per day."


def test_get_todo():
    response = requests.post("http://127.0.0.1:5000/login", auth=("test", "password"))
    token = response.json()["token"]

    response = requests.get("http://127.0.0.1:5000/todos", headers={"x-access-token": token})

    todos = response.json()["todos"]

    # check if the data of the first record
    assert todos[0]["task"] == "sample todo task"


def test_todo_delete():
    response = requests.post("http://127.0.0.1:5000/login", auth=("test", "password"))
    token = response.json()["token"]

    response = requests.get("http://127.0.0.1:5000/todos", headers={"x-access-token": token})

    todos = response.json()["todos"]

    todo_id = todos[0]["id"]

    response = requests.post("http://127.0.0.1:5000/todo/delete", json={"todo_id": todo_id},
                             headers={"x-access-token": token})

    assert response.status_code == 200
    assert response.json()["message"] == "todo deleted."


def test_user_delete():
    response = requests.post("http://127.0.0.1:5000/login", auth=("test", "password"))
    token = response.json()["token"]

    response = requests.post("http://127.0.0.1:5000/user/delete", headers={"x-access-token": token})

    assert response.status_code == 200
    assert response.json()["message"] == "user deleted."
