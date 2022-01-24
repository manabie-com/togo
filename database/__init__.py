from .models.assignment import Assignment
from .models.base import Base, engine
from .models.calender import Calender
from .models.task import Task
from .models.user import User


def init_databases():
    Base.metadata.create_all(engine)
