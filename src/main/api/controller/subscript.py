from src.main.util import load_config
from sqlalchemy import create_engine, and_
from sqlalchemy.sql import update
from sqlalchemy.orm.session import Session
from src.main.model import Task, Pricing, Base, Invoice, User
from uuid import uuid4
from datetime import datetime


def subscript(user_id: str, pricing_id: str):
    connect_string = load_config().get("DATABASE", {}).get("connection", None)
    if connect_string is None:
        raise Exception("Fail to connect to database!")
    engine = create_engine(connect_string)
    Base.metadata.create_all(engine, checkfirst=True)
    invoice_id = uuid4().hex
    invoice = Invoice.from_dict({
        "id": invoice_id,
        "pricing_id": pricing_id,
        "paid_date": datetime.now(),
        "duration": 30
    })
    print(user_id, pricing_id)
    with Session(engine) as session, session.begin():
        session.add(invoice)
        session.execute(
            update(User).where(and_(
                (User.id == user_id),
                (User.deleted == False)
            )).values(
                pricing=pricing_id
            ))
    return {"invoice_id": invoice_id}


def get_pricing_level_by(id: str = None, name: str = None) -> dict:
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
            return {}
        data = {
            "id": pricing.id,
            "name": pricing.name,
            "unit_price": pricing.unit_price,
            "daily_limit": pricing.daily_limit
        }
    return data


def get_pricing_level():
    connect_string = load_config().get("DATABASE", {}).get("connection", None)
    if connect_string is None:
        raise Exception("Fail to connect to database!")
    engine = create_engine(connect_string)
    with Session(engine) as session, session.begin():
        tasks = session.query(
            Pricing.id, Pricing.name, Pricing.unit_price, Pricing.daily_limit
        ).where(and_(
            (Pricing.deleted == False)
        ))
        data = []
        i =0
        for task in tasks:
            print(i)
            i+=1
            data.append(task._asdict())
    return data
