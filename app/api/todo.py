from app.task.service_layer.services import AddTodoListService
from app.task.service_layer.unit_of_work import TodoListUnitOfWork
from . import api_blueprint
from app.exceptions import InvalidTitleError, DuplicateTitleError, ExceedLimitationError

from flask import request

@api_blueprint.route('/todo', methods=['POST'])
def create_todo():
    try:
        data = request.json
        title = data.get("title") or ""
        limit = data.get("limit") or 0

        if not title:
            raise InvalidTitleError("Title is invalid.")
        description = data.get("description")
        todos = data.get("todos") or []

        AddTodoListService.add(title, description, todos, limit, TodoListUnitOfWork())
    except (InvalidTitleError, DuplicateTitleError, ExceedLimitationError) as e:
        return {
            "message": str(e)
        }, 400
    return {
        "message": "Todo List created successfully."
    }, 201