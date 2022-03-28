from sqlalchemy.exc import SQLAlchemyError
from sqlalchemy.orm import scoped_session, Session

from database import User, Assignment, Calender
from database.models.base import LocalSession


class DBSession:
    def __init__(self):
        self._session_factory = scoped_session(LocalSession)
        self._session: Session = self._session_factory()

    def refresh(self):
        if self._session:
            self._session.close()
        self._session = self._session_factory()

    @property
    def session(self) -> Session:
        return self._session

    @session.setter
    def session(self, session_factory):
        self._session_factory = session_factory
        self._session = self._session_factory()

    def remove_session(self):
        self._session.close()
        self._session_factory.remove()

    def get_user_by_id(self, user_id):
        try:
            result = (
                self._session.query(User)
                .filter(User.id == user_id)
                .first()
            )
            return result
        except SQLAlchemyError as e:
            print(str(e))
            return None

    def get_assignment_by_date(self, user_id, date):
        try:
            result = (
                self._session.query(Assignment)
                .filter(Assignment.user_id == user_id)
                .join(Calender)
                .filter(Calender.date == date)
                .first()
            )
            return result
        except SQLAlchemyError as e:
            print(str(e))
            return None

    def add(self, obj):
        try:
            self._session.add(obj)
            self._session.commit()
        except Exception as e:
            self._session.rollback()
            print(str(e))
