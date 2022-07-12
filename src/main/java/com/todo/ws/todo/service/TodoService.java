package com.todo.ws.todo.service;

import com.todo.ws.commons.model.ResponseEntity;
import com.todo.ws.todo.application.dto.TodoDTO;
import com.todo.ws.todo.model.Todo;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.transaction.annotation.Transactional;

public interface TodoService {

    ResponseEntity<Integer> addTodo(Long todoUserId, TodoDTO todo);

    @Transactional(readOnly = true)
    ResponseEntity<Page<Todo>> retrieveAllTodoByUser(Long userId, Pageable pageable);
}
