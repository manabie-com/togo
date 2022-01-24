from sqlalchemy import Integer, String, Column, ForeignKey

from .base import BaseEntityModel


class Task(BaseEntityModel):
    __tablename__ = "task"

    id = Column(Integer, primary_key=True, autoincrement=True)
    title = Column(String(50))
    description = Column(String(1023), nullable=True)
    assignment_id = Column(Integer, ForeignKey("assignment.id"))

    def __repr__(self):
        return f"{self.title} - {self.description}"
