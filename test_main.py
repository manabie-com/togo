from fastapi.testclient import TestClient
from datetime import date, timedelta
import pytest
from main import app
import crud
import models
from database import SessionLocal, engine

db = SessionLocal()
models.Base.metadata.create_all(bind=engine)

# Create data for test
db.query(models.User).delete()
db.query(models.Task).delete()
db.add(models.User(limit=1))
db.add(models.User(limit=10))
db.commit()
yesterday = date.today() - timedelta(days=1)
db.add(models.Task(name='todo', user_id=1, create_date=yesterday))
db.add(models.Task(name='todo1', user_id=2, create_date=yesterday))
db.add(models.Task(name='todo2', user_id=2, create_date=yesterday))
db.add(models.Task(name='todo3', user_id=2, create_date=date.today()))
db.commit()


client = TestClient(app)


@pytest.mark.unit
def test_class():
    user = models.User(limit=1)
    assert user.limit == 1
    task = models.Task(user_id=1, name='todo')
    assert task.user_id == 1
    assert task.name == 'todo'
    assert task.create_date is None


@pytest.mark.unit
def test_crud_get_user():
    user = crud.get_user(db, 2)
    assert user.id == 2
    assert user.limit == 10


@pytest.mark.unit
def test_crud_create_task():
    task = crud.create_task(db, models.Task(user_id=2, name='todo4'))
    assert task.create_date == date.today()


@pytest.mark.unit
def test_crud_count_task():
    count = crud.count_task(db, 1)
    assert count == 0
    count = crud.count_task(db, 2)
    assert count == 2


@pytest.mark.integration
def test_validate():
    data = {
        'name': 'todo1',
        'user_id': 'a',
    }
    response = client.post('/tasks/', json=data)
    assert response.status_code == 422
    data = {
        'name': 'todo1',
    }
    response = client.post('/tasks/', json=data)
    assert response.status_code == 422


@pytest.mark.integration
def test_create_task_success():
    data = {
        'name': 'todo1',
        'user_id': 1,
    }
    response = client.post('/tasks/', json=data)
    assert response.status_code == 200
    res_data = response.json()
    assert res_data['name'] == data['name']
    assert res_data['user_id'] == data['user_id']
    assert res_data['create_date'] == str(date.today())


@pytest.mark.integration
def test_create_task_user_not_exist():
    data = {
        'name': 'todo',
        'user_id': 0,
    }
    response = client.post('/tasks/', json=data)
    assert response.status_code == 400
    assert response.json() == {'detail': 'user not exist'}


@pytest.mark.integration
def test_create_task_over_limit():
    data = {
        'name': 'todo2',
        'user_id': 1,
    }
    response = client.post('/tasks/', json=data)
    assert response.status_code == 429
    assert response.json() == {'detail': 'number of tasks per day reached'}
