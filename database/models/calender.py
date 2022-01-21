from sqlalchemy import Column, Date
from sqlalchemy.orm import relationship
from sqlalchemy.orm.collections import InstrumentedList

from database.models import BaseEntityModel


class Calender(BaseEntityModel):
    __tablename__ = "calender"

    date = Column(Date, primary_key=True, unique=True)

    user_dates: InstrumentedList = relationship("DateOfUser", lazy="select", cascade="all, delete-orphan")

    def __repr__(self):
        return f"{self.date}"
