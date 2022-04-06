
# pylint: disable=attribute-defined-outside-init
import abc
from app import db

from app.task.adapters.repository import TodoListRepository

class AbstractUnitOfWork(abc.ABC):
    def __enter__(self):
        return self

    def __exit__(self, *args):
        self.rollback()

    @abc.abstractmethod
    def commit(self):
        raise NotImplementedError

    @abc.abstractmethod
    def rollback(self):
        raise NotImplementedError


class TodoListUnitOfWork(AbstractUnitOfWork):
    def __init__(self, session=db.session):
        self.session = session

    def __enter__(self):
        self.repository = TodoListRepository(self.session)
        return super().__enter__()

    def __exit__(self, *args):
        super().__exit__(*args)

    def commit(self):
        self.session.commit()
    
    def add(self, todo_list):
        self.session.add(todo_list)

    def rollback(self):
        self.session.rollback()
