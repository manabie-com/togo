from .base import BaseInterface, Base
from sqlalchemy import Column, String, Boolean, DateTime, Integer, Numeric


class Pricing(Base, BaseInterface):
    """ Mapping from python class to table ip_product_uptime_by_month """
    __tablename__ = 'pricing'
    id = Column('id', String(255), primary_key=True, unique=True)
    name = Column("name", String(255))
    unit_price = Column("unit_price", Numeric(30, 2))
    daily_limit = Column("daily_limit", Integer)
    deleted = Column("deleted", Boolean, default=False)
