package com.todo.service.task;

import com.todo.entity.TodoTask;
import com.todo.model.TodoTaskDTO;

import java.util.List;

public interface TodoTaskService {
    List<TodoTask> findTodoTaskByAppAccount_Username(String username);

    int countByAppAccount_Username(String username);

    TodoTask create(TodoTask todoTask);
    TodoTask update(Long id, TodoTaskDTO dto);
}
