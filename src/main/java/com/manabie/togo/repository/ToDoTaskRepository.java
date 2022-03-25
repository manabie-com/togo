package com.manabie.togo.repository;

import com.manabie.togo.domain.ToDoTask;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface ToDoTaskRepository extends JpaRepository<ToDoTask, String> {

    List<ToDoTask> findByUserIdAndToDoDate(String userId, String toDoDate);
}
