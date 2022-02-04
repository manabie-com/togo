from .base import BaseInterface, Base
from sqlalchemy import Column, String, Boolean, DateTime, Integer, Numeric


class Invoice(Base, BaseInterface):
    """ Mapping from python class to table ip_product_uptime_by_month """
    __tablename__ = 'invoice'
    id = Column('id', String(255), primary_key=True, unique=True)
    user_id = Column("name", String(255))
    pricing_id = Column("unit_price", String(255))
    paid_date = Column("paid_date", DateTime)
    duration = Column("duration", Integer)
