from sqlalchemy import Column, Date, Integer
from sqlalchemy.orm import relationship
from sqlalchemy.orm.collections import InstrumentedList

from database.models import BaseEntityModel


class Calender(BaseEntityModel):
    __tablename__ = "calender"

    id = Column(Integer, primary_key=True, autoincrement=True)
    date = Column(Date, primary_key=True, unique=True)

    assignments: InstrumentedList = relationship("Assignment", lazy="select", cascade="all, delete-orphan")

    def __repr__(self):
        return f"{self.date}"
