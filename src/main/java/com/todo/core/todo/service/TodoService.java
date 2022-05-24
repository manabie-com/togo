package com.todo.core.todo.service;

import com.todo.core.commons.model.GenericResponse;
import com.todo.core.todo.application.dto.TodoDTO;

public interface TodoService {

    GenericResponse<Integer> addTodo(Long todoUserId, TodoDTO todo);

}
