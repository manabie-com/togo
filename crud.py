from sqlalchemy.orm import Session
from datetime import date
import schemas
import models


def get_user(db: Session, user_id: int):
    return db.query(models.User).filter(models.User.id == user_id).first()


def count_task(db: Session, user_id: int):
    return db.query(models.Task) \
            .filter(models.Task.user_id == user_id,
                    models.Task.create_date == date.today()) \
            .count()


def create_task(db: Session, task: schemas.TaskCreate):
    new_task = models.Task(user_id=task.user_id, name=task.name)
    db.add(new_task)
    db.commit()
    db.refresh(new_task)
    return new_task
