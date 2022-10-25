from .base import BaseInterface, Base
from sqlalchemy import Column, String, Boolean


class User(Base, BaseInterface):
    """ Mapping from python class to table ip_product_uptime_by_month """
    __tablename__ = 'user'
    id = Column('id', String(255), primary_key=True, unique=True)
    username = Column('username', String(255))
    email = Column('email', String(255))
    salt = Column("salt", String(255))
    hashed = Column("hashed", String(255))
    name = Column("fullname", String(255))
    pricing = Column("pricing_id",String(255))
    deleted = Column("deleted", Boolean, default=False)
