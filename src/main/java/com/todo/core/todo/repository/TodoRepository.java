package com.todo.core.todo.repository;

import com.todo.core.todo.model.Todo;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface TodoRepository extends JpaRepository<Todo, Long> {

    List<Todo> findAllByTodoUserId(Long id);
}
