package com.api.todo.repositories;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import com.api.todo.entities.Task;

@Repository
public interface TaskRepository extends JpaRepository<Task, String> {
    @Query(value = "SELECT COUNT(t.id) FROM task t WHERE t.user_id=:user_id and DATE(t.created_date)=:created_date", nativeQuery = true)
    int getCountTaskOfUser(@Param("user_id") long userId, @Param("created_date") String createdDate);
}
