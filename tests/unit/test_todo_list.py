from unittest import mock
from app.task.adapters.repository import AbstractRepository
from app.task.service_layer.services import AddTodoListService
from app.task.service_layer.unit_of_work import AbstractUnitOfWork
import pytest
from app.exceptions import DuplicateTitleError, ExceedLimitationError

class FakeTodoListRepository(AbstractRepository):
    todo_lists = []

    def __init__(self, todo_lists):
        self.todo_lists = todo_lists
    
    def add(self, todo_list):
        self.todo_lists.append(todo_list)
    
    def get(self, title):
        return next((todo for todo in self.todo_lists if todo.title == title), None)

class FakeTodoListUnitOfWork(AbstractUnitOfWork):
    def __init__(self, default_todo_lists=set()) -> None:
        self.session = default_todo_lists
        self.repository = FakeTodoListRepository(self.session)
        self.committed = False
    
    def commit(self):
        self.committed = True
    
    def add(self, todo_list):
        self.session.add(todo_list)

    def rollback(self):
        pass

def test_add_todo_list():
    uow = FakeTodoListUnitOfWork()
    AddTodoListService.add("test title", "", [], 0, uow)
    assert uow.repository.get("test title") is not None
    assert uow.committed

def test_add_duplicate_todo_list():
    uow = FakeTodoListUnitOfWork([mock.Mock(title="test title", description="")])
    with pytest.raises(DuplicateTitleError):
        AddTodoListService.add("test title", "", [], 0, uow)
        assert uow.repository.get("test title") is not None
        assert uow.committed is False

def test_todos_greater_than_limitation():
        uow = FakeTodoListUnitOfWork()

        with pytest.raises(ExceedLimitationError):
            AddTodoListService.add("test title 3", "", [
                {"title": "test title 1", "description": ""}, 
                {"title": "test title 2", "description": ""}
            ], 1, uow)
            assert uow.repository.get("test title 3") is not None
            assert uow.committed is False
