from .base import BaseInterface, Base
from sqlalchemy import Column, String, Boolean, DateTime


class Task(Base, BaseInterface):
    """ Mapping from python class to table ip_product_uptime_by_month """
    __tablename__ = 'task'
    id = Column('id', String(255), primary_key=True, unique=True)
    user_id = Column('user_id', String(255))
    summary = Column('summary', String(255))
    description = Column('description', String(255))
    created = Column('created', DateTime)
    last_modified = Column('last_modified', DateTime)
    pricing = Column("pricing", String(255))
    finish = Column("finish", Boolean,default=False)
    deleted = Column("deleted", Boolean, default=False)
