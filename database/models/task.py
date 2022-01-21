from sqlalchemy import Integer, String, Column
from sqlalchemy.orm import relationship
from sqlalchemy.orm.collections import InstrumentedList

from .base import BaseEntityModel


class Task(BaseEntityModel):
    __tablename__ = "task"

    id = Column(Integer, primary_key=True, autoincrement=True)
    title = Column(String(50))
    description = Column(String(1023), nullable=True)

    assignments: InstrumentedList = relationship("Assignment", lazy="select", cascade="all, delete-orphan")

    def __repr__(self):
        return f"{self.title} - {self.description}"
