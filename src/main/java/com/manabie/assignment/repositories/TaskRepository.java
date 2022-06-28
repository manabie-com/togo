package com.manabie.assignment.repositories;

import com.manabie.assignment.models.User;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.CrudRepository;

import com.manabie.assignment.models.Task;
import org.springframework.data.repository.query.Param;

import java.sql.Date;
import java.util.List;

public interface TaskRepository extends CrudRepository<Task, Integer> {

    @Query("SELECT t FROM Task t WHERE t.user = :user and t.taskDate = :taskDate")
    List<Task> findAllByUserIdAndDate(@Param("user") User user, @Param("taskDate") Date taskDate);
}
