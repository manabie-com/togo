package com.manabie.todotaskapplication.repository.task;

import com.manabie.todotaskapplication.data.model.Task;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.CrudRepository;

import javax.transaction.Transactional;
import java.util.List;
import java.util.UUID;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
public interface TaskRepository extends JpaRepository<Task, UUID> {
    @Modifying
    @Transactional
    @Query("UPDATE Task t set t.name = ?1, t.content = ?2 where t.id = ?3")
    int updateTask(String name, String content, UUID id);

}
