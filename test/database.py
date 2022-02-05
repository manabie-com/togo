from src.util import load_config, logger
from src.model import Base, Pricing, User
from sqlalchemy import create_engine
from sqlalchemy.orm.session import Session
from uuid import uuid4
from src.api.controller.auth import create_user

PRICING = [
    {
        "id": uuid4().hex,
        "name": "Basic",
        "unit_price": 0,
        "daily_limit": 5,
    },
    {
        "id": uuid4().hex,
        "name": "Standard",
        "unit_price": 2.99,
        "daily_limit": 10,
    }, {
        "id": uuid4().hex,
        "name": "Premium",
        "unit_price": 4.99,
        "daily_limit": 40,
    }, {
        "id": uuid4().hex,
        "name": "Enterprise",
        "unit_price": 9.99,
        "daily_limit": 80,
    }
]
USER = [
    {
        "email": "test-mail-1@gmail.com",
        "password": "easy-peasy",
        "fullname": "First User"
    },
    {
        "email": "test-mail-2@gmail.com",
        "password": "lemon-squeeze",
        "fullname": "Second User"
    },
    {
        "email": "test-mail-3@gmail.com",
        "password": "naruto-tobacco",
        "fullname": "Third User"
    }
]


def init_db():
    init_db_mock_data()

    for i, user in enumerate(USER):
        email = user.pop("email")
        password = user.pop("password")
        logger.info(f"{email} - {password}")
        USER[i] = {**user, **create_user(email, password, **user), "password": password}


def init_db_mock_data():
    connect_string = load_config().get("DATABASE", {}).get("connection", None)
    if connect_string is None:
        raise Exception("Fail to connect to database!")
    engine = create_engine(connect_string)
    with Session(engine) as session, session.begin():
        Base.metadata.drop_all(engine)
        Base.metadata.create_all(engine)
        for pricing in PRICING:
            session.add(Pricing.from_dict(pricing))
