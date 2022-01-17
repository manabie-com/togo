from sqlalchemy.orm import Session
from fastapi import APIRouter, Depends, HTTPException
from datetime import date
import models
import schemas
from database import SessionLocal, engine


router = APIRouter()
models.Base.metadata.create_all(bind=engine)


def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()


@router.post('/', response_model=schemas.Task)
def create_task(task: schemas.TaskCreate, db: Session = Depends(get_db)):
    user = db.query(models.User).filter(models.User.id == task.user_id).first()
    if not user:
        raise HTTPException(400, 'user not exist')

    if user.limit > 0:
        n_task = db.query(models.Task)\
            .filter(models.Task.user_id == task.user_id,
                    models.Task.create_date == date.today())\
            .count()
        if n_task >= user.limit:
            raise HTTPException(429, 'number of tasks per day reached')

    new_task = models.Task(user_id=task.user_id, text=task.text)
    db.add(new_task)
    db.commit()
    db.refresh(new_task)
    return new_task
