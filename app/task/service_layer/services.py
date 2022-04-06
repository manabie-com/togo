from app.task.adapters.repository import TodoListRepository
from app.task.domain.model import TodoItem, TodoList, AbstractTodo
from app.task.service_layer.unit_of_work import TodoListUnitOfWork
from typing import List
from app.exceptions import DuplicateTitleError, ExceedLimitationError

class AddTodoListService:
    @staticmethod
    def add(title:str, description:str, todos: List[dict], limit: int, uow: TodoListUnitOfWork):
        with uow:
            existing_todo_list = uow.repository.get(title=title)
            if existing_todo_list:
                raise DuplicateTitleError("Cannot add todo list due to duplicate title.")

            todo_list = TodoList(
                title=title, 
                description=description, 
                limit=limit
            )

            if limit != 0 and len(todos) > limit:
                raise ExceedLimitationError("Cannot input more todos than limitation")

            todo_lists, todo_items = AddTodoListService.create_todo(todos)
            
            todo_list.children.extend(todo_lists)
            todo_list.todo_items.extend(todo_items)
            
            uow.add(todo_list)
            uow.commit()

    def create_todo(todos: List[dict]):
        todo_lists = []
        todo_items = []
        for todo in todos:
            if isinstance(todo, dict) and todo.get("title", ""):
                if todo.get("todos"):
                    todo_lists, todo_items = AddTodoListService.create_todo(todo.get("todos"))

                    todo_lists.append(
                        TodoList(
                            title=todo.get("title", ""), 
                            description=todo.get("description"),
                            limit=0,
                            children=todo_lists,
                            todo_items=todo_items
                        )
                    )
                else:
                    todo_items.append(
                        TodoItem(
                            title=todo.get("title", ""), 
                            description=todo.get("description"),
                        )
                    )
        return todo_lists, todo_items
