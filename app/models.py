from sqlalchemy import ForeignKey, PrimaryKeyConstraint
from .app import db


# User model
class Users(db.Model):
    __tablename__ = "users"
    id = db.Column(db.Integer, primary_key=True)
    public_id = db.Column(db.Integer)
    name = db.Column(db.String(50), unique=True)
    password = db.Column(db.String(50))
    limit_per_day = db.Column(db.Integer, nullable=False)
    tasks = db.relationship("Todos")


# Todo task model
class Todos(db.Model):
    __tablename__ = "todos"
    id = db.Column(db.Integer, primary_key=True)
    task = db.Column(db.String(100), nullable=False)
    user_id = db.Column(db.Integer, ForeignKey("users.public_id"), nullable=False)
    date_time = db.Column(db.Date)
