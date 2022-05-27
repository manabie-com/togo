package com.todo.core.todo.service;

import com.todo.core.commons.model.GenericResponse;
import com.todo.core.todo.application.dto.TodoDTO;
import com.todo.core.todo.model.Todo;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.transaction.annotation.Transactional;

public interface TodoService {

    GenericResponse<Integer> addTodo(Long todoUserId, TodoDTO todo);

    @Transactional(readOnly = true)
    GenericResponse<Page<Todo>> retrieveAllTodoByUser(Long userId, Pageable pageable);
}
