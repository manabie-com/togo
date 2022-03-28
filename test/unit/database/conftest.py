import os
import pytest
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker, scoped_session

from database import Base

Session = sessionmaker()
engine = create_engine(
    "postgresql+psycopg2://{}:{}@{}:{}/{}".format(
        os.environ.get("DB_USER", "postgres"),
        os.environ.get("DB_PASSWORD", "1234"),
        os.environ.get("DB_HOST", "localhost"),
        os.environ.get("DB_PORT", "5432"),
        os.environ.get("DB_NAME", "postgres"),
    )
)


@pytest.fixture(scope="module")
def connection():
    connection = engine.connect()
    yield connection
    connection.close()


@pytest.fixture(scope="module")
def setup_database(connection):
    Base.metadata.bind = connection
    Base.metadata.create_all()
    yield
    Base.metadata.drop_all()


@pytest.fixture(scope='function')
def db_session(setup_database, connection):
    """Create a database session
    Using a transaction, example:
    transaction = connection.begin()
    yield scope_session(sessionmaker(autocommit=False, autoflush=False, bin=connection))
    transaction.rollback()
    """
    session = scoped_session(sessionmaker(autocommit=False, autoflush=False, bind=connection))
    yield session


def seek_data(db_session, records):
    for record in records:
        db_session.add(record)
    db_session.commit()


def cleanup(db_session, records):
    for record in records:
        db_session.delete(record)
    db_session.commit()
