import os

from sqlalchemy import create_engine, Column, Integer
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import Session, scoped_session, sessionmaker
from sqlalchemy.util.compat import contextmanager

engine = create_engine(
    "postgresql+psycopg2://{}:{}@{}:{}/{}".format(
        os.environ.get("DB_USER", "postgres"),
        os.environ.get("DB_PASSWORD", "1234"),
        os.environ.get("DB_HOST", "localhost"),
        os.environ.get("DB_PORT", "5432"),
        os.environ.get("DB_NAME", "postgres"),
    )
)
Base = declarative_base()


class BaseEntityModel(Base):
    __abstract__ = True
    id = Column(Integer, primary_key=True, autoincrement=True)


LocalSession = sessionmaker(bind=engine, autocommit=False, autoflush=False)


@contextmanager
def session_scope() -> Session:
    scope = scoped_session(LocalSession)
    session = scope()
    try:
        yield session
        session.commit()
    except Exception as ex:
        session.rollback()
        raise ex
    finally:
        session.close()
        session.remove()

