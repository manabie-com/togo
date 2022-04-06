import pytest
from app.task.service_layer.unit_of_work import TodoListUnitOfWork
from app import create_app, db

app = create_app("testing")

def test_roll_back_on_exception():
    class MockException(Exception):
        pass
    with app.app_context():
        uow = TodoListUnitOfWork()
        with pytest.raises(MockException):
            with uow:
                db.session.execute(
                "INSERT INTO todo_list (title, description) VALUES (:title, :description)",
                    dict(title="test title", description="test desc"),
                )
                raise MockException()
        data = list(db.session.execute("SELECT * FROM todo_list WHERE title='test title'"))
        assert data == []

def test_roll_back_on_uncommitted():
    with app.app_context():
        uow = TodoListUnitOfWork()
        with uow:
            db.session.execute(
            "INSERT INTO todo_list (title, description) VALUES (:title, :description)",
                dict(title="test title", description="test desc"),
            )
        data = list(db.session.execute("SELECT * FROM todo_list WHERE title='test title'"))
        assert data == []