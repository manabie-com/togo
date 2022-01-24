from datetime import datetime
from unittest.mock import MagicMock

import pytest
from sqlalchemy.exc import SQLAlchemyError


from database import User, Assignment
from database.models.queries import DBSession


@pytest.fixture()
def db(db_session):
    db = DBSession()
    db._session = db_session
    return db


@pytest.fixture
def db_magic():
    db = DBSession()
    db._session = MagicMock()
    return db


def test_get_user_by_id_sql_exception(db_magic):
    db_magic.session.query.side_effect = SQLAlchemyError("error")
    rs = db_magic.get_user_by_id(1)
    assert db_magic._session.query.call_count, 1


def test_get_user_by_id(db):
    user = User(name="user_test", ability=3)
    db.add(user)

    rs = db.get_user_by_id(1)
    assert isinstance(rs, User)


def test_get_assignment_by_date(db):
    assignment = Assignment(user_id=1, date=datetime.strptime('24/01/2022', "%d/%m/%Y"))
    db.add(assignment)
    rs = db.get_assignment_by_date(1, datetime.strptime('24/01/2022', "%d/%m/%Y"))
    assert isinstance(rs, Assignment)


def test_add(db_magic):
    user = User(name="user_test")
    db_magic.add(user)

    assert db_magic.session.add.call_count == 1
    assert db_magic.session.commit.call_count == 1


def test_add_exception(db_magic):
    user = User(name="user_test")
    db_magic.session.add.side_effect = SQLAlchemyError("error")
    db_magic.add(user)

    assert db_magic.session.add.call_count == 1
    assert db_magic.session.commit.call_count == 0
    assert db_magic.session.rollback.call_count == 1



