from src.main.util import load_config
from sqlalchemy import create_engine
from sqlalchemy.orm.session import Session
from src.main.model import Task, Pricing, User
from sqlalchemy.sql.expression import select, update
from sqlalchemy import and_
from typing import List
from datetime import datetime, timedelta
from uuid import uuid4
from ..error.basic import HTTPError


# Base.metadata.create_all(engine, checkfirst=True)
# session.add(Pricing.from_dict({
#     "id": uuid4().hex,
#     "name": "Basic",
#     "unit_price": 0,
#     "daily_limit": 5,
# }))
# session.add(Pricing.from_dict({
#     "id": uuid4().hex,
#     "name": "Standard",
#     "unit_price": 2.99,
#     "daily_limit": 10,
# }))
# session.add(Pricing.from_dict({
#     "id": uuid4().hex,
#     "name": "Premium",
#     "unit_price": 4.99,
#     "daily_limit": 40,
# }))
# session.add(Pricing.from_dict({
#     "id": uuid4().hex,
#     "name": "Enterprise",
#     "unit_price": 9.99,
#     "daily_limit": 80,
# }))


def daily_task_check(user_id: str):
    connect_string = load_config().get("DATABASE", {}).get("connection", None)
    if connect_string is None:
        raise Exception("Fail to connect to database!")
    engine = create_engine(connect_string)

    with Session(engine) as session, session.begin():
        begin_date = datetime.now()
        begin_date = begin_date.replace(hour=0, minute=0, second=0, microsecond=0)
        end_date = begin_date + timedelta(days=1)
        no_created_task = session.query(Task).where(and_(
            (Task.user_id == user_id),
            (Task.created >= begin_date),
            (Task.created < end_date),
            (Task.deleted == False)
        )).count()
        user = session.query(User).where(User.id == user_id).first()
        pricing_info = get_pricing_level(id=user.pricing)
        if pricing_info is None:
            raise HTTPError(404, f"Pricing Level of {user_id} is not existed!")
        if no_created_task < int(pricing_info.get("daily_limit")):
            return True
        return False


def get_task_of(user_id) -> List[dict]:
    connect_string = load_config().get("DATABASE", {}).get("connection", None)
    if connect_string is None:
        raise Exception("Fail to connect to database!")
    engine = create_engine(connect_string)
    with Session(engine) as session, session.begin():
        tasks = session.query(
            Task.id, Task.summary, Task.description, Task.finish, Task.created, Task.last_modified
        ).where(and_(
            (Task.user_id == user_id)),
            (Task.deleted == False)
        )
        data = []
        for task in tasks:
            data.append(task._asdict())
    return data


def get_pricing_level(id: str = None, name: str = None) -> dict:
    connect_string = load_config().get("DATABASE", {}).get("connection", None)
    if connect_string is None:
        raise Exception("Fail to connect to database!")
    engine = create_engine(connect_string)
    with Session(engine) as session, session.begin():
        if name is not None:
            pricing = session.query(Pricing).where(Pricing.name == name).first()
        elif id is not None:
            pricing = session.query(Pricing).where(Pricing.id == id).first()
        else:
            pricing = None
        if pricing is None:
            return None
        data = {
            "id": pricing.id,
            "name": pricing.name,
            "unit_price": pricing.unit_price,
            "daily_limit": pricing.daily_limit
        }
    return data


def create_task(user_id, task_data: dict):
    connect_string = load_config().get("DATABASE", {}).get("connection", None)
    if connect_string is None:
        raise Exception("Fail to connect to database!")
    engine = create_engine(connect_string)
    task_id = uuid4().hex
    task = Task.from_dict({
        **task_data,
        "id": task_id,
        "user_id": user_id,
        "created": datetime.now(),
        "last_modified": datetime.now()
    })
    if daily_task_check(user_id):
        with Session(engine) as session, session.begin():
            session.add(task)
        return task_id
    else:
        raise HTTPError(429, f"Daily limit for user '{user_id}' exceed! Please upgrade to higher pricing options.")


def update_task(task_id, task_data):
    connect_string = load_config().get("DATABASE", {}).get("connection", None)
    if connect_string is None:
        raise Exception("Fail to connect to database!")
    engine = create_engine(connect_string)
    with Session(engine) as session, session.begin():
        session.execute(
            update(Task).where(Task.id == task_id).values(
                **task_data
            ))
    pass


def delete_task(task_id):
    connect_string = load_config().get("DATABASE", {}).get("connection", None)
    if connect_string is None:
        raise Exception("Fail to connect to database!")
    engine = create_engine(connect_string)
    with Session(engine) as session, session.begin():
        session.execute(
            update(Task).where(Task.id == task_id).values(
                deleted=True
            ))
    pass
