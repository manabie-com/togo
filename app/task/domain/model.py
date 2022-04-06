from dataclasses import dataclass
from datetime import date
from typing import Optional, List, Set
from abc import ABC, abstractmethod
from app import db
from sqlalchemy import Column, Table, func, DateTime, String, Integer, ForeignKey

class Base(db.Model):
    __abstract__ = True

    created_on = db.Column(db.DateTime, default=db.func.now())
    updated_on = db.Column(db.DateTime, default=db.func.now(), onupdate=db.func.now())

class AbstractTodo(Base):
    __abstract__ = True

    id = db.Column(db.Integer, primary_key=True, autoincrement=True)

    title = db.Column(db.String(255), unique=True)
    description = db.Column(db.String(255))
    due_date = db.Column(db.DateTime)

class TodoItem(AbstractTodo):
    todo_list_id = db.Column(Integer, db.ForeignKey("todo_list.id"))

class TodoList(AbstractTodo):
    limit = db.Column(Integer, default=0)
    todo_items = db.relationship("TodoItem")
    parent_id = db.Column(Integer, ForeignKey('todo_list.id'))
    children = db.relationship("TodoList",
        backref=db.backref('parent', remote_side='TodoList.id')
    )
