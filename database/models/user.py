from sqlalchemy import Integer, String, Column
from sqlalchemy.orm import relationship
from sqlalchemy.orm.collections import InstrumentedList

from .base import BaseEntityModel


class User(BaseEntityModel):
    __tablename__ = "user"

    id = Column(Integer, primary_key=True, autoincrement=True)
    name = Column(String(250))
    ability = Column(Integer)
    assignments: InstrumentedList = relationship("assignmet", lazy="select", cascade="all, delete-orphan")

    def __repr__(self):
        return f"{self.name}"

