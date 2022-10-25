import os
import tempfile

import pytest
import os
import json
from src import create_app
from .database import init_db, PRICING, USER
from src.util import logger, load_config
from src.api.services import create_token


@pytest.fixture
def client():
    app = create_app("test")
    os.environ["config_file"] = "config/config-test.yml"
    init_db()
    with app.test_client() as client:
        yield client


def test_pricing(client):
    rsp = client.get("/api/subscription", follow_redirects=True)
    data = json.loads(rsp.data)
    pricing_option_ids = sorted(list(map(lambda x: x.get("id"), PRICING)))
    dest_id = sorted(list(map(lambda x: x.get("id"), data.get("data"))))
    assert ".".join(pricing_option_ids) == ".".join(dest_id)


def test_sign_up(client):
    """Start with a blank database."""
    TEST_DATA = [
        (USER[0], 500),
        (USER[1], 500),
        (USER[2], 500),
        ({
             "email": "test-mail-4@gmail.com",
             "password": "easy-peasy",
             "fullname": "Fifth User"
         }, 201)
    ]
    for info, status in TEST_DATA:
        logger.info(info)
        res = client.post('/api/auth/sign-up', json=info, follow_redirects=True)
        assert res.status_code == status


def test_sign_in(client):
    """Start with a blank database."""
    TEST_DATA = [
        (USER[0], create_token({"userId": USER[0].get("id")})),
        (USER[1], create_token({"userId": USER[1].get("id")})),
        (USER[2], create_token({"userId": USER[2].get("id")})),
    ]
    for info, token in TEST_DATA:
        logger.info(info)
        res = client.post('/api/auth/sign-in', json={
            "email": info.get("email"),
            "password": info.get("password")
        }, follow_redirects=True)
        data = json.loads(res.data)
        assert data.get("token") == token


def test_create_task(client):
    task_data = {
        "summary": "hehe",
        "description": "hihi"
    }
    TEST_DATA = [
        (create_token({"userId": USER[0].get("id")}), 5, 201),
        (create_token({"userId": USER[1].get("id")}), 5, 201),
        (create_token({"userId": USER[2].get("id")}), 5, 201),
    ]
    for token, amount, status in TEST_DATA:
        for i in range(amount + 1):
            res = client.post("/api/task", json=task_data, follow_redirects=True,headers={'Authorization':f'Bearer {token}'})
            if i == amount:
                assert res.status_code != 201
            else:
                assert res.status_code == 201
