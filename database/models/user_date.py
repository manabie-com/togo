from sqlalchemy import Column, Integer, Date, ForeignKey

from database.models import BaseEntityModel


class DateOfUser(BaseEntityModel):
    __tablename__ = "user_date"
    id = Column(Integer, primary_key=True, autoincrement=True)
    user_id = Column(Integer, ForeignKey("user.id"))
    date = Column(Date, ForeignKey("calender.date"))

    def __repr__(self):
        return f"{self.date} - {self.user_id}"
