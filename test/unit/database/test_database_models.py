from datetime import datetime

from database import User, Assignment, Task, Calender


def test_user():
    user = User(
        id=1,
        name="test",
        ability=3,
    )
    assert isinstance(user, User)
    assert str(user), "test"


def test_assignment():
    assignment = Assignment(
        id=1,
        user_id=1,
        date=datetime.strptime('24/01/2022', "%d/%m/%Y")
    )
    assert isinstance(assignment, Assignment)
    assert str(assignment), '1 - 1 - 2022-01-24 00:00:00'


def test_task():
    task = Task(
        id=1,
        title='test',
        description='test',
        assignment_id=1
    )
    assert isinstance(task, Task)
    assert str(task), '1 - test'


def test_calender():
    calendar = Calender(
        id='1',
        date=datetime.strptime('24/01/2022', "%d/%m/%Y")
    )
    assert isinstance(calendar, Calender)
    assert str(calendar), '2022-01-24 00:00:00'

