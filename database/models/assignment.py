from sqlalchemy import Column, Integer, Date, ForeignKey, String

from database.models import BaseEntityModel

OPEN = "Open"
IP = "In Process"
RV = "Review"
DN = "Done"
C = "Close"


class Assignment(BaseEntityModel):
    __tablename__ = "assignment"
    id = Column(Integer, primary_key=True, autoincrement=True)
    user_date_id = Column(Integer, ForeignKey("user_date.id"))
    task_id = Column(Integer, ForeignKey("task.id"))
    status = Column(String(10), default=OPEN)
    mark = Column(Integer, nullable=True)
    comment = Column(String(255))

    def __repr__(self):
        return f"{self.user_date_id} - {self.status}"
