package com.manabie.todo.repository;

import com.manabie.todo.entity.TaskCounterEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface TaskCounterRepository extends JpaRepository<TaskCounterEntity, String> {
    TaskCounterEntity findByUserId(Long userId);
}
