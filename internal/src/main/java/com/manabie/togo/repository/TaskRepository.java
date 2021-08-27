package com.manabie.togo.repository;

import com.manabie.togo.model.Task;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

/**
 * Data access layer for Tasks
 * @author mupmup
 */
@Repository
public interface TaskRepository extends JpaRepository<Task, String> {

    
}
