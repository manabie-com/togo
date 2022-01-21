from .models.base import Base, engine
from .models.user import User
from .models.task import Task
from .models.assignment import Assignment
from .models.user_date import DateOfUser
from .models.calender import Calender


def init_databases():
    Base.metadata.create_all(engine)
