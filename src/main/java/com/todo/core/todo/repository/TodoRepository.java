package com.todo.core.todo.repository;

import com.todo.core.todo.model.Todo;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;

import java.time.LocalDate;
import java.util.List;

public interface TodoRepository extends JpaRepository<Todo, Long> {

    Page<Todo> findAllByTodoUserId(Long id, Pageable pageable);

    int countByDateCreatedAndTodoUserId(LocalDate dateCreate, Long userId);
}
