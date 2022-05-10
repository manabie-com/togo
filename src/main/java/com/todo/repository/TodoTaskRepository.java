package com.todo.repository;


import com.todo.entity.TodoTask;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;


@Repository
public interface TodoTaskRepository extends JpaRepository<TodoTask, Long> {

    List<TodoTask> findTodoTaskByAppAccount_Username(String username);

    int countByAppAccount_Username(String username);
}