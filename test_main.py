from fastapi.testclient import TestClient
from datetime import date, timedelta
from main import app
from app import models, crud
from app.database import SessionLocal, engine

db = SessionLocal()
models.Base.metadata.create_all(bind=engine)


client = TestClient(app)


def test_class():
    user = models.User(limit=1)
    assert user.limit == 1
    task = models.Task(user_id=1, name='todo')
    assert task.user_id == 1
    assert task.name == 'todo'
    assert task.create_date is None


def test_database_create_data():
    db.query(models.User).delete()
    db.query(models.Task).delete()
    db.add(models.User(limit=1))
    db.add(models.User(limit=3))
    db.commit()
    yesterday = date.today() - timedelta(days=1)
    db.add(models.Task(name='todo', user_id=1, create_date=yesterday))
    db.add(models.Task(name='todo1', user_id=2, create_date=yesterday))
    db.add(models.Task(name='todo2', user_id=2, create_date=yesterday))
    db.add(models.Task(name='todo3', user_id=2, create_date=date.today()))
    db.commit()

    users = db.query(models.User).all()
    assert users[0].id == 1
    assert users[0].limit == 1
    assert users[0].tasks[0].name == 'todo'
    assert users[0].tasks[0].create_date == yesterday
    assert users[1].id == 2
    assert users[1].limit == 3


def test_crud_get_user():
    user = crud.get_user(db, 2)
    assert user.id == 2
    assert user.limit == 3


def test_crud_create_task():
    task = crud.create_task(db, models.Task(user_id=2, name='todo4'))
    assert task.create_date == date.today()


def test_crud_count_task():
    count = crud.count_task(db, 1)
    assert count == 0
    count = crud.count_task(db, 2)
    assert count == 2


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


def test_create_task_user_not_exist():
    data = {
        'name': 'todo',
        'user_id': 0,
    }
    response = client.post('/tasks/', json=data)
    assert response.status_code == 400
    assert response.json() == {'detail': 'user not exist'}


def test_create_task_over_limit():
    data = {
        'name': 'todo2',
        'user_id': 1,
    }
    response = client.post('/tasks/', json=data)
    assert response.status_code == 429
    assert response.json() == {'detail': 'number of tasks per day reached'}
