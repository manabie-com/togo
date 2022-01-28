package com.thang.todo.repositories;

import com.thang.todo.entities.Task;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface TaskRepository extends JpaRepository<Task, Long> {
    List<Task> findByUserId(Long userId);

    Optional<Task> findByIdAndUserId(Long id, Long userId);

    @Query(value = "select count(*) from tasks where updated_at >= CURDATE() and updated_at < DATE_ADD(CURDATE(), INTERVAL 1 DAY);  ", nativeQuery = true)
    Long countByUserId(Long userId);
}