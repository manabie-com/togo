import abc
from app.task.domain import model
from typing import List

class AbstractRepository(abc.ABC):

    @abc.abstractmethod
    def add(self, todo: model.Base):
        raise NotImplementedError

    @abc.abstractmethod
    def get(self, title) -> model.Base:
        raise NotImplementedError



class TodoListRepository(AbstractRepository):

    def __init__(self, session):
        self.session = session

    def add(self, todo: model.TodoList):
        self.session.add(todo)

    def get(self, title):
        return self.session.query(model.TodoList).filter_by(title=title).first()

    def add_parent(self, todo_list: model.TodoList, parent: model.TodoList):
        todo_list.parent = parent
    
    def add_todo_item(self, todo_list: model.TodoList, todo_item: model.TodoItem):
        todo_list.todo_items.append(todo_item)
    
    def add_todo_items(self, todo_list: model.TodoList, todo_items: List[model.TodoItem]):
        for item in todo_items:
            todo_list.todo_items.append(item)

class TodoItemRepository(AbstractRepository):

    def __init__(self, session):
        self.session = session

    def add(self, todo: model.TodoItem):
        self.session.add(todo)

    def get(self, title):
        return self.session.query(model.TodoItem).filter_by(title=title).first()

    def add_parent(self, todo_item: model.TodoItem, parent: model.TodoList):
        todo_item.parent = parent