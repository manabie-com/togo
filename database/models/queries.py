from sqlalchemy.orm import scoped_session, Session

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
