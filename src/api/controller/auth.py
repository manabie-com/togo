from src.util import load_config, logger
from sqlalchemy import create_engine
from sqlalchemy.orm.session import Session
from src.model import User
from uuid import uuid4
from ..services import create_token
from .subscript import get_pricing_level_by
import bcrypt


def create_user(email: str, password: str, **kwargs) -> dict:
    fullname = kwargs.get("fullname")
    connect_string = load_config().get("DATABASE", {}).get("connection", None)
    if connect_string is None:
        raise Exception("Fail to connect to database!")
    engine = create_engine(connect_string)
    if not check_email(email):
        raise Exception("An user with email already existed!")
    user_id = uuid4().hex
    salt = bcrypt.gensalt(10)
    hashed = bcrypt.hashpw(password, salt)
    basic_pricing_level_info = get_pricing_level_by(name="Basic")
    basic_pricing_id = basic_pricing_level_info.get("id")
    user_dict = {
        "id": user_id,
        "username": user_id,
        "email": email,
        "salt": salt,
        "hashed": hashed,
        "name": fullname,
        "pricing": basic_pricing_id
    }
    user = User.from_dict(user_dict)
    with Session(engine) as session, session.begin():
        session.add(user)
    return user_dict


def check_email(email: str) -> bool:
    connect_string = load_config().get("DATABASE", {}).get("connection", None)
    if connect_string is None:
        raise Exception("Fail to connect to database!")
    engine = create_engine(connect_string)
    with Session(engine) as session, session.begin():
        no_email = session.query(User).where(User.email == email).count()
        return no_email < 1


def validate_credential(email: str, password: str) -> str:
    logger.info(email)
    logger.info(password)
    connect_string = load_config().get("DATABASE", {}).get("connection", None)
    if connect_string is None:
        raise Exception("Fail to connect to database!")
    engine = create_engine(connect_string)
    with Session(engine) as session, session.begin():
        user = session.query(User).where(User.email == email).first()
        hashed = bcrypt.hashpw(password, user.salt)
        if hashed == user.hashed:
            return create_token({
                "userId": user.id
            })
    return ""
