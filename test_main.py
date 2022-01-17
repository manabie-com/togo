from fastapi.testclient import TestClient
from main import app


client = TestClient(app)


def test_create_task():
    data = {
        'text': 'todo',
        'user_id': 1,
    }
    response = client.post('/tasks/', json=data)
    assert response.status_code == 200
    res_data = response.json()
    assert res_data['text'] == data['text']
    assert res_data['user_id'] == data['user_id']


def test_create_task_user_not_exist():
    data = {
        'text': 'todo',
        'user_id': 0,
    }
    response = client.post('/tasks/', json=data)
    assert response.status_code == 400
    assert response.json() == {'detail': 'user not exist'}


def test_create_task_over_limit():
    data = {
        'text': 'todo',
        'user_id': 3,
    }
    response = client.post('/tasks/', json=data)
    assert response.status_code == 429
    assert response.json() == {'detail': 'number of tasks per day reached'}
