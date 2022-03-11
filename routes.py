from sqlalchemy.orm import Session
from fastapi import APIRouter, Depends, HTTPException
from database import SessionLocal, engine
import crud
import schemas
import models

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
    user = crud.get_user(db, task.user_id)
    if not user:
        raise HTTPException(400, 'user not exist')

    if user.limit > 0:
        n_task = crud.count_task(db, task.user_id)
        if n_task >= user.limit:
            raise HTTPException(429, 'number of tasks per day reached')

    return crud.create_task(db, task)
