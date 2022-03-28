from datetime import datetime
from unittest.mock import MagicMock

import pytest
from pytest_mock import MockerFixture

from database import User, Assignment, Task
from handler.task import HandlerTask


@pytest.fixture
def handler():
    handler = HandlerTask(MagicMock())
    return handler


def test_add_and_assign_task(handler, mocker: MockerFixture):
    task = {
        "user_id": 1,
        "date": "24/01/2022"
    }
    mock_ability_caculator = mocker.patch.object(HandlerTask, "ability_caculator", return_value=True)
    handler.assign_task = MagicMock()

    rs = handler.add_and_assign_task(task)

    assert rs, True
    assert mock_ability_caculator.call_count == 1
    assert handler.assign_task.call_count == 1


def test_add_and_assign_task_false(handler, mocker: MockerFixture):
    task = {
        "user_id": 1,
        "date": "24/01/2022"
    }
    mock_ability_caculator = mocker.patch.object(HandlerTask, "ability_caculator", return_value=False)
    rs = handler.add_and_assign_task(task)
    assert (rs==False), True
    assert mock_ability_caculator.call_count == 1


def test_ability_caculator(handler):
    user = User(id=1, name="test", ability=3)
    assignment = Assignment(user_id=1, date=datetime.strptime('24/01/2022', "%d/%m/%Y"))
    task = Task(id=1, title="1")
    task2 = Task(id=2, title="2")
    assignment.tasks.append(task)
    assignment.tasks.append(task2)
    handler._db.get_user_by_id.return_value = user
    handler._db.get_assignment_by_date.return_value = assignment

    rs = handler.ability_caculator(1, '24/01/2022')

    assert rs, True


def test_ability_caculator_false(handler):
    user = User(id=1, name="test", ability=3)
    assignment = Assignment(user_id=1, date=datetime.strptime('24/01/2022', "%d/%m/%Y"))
    task = Task(id=1, title="1")
    task2 = Task(id=2, title="2")
    task3 = Task(id=3, title="3")
    assignment.tasks.append(task)
    assignment.tasks.append(task2)
    assignment.tasks.append(task3)
    handler._db.get_user_by_id.return_value = user
    handler._db.get_assignment_by_date.return_value = assignment

    rs = handler.ability_caculator(1, '24/01/2022')

    assert (rs==False), True


def test_assign_task(handler):
    task = Task(id=1, title="test")
    handler.assign_task(task)

    assert handler._db.add.call_count, 1
    assert handler._db.remove_session.call_count, 1
